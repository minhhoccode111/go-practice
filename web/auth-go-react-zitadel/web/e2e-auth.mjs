// End-to-end smoke test: drives the PKCE login flow in headless Chrome.
// Usage: node e2e-auth.mjs
import puppeteer from "puppeteer-core";

const APP = "http://localhost:5174";
const USER = process.env.E2E_USER ?? "zitadel-admin@zitadel.localhost";
const PASS = process.env.E2E_PASS ?? "Password1!";

const log = (...a) => console.log("[e2e]", ...a);

function sleep(ms) {
  return new Promise((r) => setTimeout(r, ms));
}

async function waitFor(fn, desc, timeout = 20000) {
  const start = Date.now();
  while (Date.now() - start < timeout) {
    try {
      const v = await fn();
      if (v) return v;
    } catch {}
    await sleep(400);
  }
  throw new Error(`timeout waiting for ${desc}`);
}

const browser = await puppeteer.launch({
  executablePath: "/usr/bin/google-chrome",
  headless: "new",
  args: ["--no-sandbox", "--disable-dev-shm-usage"],
});

try {
  const page = await browser.newPage();
  await page.setViewport({ width: 1280, height: 900 });

  log(`open ${APP}`);
  await page.goto(APP, { waitUntil: "networkidle2" });

  log("click Sign in");
  const signInBtn = await waitFor(() => page.$("button:not([disabled])"));
  const btnText = await signInBtn.evaluate((b) => b.textContent);
  log(`button: "${btnText}"`);
  await signInBtn.click();

  log("waiting for ZITADEL login page …");
  await waitFor(() => page.url().includes("/login/login") || page.url().includes("authRequest"), "redirect to login");
  log(`on ${page.url().slice(0, 80)}`);

  const fillField = async (selector, value) => {
    const el = await waitFor(() => page.$(selector), `field ${selector}`, 15000);
    await el.click({ clickCount: 3 });
    await el.type(value);
  };
  const clickByText = async (label) => {
    const handle = await waitFor(
      () => page.evaluateHandle((l) => {
        const btn = Array.from(document.querySelectorAll("button")).find((b) => b.textContent?.trim().toLowerCase().includes(l));
        return btn || null;
      }, label.toLowerCase()),
      `button "${label}"`,
      15000
    );
    const ok = !(await handle.evaluate((h) => h === null));
    if (!ok) {
      const dbg = await page.evaluate(() => ({
        text: document.body?.innerText?.slice(0, 500),
        btns: Array.from(document.querySelectorAll("button, input[type=submit], a"))
          .map((b) => (b.textContent || b.getAttribute("value") || b.getAttribute("href") || "").trim())
          .filter(Boolean),
      }));
      process.stderr.write("[e2e] DEBUG: " + JSON.stringify(dbg) + "\n");
      throw new Error(`button "${label}" not found`);
    }
    await handle.click();
  };

  // Login V1 is two-step: loginName → Next → password → submit.
  await fillField('input[name="loginName"]', USER);
  log("step 1: Next …");
  await clickByText("Next");
  await fillField('input[type="password"]', PASS);
  log("step 2: submit …");
  await Promise.all([
    page.waitForNavigation({ waitUntil: "networkidle2", timeout: 30000 }).catch(() => {}),
    clickByText("Next"),
  ]);

  // Initial admin is forced to TOTP setup once; skip it. The Skip button's
  // click handler is flaky under automation, so submit the form via
  // requestSubmit(skipButton), which works reliably.
  const mfaText = await page.evaluate(() => document.body?.innerText ?? "");
  if (mfaText.includes("2-Factor") || mfaText.includes("MFA")) {
    log("MFA setup prompt detected; skipping …");
    const nav = page.waitForNavigation({ waitUntil: "networkidle2", timeout: 30000 }).catch(() => {});
    await page.evaluate(() => {
      const form = document.querySelector("form");
      const skip = document.querySelector("button[name=skip]");
      if (form && skip) form.requestSubmit(skip);
    });
    await nav;
    await sleep(1000);
  }

  log("waiting for callback → app …");
  try {
    await waitFor(() => page.url().startsWith(APP), "back to app", 30000);
  } catch (e) {
    const dbg = await page.evaluate(() => ({
      url: location.href,
      text: document.body?.innerText?.slice(0, 500),
      inputs: Array.from(document.querySelectorAll("input")).map((i) => ({ name: i.name, type: i.type, value: i.value })),
      btns: Array.from(document.querySelectorAll("button")).map((b) => b.textContent?.trim()),
    }));
    process.stderr.write("[e2e] DEBUG after submit: " + JSON.stringify(dbg) + "\n");
    throw e;
  }
  log(`back on ${page.url()}`);

  // App should show signed-in state + profile from backend.
  await waitFor(() => page.evaluate(() => document.body.innerText.includes("Signed in")), "signed in state", 20000);
  const body = await page.evaluate(() => document.body.innerText);
  log("--- page text ---");
  console.log(body.slice(0, 800));
  log("---");

  if (!body.includes("Profile from Go API")) throw new Error("profile section missing");
  await waitFor(() => page.evaluate(() => !document.body.innerText.includes("Loading profile…")), "profile loaded", 15000);
  const finalBody = await page.evaluate(() => document.body.innerText);
  if (finalBody.includes("API 401") || finalBody.includes("invalid token")) {
    throw new Error("/api/me rejected token:\n" + finalBody.slice(0, 600));
  }
  if (!/zitadel-admin|@zitadel\.localhost|sub/.test(finalBody)) {
    throw new Error("profile did not include user info:\n" + finalBody.slice(0, 600));
  }
  log("E2E OK: login + /api/me profile verified ✔");
} finally {
  await browser.close();
}

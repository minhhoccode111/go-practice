# Header + Footer Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development or superpowers:executing-plans. Steps use checkbox (`- [ ]`) syntax.

**Goal:** Add dark terminal-themed header with nav links and sticky footer to the site layout.

**Architecture:** Two new templ components (`components/header.templ`, `components/footer.templ`) composed into `layout/app.templ`. HTMX handles partial page loads on nav clicks. Alpine handles mobile hamburger toggle.

**Tech Stack:** Go, templ, HTMX 2.0, Alpine.js 3, daisyUI 5, Tailwind CSS v4

---

### Task 1: Create Header Component

**Files:**

- Create: `components/header.templ`

- [ ] **Step 1: Write `components/header.templ`**

```templ
package components

templ Header() {
    <nav class="navbar bg-base-100 shadow-sm sticky top-0 z-50 border-b border-base-300">
        <div class="navbar-start">
            <div class="dropdown" x-data="{ open: false }">
                <div tabindex="0" role="button" class="btn btn-ghost lg:hidden" @click="open = !open">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16"/>
                    </svg>
                </div>
                <ul tabindex="-1" class="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow" x-show="open" @click.outside="open = false">
                    <li><a href="/" hx-get="/" hx-target="#main" hx-push-url="true">Home</a></li>
                    <li><a href="/hello" hx-get="/hello" hx-target="#main" hx-push-url="true">Hello</a></li>
                </ul>
            </div>
            <a href="/" class="btn btn-ghost text-xl tracking-tight font-mono" hx-get="/" hx-target="#main" hx-push-url="true">
                <span class="text-[#3fb950]">&gt;</span> try-htmx
            </a>
        </div>
        <div class="navbar-end hidden lg:flex">
            <ul class="menu menu-horizontal px-1">
                <li><a href="/" hx-get="/" hx-target="#main" hx-push-url="true" class="font-mono">Home</a></li>
                <li><a href="/hello" hx-get="/hello" hx-target="#main" hx-push-url="true" class="font-mono">Hello</a></li>
            </ul>
        </div>
    </nav>
}
```

- [ ] **Step 2: Generate Go code**

Run: `templ generate`
Expected: `components/header_templ.go` created

- [ ] **Step 3: Verify it compiles**

Run: `go build ./...`
Expected: no errors

---

### Task 2: Create Footer Component

**Files:**

- Create: `components/footer.templ`

- [ ] **Step 1: Write `components/footer.templ`**

```templ
package components

templ Footer() {
    <footer class="footer footer-center bg-base-200 text-base-content p-6 border-t border-base-300">
        <aside>
            <p class="font-mono text-sm text-base-content/70">
                Built with <a href="https://go.dev" class="link link-hover text-[#3fb950]" target="_blank">Go</a>
                &amp; <a href="https://htmx.org" class="link link-hover text-[#3fb950]" target="_blank">HTMX</a>
                &mdash; <span class="text-base-content/50">&copy; 2026</span>
            </p>
        </aside>
    </footer>
}
```

- [ ] **Step 2: Generate Go code**

Run: `templ generate`
Expected: `components/footer_templ.go` created

- [ ] **Step 3: Verify it compiles**

Run: `go build ./...`
Expected: no errors

---

### Task 3: Update Layout

**Files:**

- Modify: `layout/app.templ`

- [ ] **Step 1: Update `layout/app.templ`**

Replace the entire file content:

```templ
package layout

type AppData struct {
	PageTitle string
}

templ App(data AppData, content templ.Component) {
	<!DOCTYPE html>
	<html lang="en" data-theme="dark">
		<head>
			<meta charset="UTF-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
			<title>{ data.PageTitle } &mdash; try-htmx</title>
			<link href="/assets/css/output.css" rel="stylesheet"/>
			<script defer src="/assets/js/htmx.min.js"></script>
			<script defer src="/assets/js/alpine.min.js"></script>
		</head>
		<body class="min-h-screen flex flex-col">
			{ components.Header() }
			<main id="main" class="flex-1">
				@content
			</main>
			{ components.Footer() }
		</body>
	</html>
}
```

- [ ] **Step 2: Generate Go code**

Run: `templ generate`
Expected: `layout/app_templ.go` updated

- [ ] **Step 3: Verify it compiles**

Run: `go build ./...`
Expected: no errors

---

### Task 4: Rebuild CSS and Run

- [ ] **Step 1: Rebuild Tailwind CSS**

Run: `./tailwindcss -i assets/css/input.css -o assets/css/output.css`
Expected: no errors, `assets/css/output.css` updated

- [ ] **Step 2: Run the server**

Run: `go run .`
Expected: server starts on port 8080 (or from .env)

- [ ] **Step 3: Verify in browser**

Open `http://localhost:8080`

- Header visible with "> try-htmx" logo, Home and Hello links
- Green accent on logo `>`
- Mobile: hamburger shows on narrow viewport
- Nav links load content via HTMX (no full page reload)
- Footer visible at bottom: "Built with Go & HTMX — © 2026"

---

### Task 5: Cleanup

- [ ] **Step 1: Remove empty class attributes**

In `components/hello.templ`, remove the empty `class=""` attributes:

```templ
package components

templ Hello(name string) {
	<div>Hello, { name }!</div>
	<button class="btn btn-primary">Default</button>
}
```

- [ ] **Step 2: Regenerate and verify**

Run: `templ generate && go build ./...`
Expected: no errors

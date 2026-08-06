import type { UserManagerSettings } from "oidc-client-ts";

const issuer = import.meta.env.VITE_ZITADEL_ISSUER;
const clientId = import.meta.env.VITE_ZITADEL_CLIENT_ID;
const scope = import.meta.env.VITE_ZITADEL_SCOPE;
const redirectUri = import.meta.env.VITE_ZITADEL_REDIRECT_URI;
const postLogoutRedirectUri = import.meta.env
  .VITE_ZITADEL_POST_LOGOUT_REDIRECT_URI;

export const oidcConfig: UserManagerSettings = {
  authority: issuer,
  client_id: clientId,
  redirect_uri: redirectUri,
  post_logout_redirect_uri: postLogoutRedirectUri,
  response_type: "code",
  scope,
  automaticSilentRenew: true,
  filterProtocolClaims: true,
  loadUserInfo: true,
};

export function apiUrl(path: string): string {
  return `${import.meta.env.VITE_API_URL}${path}`;
}

export async function fetchMe(
  accessToken: string,
): Promise<Record<string, unknown>> {
  const res = await fetch(apiUrl("/api/me"), {
    headers: { Authorization: `Bearer ${accessToken}` },
  });
  if (!res.ok) {
    throw new Error(`API ${res.status}: ${await res.text()}`);
  }
  return res.json();
}

import type { User } from "oidc-client-ts";
import type { AuthProviderProps } from "react-oidc-context";

// Required by react-oidc-context: strips code/state from the URL after login.
// Without it, refreshing the callback URL re-runs signinCallback with a
// consumed state and breaks silent renew.
export const onSigninCallback = (_user: User | undefined): void => {
  window.history.replaceState({}, document.title, window.location.pathname);
};

const issuer = import.meta.env.VITE_ZITADEL_ISSUER;
const clientId = import.meta.env.VITE_ZITADEL_CLIENT_ID;
const scope = import.meta.env.VITE_ZITADEL_SCOPE;
const redirectUri = import.meta.env.VITE_ZITADEL_REDIRECT_URI;
const postLogoutRedirectUri = import.meta.env
  .VITE_ZITADEL_POST_LOGOUT_REDIRECT_URI;

export const oidcConfig: AuthProviderProps = {
  authority: issuer,
  client_id: clientId,
  redirect_uri: redirectUri,
  post_logout_redirect_uri: postLogoutRedirectUri,
  response_type: "code",
  scope,
  automaticSilentRenew: true,
  filterProtocolClaims: true,
  loadUserInfo: true,
  onSigninCallback,
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

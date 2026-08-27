import { useCallback, useEffect, useState } from "react";
import { useAuth } from "react-oidc-context";
import { fetchMe } from "./auth";

type UserInfo = Record<string, unknown>;

function useProfile(accessToken: string | undefined) {
  const [info, setInfo] = useState<UserInfo | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!accessToken) {
      setInfo(null);
      setError(null);
      return;
    }
    let cancelled = false;
    fetchMe(accessToken)
      .then((data) => {
        if (!cancelled) setInfo(data);
      })
      .catch((err: Error) => {
        if (!cancelled) setError(err.message);
      });
    return () => {
      cancelled = true;
    };
  }, [accessToken]);

  return { info, error };
}

export default function App() {
  const auth = useAuth();
  const accessToken = auth.user?.access_token;
  const { info, error } = useProfile(accessToken);

  const signIn = useCallback(() => auth.signinRedirect(), [auth]);
  const signOut = useCallback(
    () =>
      auth.signoutRedirect({
        post_logout_redirect_uri: import.meta.env
          .VITE_ZITADEL_POST_LOGOUT_REDIRECT_URI,
      }),
    [auth],
  );

  if (auth.isLoading) {
    return <div className="page">Loading…</div>;
  }

  return (
    <div className="page">
      <h1>React + Go + ZITADEL auth example</h1>

      {auth.error && <p className="error">Auth error: {auth.error.message}</p>}

      {auth.isAuthenticated ? (
        <>
          <p>
            Signed in as{" "}
            <strong>
              {auth.user?.profile.preferred_username ??
                auth.user?.profile.email ??
                auth.user?.profile.sub}
            </strong>
          </p>
          <button onClick={signOut}>Sign out</button>

          <h2>Profile from Go API</h2>
          {error ? (
            <p className="error">{error}</p>
          ) : info ? (
            <pre>{JSON.stringify(info, null, 2)}</pre>
          ) : (
            <p>Loading profile…</p>
          )}
        </>
      ) : (
        <>
          <p>Not signed in.</p>
          <button onClick={signIn}>Sign in</button>
        </>
      )}
    </div>
  );
}

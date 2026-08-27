/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_ZITADEL_ISSUER: string;
  readonly VITE_ZITADEL_CLIENT_ID: string;
  readonly VITE_ZITADEL_SCOPE: string;
  readonly VITE_ZITADEL_REDIRECT_URI: string;
  readonly VITE_ZITADEL_POST_LOGOUT_REDIRECT_URI: string;
  readonly VITE_API_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}

# React + Go + ZITADEL — username/password auth example

A complete local example of OIDC username/password login:

- **React SPA** (Vite + `oidc-client-ts`) — Authorization Code flow with PKCE.
- **Go API** (stdlib + `coreos/go-oidc/v3`) — validates the access token's
  signature/JWKS against ZITADEL and returns the user profile.
- **ZITADEL** + **PostgreSQL** — local infrastructure via Docker Compose.

```
Browser (React :5174)  --PKCE authorize/callback-->  ZITADEL (:8082)
        |                                                 |
        |  Bearer token (JWT)                             |  hosted login UI
        v                                                 v
   Go API (:8083)  --OIDC discovery + JWKS verify-->   Postgres
```

## Prerequisites

- Docker + Docker Compose
- Go 1.22+
- Node 20+

## Quick start

```sh
cp .env.example .env        # default admin password is Password1!
docker compose up -d --wait # starts ZITADEL (:8082) + Postgres
bash scripts/setup.sh       # creates project + OIDC SPA app, writes web/.env and backend/.env

# backend (terminal 1)
cd backend
set -a && . ../backend/.env && set +a
go run ./cmd/server

# frontend (terminal 2)
cd web
npm install
npm run dev                 # http://localhost:5174
```

Open http://localhost:5174, click **Sign in**, and log in with
`zitadel-admin@zitadel.localhost` / `Password1!`.

The admin console lives at http://localhost:8082/ui/console
(login with the same user). Add a human user there and log in with those
credentials in the app.

## How provisioning works

ZITADEL has no built-in "create an app" file, so `scripts/setup.sh` automates it:

1. Reads the first-instance machine-user PAT from `zitadel/admin.pat`
   (written by ZITADEL during `start-from-init`; expires 2099).
2. Calls the ZITADEL v2 API to create the project and the OIDC
   user-agent (SPA) application with the PKCE redirect URIs.
3. Writes `web/.env` (issuer, client id, scope, redirect URIs) and
   `backend/.env` (issuer, project id).

Idempotent: re-running skips if `web/.env` already has a client id.
Delete `web/.env` + `backend/.env` and re-run to provision a fresh instance.

> **Reset everything:** `docker compose down -v` wipes the Postgres volume
> (new instance → new project/app IDs → re-run `scripts/setup.sh`).

## Ports

| Service         | URL                              |
| --------------- | -------------------------------- |
| React SPA       | http://localhost:5174            |
| ZITADEL         | http://localhost:8082            |
| ZITADEL console | http://localhost:8082/ui/console |
| Go API          | http://localhost:8083            |

Change the host ports in `.env` (`ZITADEL_PORT`) and `scripts/setup.sh`.

## Notes

- Login V1 is used (V2 needs a separate `zitadel-login` container).
- The initial admin is forced through one-time TOTP setup on first login;
  use **Skip** there (the e2e test does this).
- Access tokens are JWTs containing `aud = [client_id, project_id]`; the Go
  API validates the project id is in `aud`.

## End-to-end test

With backend + frontend running:

```sh
cd web && node e2e-auth.mjs
```

Drives headless Chrome through the full PKCE login and asserts the Go API
returns the profile.

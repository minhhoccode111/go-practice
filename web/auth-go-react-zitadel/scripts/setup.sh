#!/usr/bin/env bash
# Creates the ZITADEL project + OIDC SPA application and writes frontend/backend env files.
# Idempotent: skips if app already provisioned (web/.env exists with a client id).
set -euo pipefail

ZITADEL_PORT="${ZITADEL_PORT:-8082}"
BASE_URL="http://localhost:${ZITADEL_PORT}"
PAT_FILE="zitadel/admin.pat"
WEB_ENV="web/.env"
BACKEND_ENV="backend/.env"

app_exists() {
  [[ -f "$WEB_ENV" ]] && grep -q '^VITE_ZITADEL_CLIENT_ID=.\+' "$WEB_ENV" 2>/dev/null
}

if app_exists; then
  echo "Provisioning already done ($WEB_ENV exists). Skipping."
  exit 0
fi

# Wait for ZITADEL to be ready.
echo "Waiting for ZITADEL at $BASE_URL ..."
for i in $(seq 1 60); do
  if curl -fsS "$BASE_URL/debug/healthz" >/dev/null 2>&1; then
    break
  fi
  [[ "$i" == 60 ]] && { echo "ZITADEL not healthy after 60s" >&2; exit 1; }
  sleep 2
done

# Read the first-instance machine PAT (root-owned host file written by the container).
echo "Reading machine PAT ..."
PAT=""
if [[ -r "$PAT_FILE" ]]; then
  PAT="$(cat "$PAT_FILE")"
else
  for i in $(seq 1 30); do
    if PAT="$(docker compose exec -T zitadel cat /zitadel/admin.pat 2>/dev/null)"; then
      [[ -n "$PAT" ]] && break
    fi
    [[ "$i" == 30 ]] && { echo "admin.pat not found after 30s" >&2; exit 1; }
    sleep 2
  done
fi
PAT="$(printf '%s' "$PAT" | tr -d '\r\n')"
[[ -z "$PAT" ]] && { echo "Empty PAT" >&2; exit 1; }
echo "PAT acquired."

AUTH=(-H "Authorization: Bearer $PAT" -H "Content-Type: application/json")

post() { # post <url> <json>
  curl -fsS -X POST "$BASE_URL$1" "${AUTH[@]}" -d "$2"
}

echo "Reading org id ..."
ORG_ID="$(curl -fsS "$BASE_URL/management/v1/orgs/me" "${AUTH[@]}" | jq -r .org.id)"
echo "Org: $ORG_ID"

echo "Creating project ..."
PROJECT_ID="$(post "/zitadel.project.v2.ProjectService/CreateProject" "{\"organizationId\":\"$ORG_ID\",\"name\":\"auth-example\"}" | jq -r .projectId)"
echo "Project: $PROJECT_ID"

echo "Creating OIDC app ..."
APP_JSON=$(post "/zitadel.application.v2.ApplicationService/CreateApplication" \
  "{\"projectId\":\"$PROJECT_ID\",\"name\":\"react-spa\",
    \"oidcConfiguration\":{
      \"redirectUris\":[\"http://localhost:5174/auth/callback\"],
      \"postLogoutRedirectUris\":[\"http://localhost:5174\"],
      \"responseTypes\":[\"OIDC_RESPONSE_TYPE_CODE\"],
      \"grantTypes\":[\"OIDC_GRANT_TYPE_AUTHORIZATION_CODE\",\"OIDC_GRANT_TYPE_REFRESH_TOKEN\"],
      \"appType\":\"OIDC_APP_TYPE_USER_AGENT\",
      \"authMethodType\":\"OIDC_AUTH_METHOD_TYPE_NONE\",
      \"version\":\"OIDC_VERSION_1_0\",
      \"devMode\":true,
      \"accessTokenType\":\"OIDC_TOKEN_TYPE_JWT\",
      \"accessTokenRoleAssertion\":true,
      \"idTokenUserinfoAssertion\":true
    }}")
CLIENT_ID="$(printf '%s' "$APP_JSON" | jq -r .oidcConfiguration.clientId)"
echo "App: $CLIENT_ID"

SCOPE="openid profile email urn:zitadel:iam:org:project:id:${PROJECT_ID}:aud"

mkdir -p web backend
cat > "$WEB_ENV" <<EOF
VITE_ZITADEL_ISSUER=$BASE_URL
VITE_ZITADEL_CLIENT_ID=$CLIENT_ID
VITE_ZITADEL_SCOPE=$SCOPE
VITE_ZITADEL_REDIRECT_URI=http://localhost:5174/auth/callback
VITE_ZITADEL_POST_LOGOUT_REDIRECT_URI=http://localhost:5174
VITE_API_URL=http://localhost:8083
EOF

cat > "$BACKEND_ENV" <<EOF
ZITADEL_ISSUER=$BASE_URL
ZITADEL_PROJECT_ID=$PROJECT_ID
PORT=8083
EOF

echo
echo "Provisioned. Wrote $WEB_ENV and $BACKEND_ENV"
echo "Admin console: $BASE_URL/ui/console (login: ${ZITADEL_FIRSTINSTANCE_ORG_HUMAN_USERNAME:-zitadel-admin}@zitadel.localhost / ${ZITADEL_FIRSTINSTANCE_ORG_HUMAN_PASSWORD:-Password1!})"

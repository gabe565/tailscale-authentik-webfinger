# Tailscale Authentik Webfinger

A simple Go app that serves a `/.well-known/webfinger` endpoint for using Tailscale SSO with Authentik.

## Features
- **Configurable issuer URL:** Uses `--domain` and `--app-name` to generate the issuer URL.
- **Dynamic accounts:** Reads the `resources` query param to properly handle multiple accounts.

## Usage

```shell
docker run --detach \
  --env AK_HOST=auth.example.com \
  --env AK_APP_NAME=tailscale \
  --port 3000:3000 \
  ghcr.io/gabe565/tailscale-authentik-webfinger
```

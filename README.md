# containers

Custom container images built and maintained by **JoestarLabs**, published to GitHub Container Registry (GHCR).

## Features

- **Multi-Arch:** Native `linux/amd64` & `linux/arm64` builds (no emulation overhead).
- **Automation:** Rebuilt on push via parallel native runners and merged via Docker manifests.

## Image Registry

- **`caddy-docker-cloudflare`** (`ghcr.io/joestarlabs/caddy-docker-cloudflare:latest`)
  - Custom Caddy with: `caddy-docker-proxy`, `cloudflare-dns`, `maxmind-geolocation`.
- **`plezy-relay`** (`ghcr.io/joestarlabs/plezy-relay:latest`)
  - Lightweight Go media status proxy, derived from [edde746/plezy](https://github.com/edde746/plezy).

# containers

Custom container images built and maintained by **JoestarLabs**, published to GitHub Container Registry (GHCR).

## Features

- **Multi-Arch:** Native `linux/amd64` & `linux/arm64` builds (no emulation overhead).
- **Automation:** Rebuilt on push via parallel native runners and merged via Docker manifests.

## Image Registry

### caddy-docker-cloudflare (`ghcr.io/joestarlabs/caddy-docker-cloudflare:latest`)
- Custom Caddy with: `caddy-docker-proxy`, `cloudflare-dns`, `maxmind-geolocation`.

<details>
<summary>Docker Compose Example</summary>

```yaml
services:
  caddy:
    image: ghcr.io/joestarlabs/caddy-docker-cloudflare:latest
    container_name: caddy
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
      - "443:443/udp"
    environment:
      - CF_API_TOKEN=your_cloudflare_api_token
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - caddy_data:/data
      - caddy_config:/config
    networks:
      - caddy

networks:
  caddy:
    name: caddy

volumes:
  caddy_data:
  caddy_config:
```

</details>

### plezy-relay (`ghcr.io/joestarlabs/plezy-relay:latest`)
- Lightweight Go media status proxy, derived from [edde746/plezy/server](https://github.com/edde746/plezy/tree/main/server).

<details>
<summary>Docker Compose Example</summary>

```yaml
services:
  plezy-relay:
    image: ghcr.io/joestarlabs/plezy-relay:latest
    container_name: plezy-relay
    restart: unless-stopped
    ports:
      - "8080:8080"
    environment:
      - TRUSTED_PROXY_CIDRS=10.0.0.0/8,172.16.0.0/12,192.168.0.0/16
      # - OAUTH_BASE_URL=https://relay.example.com
      # - MAL_CLIENT_ID=your_mal_client_id
      # - ANILIST_CLIENT_ID=your_anilist_client_id
      # - ANILIST_CLIENT_SECRET=your_anilist_client_secret
    volumes:
      - plezy_data:/data

volumes:
  plezy_data:
```

</details>

### og-relay (`ghcr.io/joestarlabs/og-relay:latest`)
- Lightweight Go server serving Open Graph link preview metadata and dynamic card images for Authentik-protected domains.

<details>
<summary>Docker Compose Example</summary>

```yaml
services:
  og-relay:
    image: ghcr.io/joestarlabs/og-relay:latest
    container_name: og-relay
    restart: unless-stopped
    ports:
      - "8080:8080"
```

</details>

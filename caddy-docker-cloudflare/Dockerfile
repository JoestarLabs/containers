# Pin the Caddy version for stability
ARG CADDY_VERSION=2.11.3 # May 12 2026

FROM caddy:${CADDY_VERSION}-builder AS builder

ARG CADDYDOCKERPROXY_VERSION=v2@v2.12.0 # Mar 12 2026
ARG CLOUDFLAREDNS_VERSION=a8737d0 # Mar 23 2026
#ARG CROWDSECBOUNCERHTTP_VERSION=v0.10.1 # Mar 26 2026
ARG MAXMINDGEOLOCATION_VERSION=9066f91 # Mar 5 2025

RUN xcaddy build \
    --with github.com/lucaslorentz/caddy-docker-proxy/${CADDYDOCKERPROXY_VERSION} \
    --with github.com/caddy-dns/cloudflare@${CLOUDFLAREDNS_VERSION} \
#    --with github.com/hslatman/caddy-crowdsec-bouncer/http@${CROWDSECBOUNCERHTTP_VERSION} \
    --with github.com/porech/caddy-maxmind-geolocation@${MAXMINDGEOLOCATION_VERSION}

FROM caddy:${CADDY_VERSION}-alpine

COPY --from=builder /usr/bin/caddy /usr/bin/caddy

CMD ["caddy", "docker-proxy"]

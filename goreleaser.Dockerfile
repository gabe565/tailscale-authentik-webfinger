FROM gcr.io/distroless/static:nonroot
LABEL org.opencontainers.image.source="https://github.com/gabe565/tailscale-authentik-webfinger"
COPY tailscale-authentik-webfinger /
ENTRYPOINT ["/tailscale-authentik-webfinger"]

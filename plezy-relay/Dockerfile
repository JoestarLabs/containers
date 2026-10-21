# Update the readable tag, index digest, and platforms together using the controlled process in CONTRIBUTING.md.
# Platforms: linux/amd64, linux/arm64
FROM golang:1.27.3-alpine3.23@sha256:895e5feab11a4b3774d24ee3c9579028456b82ad37f42a9eb6f82adaf776a201 AS build
# Set by Docker Buildx automatically
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /relay .

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /relay /relay
VOLUME /data
EXPOSE 8080
ENV GOMEMLIMIT=64MiB
ENTRYPOINT ["/relay"]

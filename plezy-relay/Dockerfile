# Update the readable tag, index digest, and platforms together using the controlled process in CONTRIBUTING.md.
# Platforms: linux/amd64, linux/arm64
FROM golang:1.27.1-alpine3.23@sha256:96a6b037cd95ee2d72dc63fec59ad8250110fe795111d783d97aa980ec59ec32 AS build
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

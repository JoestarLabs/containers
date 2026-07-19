FROM golang:1.26-alpine AS build
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

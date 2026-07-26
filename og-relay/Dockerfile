FROM golang:1.26-alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /og-relay main.go

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /og-relay /app/og-relay

EXPOSE 8080

USER 65534:65534

ENTRYPOINT ["/app/og-relay"]

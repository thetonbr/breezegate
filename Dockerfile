# Build stage
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/build/breezegate ./cmd/app

# Runtime stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/build/breezegate /app/breezegate

# Note: config.json should be provided at runtime via volume mount or ConfigMap
# COPY config.json /app/config.json

EXPOSE 80
EXPOSE 443

ENTRYPOINT ["/app/breezegate"]

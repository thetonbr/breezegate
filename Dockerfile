FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/build/breezegate /app/breezegate

COPY config.json /app/config.json

EXPOSE 80
EXPOSE 443

ENTRYPOINT ["/app/breezegate"]

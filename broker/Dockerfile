FROM golang:1.21.1-bookworm as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/broker ./cmd/api
RUN chmod +x /app/broker

FROM alpine:latest as final
RUN apk update && apk upgrade
WORKDIR /app
COPY --chown=1000:1000 --from=builder /app/broker .
CMD ["./broker"]

FROM golang:1.21-bookworm as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o /app/auth ./cmd

FROM alpine:latest as final
WORKDIR /app
COPY --from=builder /app/auth .
CMD [ "./auth" ]

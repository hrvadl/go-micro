FROM golang:1.21.1-bookworm as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/frontend ./cmd/web

FROM alpine:latest as final
WORKDIR /app
COPY --from=builder /app/frontend .
CMD [ "./frontend" ]

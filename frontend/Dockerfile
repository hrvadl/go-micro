FROM golang:1.21.1-bookworm as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/web
CMD [ "./cmd/web/frontend" ]

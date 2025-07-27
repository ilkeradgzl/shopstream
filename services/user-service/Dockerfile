# -----------------------
# Build Stage
# -----------------------
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o user-service ./cmd/server

# -----------------------
# Final Stage
# -----------------------
FROM alpine:3.17
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/user-service .
ENV SERVER_PORT=8081

CMD ["./user-service"]

# Multi-stage Dockerfile for Go backend on Linux ARM / x86_64
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install ca-certificates and git
RUN apk add --no-cache ca-certificates git

# Copy dependency manifests
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# --- Runtime Stage ---
FROM alpine:latest AS runner

WORKDIR /app

# Install runtime dependencies (ca-certificates for HTTPS/TLS, tzdata)
RUN apk add --no-cache ca-certificates tzdata

# Copy binary from builder
COPY --from=builder /app/main /app/main

# Default environment variables
ENV PORT=8080

EXPOSE 8080

CMD ["/app/main"]

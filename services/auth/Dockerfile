FROM golang:1.24-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY ./services/auth/go.mod ./services/auth/go.sum ./
COPY ./contracts/users ./contracts/users

# Download dependencies
RUN go mod download

# Copy source code
COPY ./services/auth .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./cmd/app
FROM alpine:latest
WORKDIR /app
COPY --from=builder /auth-service /app/auth-service
COPY --from=builder /app/config /app/config
ENV APP_ENV=production
ENV CONFIG_PATH=/app/config/prod.yml
EXPOSE 50051
CMD ["/app/auth-service", "--config", "/app/config/prod.yml"]
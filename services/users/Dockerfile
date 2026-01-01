FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Copy go.mod and go.sum for dependency caching
COPY ./services/users/go.mod ./services/users/go.sum ./
COPY ./contracts/users ./contracts/users

# Download dependencies
RUN go mod download

# Copy source code
COPY ./services/users .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs

# Expose port
EXPOSE 50051

# Run binary
CMD ["./main"]

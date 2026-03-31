# Stage 1: Build the binary
FROM golang:1.26.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build for ARM64 (native on your runner)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main .

# Stage 2: Dev environment for Zed
FROM golang:1.26.1-alpine AS dev
WORKDIR /app
# Install git, ssh, and basic tools for Zed
RUN apk add --no-cache git openssh-client ca-certificates
# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy source for development
COPY . .
# Install Zed-friendly tools
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/tools/cmd/goimports@latest
CMD ["sleep", "infinity"]

# Stage 3: Final lightweight image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]

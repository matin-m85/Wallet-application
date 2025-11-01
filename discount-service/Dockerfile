# ==============================
# Dockerfile for discount-service
# ==============================

# 🧱 Stage 1: Build Go binary
FROM docker.arvancloud.ir/golang:1.24.2 AS builder

# 📂 Set working directory
WORKDIR /app

# 📦 Copy go.mod & go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# 🧩 Copy the rest of the source code
COPY . .

# ⚙️ Build static binary compatible with Alpine
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o discount-service ./cmd/main.go

# 🚀 Stage 2: Create lightweight runtime image
FROM docker.arvancloud.ir/alpine:latest

WORKDIR /app

# 📥 Copy binary from builder
COPY --from=builder /app/discount-service .

# ✅ Make it executable
RUN chmod +x /app/discount-service

# 🌐 Expose service port
EXPOSE 8080

# 🏁 Start service
CMD ["./discount-service"]

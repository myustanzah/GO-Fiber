# Gunakan image golang resmi untuk build
FROM golang:1.25 AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum lalu download dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN go build -o main .

# Gunakan image minimal untuk run
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/main .

# Expose port aplikasi
EXPOSE 3000

# Jalankan aplikasi
CMD ["./main"]


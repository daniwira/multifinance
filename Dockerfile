# Build aplikasi Golang
FROM golang:1.19-alpine3.17 AS builder

RUN apk add --update --no-cache ca-certificates git

WORKDIR /app

# Salin file aplikasi Golang dan file lainnya yang diperlukan
COPY . .

# Instal dependensi yang dibutuhkan
RUN go mod download

# Build aplikasi Golang
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/app

# Stage 2: Copy binary yang sudah dibuat ke image yang baru
FROM alpine:latest

WORKDIR /app

# Salin binary yang sudah dibuat dari stage 1
COPY --from=builder /app/app .
COPY --from=builder /app/.env .

# Jalankan aplikasi saat container berjalan
ENTRYPOINT ["./app"]
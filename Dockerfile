# 1. Aşama: Build
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Modül dosyalarını kopyala ve bağımlılıkları indir
COPY go.mod go.sum ./
RUN go mod download

# Uygulama dosyalarını kopyala
COPY . .

# Uygulamayı derle (binary oluştur)
RUN go build -o main .

# 2. Aşama: Çalıştırma
FROM alpine:latest

WORKDIR /app

# CA sertifikalarını yükle (gerekirse)
RUN apk --no-cache add ca-certificates

# Build aşamasından derlenmiş binary dosyasını kopyala
COPY --from=builder /app/main .

# Uygulamayı çalıştır
CMD ["./main"]
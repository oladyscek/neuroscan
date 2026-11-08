# ---------- builder ----------
FROM golang:1.26.1-alpine AS builder
WORKDIR /app

# Сначала только go.mod/go.sum — чтобы кэшировался слой с зависимостями
COPY go.mod go.sum ./
RUN go mod download

# Теперь весь код
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

# ---------- runtime ----------
FROM alpine:3.20
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

# Папка для БД (монтируется volume'ом)
RUN mkdir -p /app/data

EXPOSE 8080
CMD ["./server"]
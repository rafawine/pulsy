# Etapa 1: Compilar binario en una imagen grande
FROM golang:latest AS builder

WORKDIR /app

# Copiar y descargar dependencias primero para aprovechar la cache
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar para Linux sin dependencias del sistema (CGO deshabilitado)
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

# Etapa 2: Crear una imagen liviana
FROM alpine:latest

# Añadir certificados para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar solo el binario desde la etapa anterior
COPY --from=builder /app/app .

CMD ["./app"]
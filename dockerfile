# Usa la imagen oficial de Go 1.23.4 como base para la compilación
FROM golang:1.23.4 AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos necesarios para descargar las dependencias
COPY go.mod go.sum ./

# Descarga las dependencias del proyecto
RUN go mod download

# Copia el resto de los archivos del proyecto al contenedor
COPY . .

# Compila el proyecto, especificando el archivo principal
RUN go build -o server ./cmd/server/main.go

# Crea una imagen ligera para ejecutar el binario compilado
FROM debian:bullseye-slim

# Establece el directorio de trabajo para la imagen final
WORKDIR /app

# Copia el binario desde la etapa de compilación
COPY --from=builder /app/server .

# Expone el puerto en el que se ejecutará la aplicación (ajústalo según tu configuración)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./server"]

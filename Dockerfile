# 1. Establecer la imagen base
FROM golang:1.21-alpine AS builder

# 2. Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# 3. Copiar los archivos go.mod y go.sum para instalar dependencias
COPY go.mod go.sum ./

# 4. Descargar las dependencias del proyecto
RUN go mod download

# 5. Copiar el código fuente
COPY . .

# 6. Compilar la aplicación
RUN go build -o cash_register ./cmd/main.go

# 7. Crear la imagen final pequeña para ejecutar el binario
FROM alpine:latest
WORKDIR /app

# 8. Copiar el binario desde la imagen de construcción
COPY --from=builder /app/cash_register .

# 9. Definir el puerto en el que la aplicación escucha
EXPOSE 8081

# 10. Ejecutar la aplicación
CMD ["./cash_register"]

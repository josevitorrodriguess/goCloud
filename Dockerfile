FROM golang:1.24.3

WORKDIR /app

# Copiar tudo
COPY . .

# Mostrar estrutura
RUN ls -la
RUN ls -la cmd/goCloud/

# Download dependências
RUN go mod download

# Build simples
RUN go build -o goCloud ./cmd/goCloud

# Verificar
RUN ls -la goCloud

EXPOSE 3050

CMD ["./goCloud"]
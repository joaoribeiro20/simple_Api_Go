# Etapa 1: Construção da aplicação Go
FROM golang:1.23 AS builder

WORKDIR /app

# Copiar o go.mod e go.sum para a imagem
COPY go.mod go.sum ./

# Baixar dependências do Go
RUN go mod download

# Copiar o código fonte para dentro do container
COPY . .

# Compilar o binário da aplicação
RUN GOOS=linux GOARCH=amd64 go build -o /simplegoapi ./cmd/main.go

# Etapa 2: Rodar a aplicação Go
FROM debian:bullseye-slim

WORKDIR /root/

# Copiar o binário da aplicação da etapa de construção
COPY --from=builder /simplegoapi .

# Expõe a porta 8080 (alterar conforme necessário)
EXPOSE 8080

# Comando para rodar o binário da aplicação
CMD ["./simplegoapi"]

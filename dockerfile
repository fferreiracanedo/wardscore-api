# Multi-stage Dockerfile para desenvolvimento e produção
FROM golang:1.21-alpine AS base

# Instalar dependências do sistema
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata \
    curl \
    postgresql-client

# Instalar Air para hot reload (versão específica compatível com Go 1.21)
RUN go install github.com/cosmtrek/air@v1.49.0

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de módulo Go
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download && go mod verify

# =============================================================================
# DEVELOPMENT STAGE
# =============================================================================
FROM base AS development

# Copiar todo o código fonte
COPY . .

# Expor porta
EXPOSE 8080

# Comando para desenvolvimento com hot reload
CMD ["air", "-c", ".air.toml"]

# =============================================================================
# PRODUCTION STAGE  
# =============================================================================
FROM base AS production

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main ./cmd/api

# Imagem final mínima para produção
FROM alpine:latest AS final

# Instalar certificados SSL e curl
RUN apk --no-cache add ca-certificates tzdata curl

WORKDIR /root/

# Copiar binário da imagem de build
COPY --from=production /app/main .

# Expor porta
EXPOSE 8080

# Comando para executar
CMD ["./main"]
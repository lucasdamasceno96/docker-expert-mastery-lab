# syntax=docker/dockerfile:1
# STAGE 1: Builder
FROM golang:1.24-alpine AS builder

# Instala certificados necessários para o SQLite e dependências
RUN apk add --no-cache ca-certificates git

WORKDIR /app

# Otimização de Cache: Baixa dependências antes de copiar o código
COPY go.mod go.sum ./
ENV GOPROXY=direct
RUN go mod download

# Copia o resto do código
COPY . .

# Build Estratégico:
# 1. --mount=type=cache: Acelera builds subsequentes mantendo o cache do compilador
# 2. CGO_ENABLED=0: Garante binário estático (sem dependência de libc externa)
# 3. -ldflags="-s -w": Remove debug symbols (reduz ~25% do tamanho)
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" \
    -o /app/api ./cmd/api/main.go

# STAGE 2: Runtime (Hardened Alpine)
FROM alpine:3.21

# Segurança: Nunca rode como root em produção
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Prepara o diretório de dados para o SQLite com permissões corretas
RUN mkdir /data && chown appuser:appgroup /data

# Copia apenas o binário final e os certificados
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder --chown=appuser:appgroup /app/api /app/api

# Troca para usuário não-privilegiado
USER appuser

# Persistência do Banco de Dados
VOLUME /data

EXPOSE 8080

ENTRYPOINT ["/app/api"]
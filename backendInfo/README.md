# 🚀 Backend Info - WardScore API Go

## 📋 **Conteúdo desta pasta**

Esta pasta contém todas as informações necessárias para implementar o backend da aplicação WardScore em **Go**.

### 📁 **Arquivos inclusos**

#### `BACKEND_TASKS.md`

**Lista completa de tarefas** organizadas por fases para desenvolvimento do backend:

- ✅ Configuração base
- ✅ API Core
- ✅ Riot API Integration
- ✅ Features avançadas
- ✅ Performance & Security
- ✅ Deploy & Production

#### `env.example`

**Arquivo de configuração** com todas as variáveis de ambiente necessárias:

- Configuração do servidor Go (porta 8080)
- PostgreSQL database
- JWT authentication
- Riot API credentials
- Redis, Mapbox e outros serviços

#### `wardscore-api-example.md`

**Código completo de exemplo** com:

- Estrutura de pastas recomendada
- Modelos GORM (User, Replay, Analysis)
- Rotas com Gin framework
- Middleware de autenticação
- Health check
- Docker Compose para desenvolvimento

## 🎯 **Como usar**

1. **Leia primeiro** o `BACKEND_TASKS.md` para entender as fases
2. **Configure** o ambiente usando o `env.example`
3. **Implemente** seguindo os exemplos em `wardscore-api-example.md`
4. **Marque** as tasks conforme for completando

## 🔥 **Start rápido**

```bash
# 1. Criar pasta da API (fora do projeto frontend)
mkdir wardscore-api
cd wardscore-api

# 2. Seguir os passos do wardscore-api-example.md
# 3. Marcar tasks no BACKEND_TASKS.md
```

## 📍 **Importante**

- A API Go deve ser criada em uma **pasta separada** do frontend Next.js
- O frontend continuará rodando na porta 3000
- A API Go rodará na porta 8080
- Comunique-se via HTTP entre frontend e backend

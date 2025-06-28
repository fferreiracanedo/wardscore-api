# 🚀 BACKEND TASKS - WardScore (Go API)

## ⚡ **FASE 1: CONFIGURAÇÃO BASE (HOJE)**

### 🔧 **Environment Setup**

- [ ] Criar pasta separada para API Go: `mkdir wardscore-api && cd wardscore-api`
- [ ] Inicializar módulo Go: `go mod init wardscore-api`
- [ ] Criar `.env.example` com todas as variáveis necessárias
- [ ] Criar `.env` para desenvolvimento
- [ ] Instalar godotenv: `go get github.com/joho/godotenv`
- [ ] Configurar estrutura de pastas (cmd, internal, pkg, configs)

### 📦 **Database Setup**

- [ ] Instalar PostgreSQL localmente ou configurar Docker
- [ ] Instalar GORM: `go get gorm.io/gorm gorm.io/driver/postgres`
- [ ] Instalar driver PostgreSQL: `go get github.com/lib/pq`
- [ ] Configurar conexão com banco de dados
- [ ] Criar arquivo de configuração do banco
- [ ] Testar conexão inicial

### 🔐 **Auth Base**

- [ ] Instalar Gin framework: `go get github.com/gin-gonic/gin`
- [ ] Instalar JWT: `go get github.com/golang-jwt/jwt/v5`
- [ ] Instalar bcrypt: `go get golang.org/x/crypto/bcrypt`
- [ ] Criar middleware de autenticação JWT
- [ ] Implementar utils para hash/verificação de senhas
- [ ] Configurar CORS: `go get github.com/gin-contrib/cors`

---

## ⚡ **FASE 2: API CORE (SEMANA 1)**

### 🏗️ **Project Structure**

- [ ] Criar estrutura MVC (models, controllers, services)
- [ ] Configurar roteamento com Gin
- [ ] Implementar middleware básico (logger, recovery, cors)
- [ ] Criar sistema de response padronizado
- [ ] Configurar validação: `go get github.com/go-playground/validator/v10`

### 👤 **User Management**

- [ ] Criar struct User no modelo
- [ ] Implementar `POST /api/users/register` (registro)
- [ ] Implementar `GET /api/users/profile` (perfil do usuário)
- [ ] Implementar `PUT /api/users/profile` (atualizar perfil)
- [ ] Implementar `DELETE /api/users/profile` (deletar usuário)
- [ ] Executar auto-migration para User

### 🎮 **Replay Management**

- [ ] Criar struct Replay no modelo
- [ ] Implementar `POST /api/replays/upload` (upload de replay)
- [ ] Implementar `GET /api/replays` (listar replays do usuário)
- [ ] Implementar `GET /api/replays/:id` (buscar replay específico)
- [ ] Implementar `DELETE /api/replays/:id` (deletar replay)
- [ ] Executar auto-migration para Replay

### 📊 **Analysis System**

- [ ] Criar struct Analysis no modelo
- [ ] Implementar processamento de arquivos .rofl
- [ ] Implementar cálculo real do WardScore
- [ ] Implementar `GET /api/analysis/:id` (buscar análise)
- [ ] Criar sistema de queue com goroutines
- [ ] Executar auto-migration para Analysis

---

## ⚡ **FASE 3: RIOT API INTEGRATION (SEMANA 2)**

### 🔑 **OAuth Riot**

- [ ] Obter credenciais da Riot Developer Portal
- [ ] Implementar `GET /api/auth/riot` (iniciar OAuth)
- [ ] Implementar `POST /api/auth/riot/callback` (callback OAuth)
- [ ] Implementar client HTTP para Riot API
- [ ] Implementar refresh token handling
- [ ] Testar fluxo completo de autenticação

### 🎯 **Riot Data Integration**

- [ ] Implementar `GET /api/riot/summoner/:name` (dados do summoner)
- [ ] Implementar `GET /api/riot/matches/:puuid` (histórico de partidas)
- [ ] Implementar `GET /api/riot/rank/:id` (dados de ranking)
- [ ] Implementar cache Redis: `go get github.com/go-redis/redis/v8`
- [ ] Implementar rate limiting para Riot API
- [ ] Criar sistema de retry para requests

---

## ⚡ **FASE 4: ADVANCED FEATURES (SEMANA 3)**

### 🏆 **Ranking System**

- [ ] Criar struct Ranking no modelo
- [ ] Implementar algoritmo de ranking por WardScore
- [ ] Implementar `GET /api/ranking/global` (ranking global)
- [ ] Implementar `GET /api/ranking/region/:region` (ranking regional)
- [ ] Implementar sistema de ligas/divisões
- [ ] Criar job para atualização periódica de rankings

### 🔍 **Search & Compare**

- [ ] Implementar `GET /api/search/summoner` (buscar summoner)
- [ ] Implementar `GET /api/compare/:id1/:id2` (comparar jogadores)
- [ ] Implementar cache para comparações
- [ ] Implementar `GET /api/compare/history` (histórico de comparações)
- [ ] Otimizar queries para comparações

### 📈 **Analytics**

- [ ] Implementar `GET /api/analytics/user/:id` (estatísticas do usuário)
- [ ] Implementar `GET /api/analytics/trends` (tendências globais)
- [ ] Implementar sistema de métricas avançadas
- [ ] Implementar `GET /api/analytics/dashboard` (dashboard analytics)
- [ ] Criar agregações no banco para performance

---

## ⚡ **FASE 5: PERFORMANCE & SECURITY (SEMANA 4)**

### ⚡ **Performance**

- [ ] Implementar cache Redis para dados frequentes
- [ ] Implementar pagination em todas as listagens
- [ ] Otimizar queries GORM com preloading
- [ ] Implementar compressão gzip: `go get github.com/gin-contrib/gzip`
- [ ] Implementar connection pooling no banco
- [ ] Implementar timeout em requests HTTP

### 🔒 **Security**

- [ ] Implementar rate limiting global: `go get github.com/gin-contrib/rate`
- [ ] Implementar validação robusta de inputs
- [ ] Implementar sanitização de dados
- [ ] Configurar headers de segurança
- [ ] Implementar HTTPS em produção
- [ ] Implementar logging de segurança

### 📝 **Logging & Monitoring**

- [ ] Implementar logging estruturado: `go get github.com/sirupsen/logrus`
- [ ] Implementar health checks: `GET /health`
- [ ] Implementar métricas Prometheus: `go get github.com/prometheus/client_golang`
- [ ] Implementar error tracking
- [ ] Configurar alertas de sistema
- [ ] Implementar profiling para performance

---

## ⚡ **FASE 6: DEPLOY & PRODUCTION (SEMANA 5)**

### 🚀 **Deploy Setup**

- [ ] Criar Dockerfile para containerização
- [ ] Configurar docker-compose para desenvolvimento
- [ ] Configurar deploy no Railway/Render/Heroku
- [ ] Configurar banco de dados em produção
- [ ] Configurar Redis em produção
- [ ] Configurar variáveis de ambiente de produção

### 🔧 **Production Optimization**

- [ ] Implementar build otimizado: `go build -ldflags="-s -w"`
- [ ] Configurar graceful shutdown
- [ ] Implementar backup automático do banco
- [ ] Configurar monitoramento de uptime
- [ ] Implementar CI/CD pipeline
- [ ] Configurar SSL/TLS

---

## 📋 **DAILY CHECKLIST**

### Antes de começar cada dia:

- [ ] Verificar se todas as variáveis de ambiente estão funcionando
- [ ] Executar `go run main.go` e verificar se não há erros
- [ ] Verificar se banco está conectado
- [ ] Testar última funcionalidade implementada

### Antes de finalizar cada dia:

- [ ] Executar `go mod tidy` para limpar dependências
- [ ] Fazer commit das mudanças
- [ ] Atualizar este arquivo marcando tasks concluídas
- [ ] Testar build de produção: `go build`
- [ ] Documentar problemas encontrados

---

## 🔥 **TASKS CRÍTICAS PARA HOJE**

1. **Criar pasta wardscore-api e inicializar módulo Go**
2. **Configurar PostgreSQL**
3. **Instalar e configurar GORM**
4. **Criar primeiro modelo (User)**
5. **Testar conexão com banco**
6. **Criar rota básica de health check**

---

## 📂 **ESTRUTURA DE PASTAS RECOMENDADA**

```
wardscore-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── models/
│   ├── controllers/
│   ├── services/
│   ├── middleware/
│   ├── database/
│   └── utils/
├── pkg/
├── configs/
├── migrations/
├── docker-compose.yml
├── Dockerfile
├── .env.example
├── .env
└── go.mod
```

---

## 📞 **QUANDO PRECISAR DE AJUDA**

- Problemas com GORM: Verificar logs de SQL com `Debug()`
- Problemas com Riot API: Verificar rate limits e credenciais
- Problemas com deployment: Verificar logs do container
- Problemas com performance: Usar `go tool pprof` para profiling
- Problemas com Go modules: Executar `go mod tidy` e `go mod download`

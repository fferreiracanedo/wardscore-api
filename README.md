# WardScore API 🎮

API para análise de replays do League of Legends, focada em avaliar e pontuar o controle de visão dos jogadores.

## 🚀 Tecnologias

- [Go 1.21](https://golang.org/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Docker](https://www.docker.com/)

## 📋 Pré-requisitos

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.21+](https://golang.org/doc/install) (opcional, apenas para desenvolvimento local)

## 🛠️ Instalação

1. Clone o repositório:

```bash
git clone https://github.com/fferreiracanedo/wardscore-api.git
cd wardscore-api
```

2. Copie o arquivo de exemplo de variáveis de ambiente:

```bash
cp backendInfo/env.example .env
```

3. Inicie os containers com Docker Compose:

```bash
docker-compose up -d
```

A API estará disponível em: http://localhost:8080

## 🌐 Endpoints

### Health Check

```
GET /health
```

### Usuários

```
POST   /api/v1/users             - Criar usuário
GET    /api/v1/users             - Listar usuários
GET    /api/v1/users/profile     - Obter perfil
PUT    /api/v1/users/profile     - Atualizar perfil
DELETE /api/v1/users/:id         - Deletar usuário
```

### Replays

```
POST   /api/v1/replays/upload    - Upload de replay
GET    /api/v1/replays           - Listar replays
GET    /api/v1/replays/:id       - Buscar replay
PUT    /api/v1/replays/:id       - Atualizar replay
DELETE /api/v1/replays/:id       - Deletar replay
```

### Análises

```
GET    /api/v1/analysis/:id                 - Buscar análise
POST   /api/v1/analysis/process/:replay_id  - Processar replay
GET    /api/v1/analysis/user/:user_id       - Análises do usuário
```

## 📊 Banco de Dados

### PostgreSQL

- Host: localhost
- Porta: 5432
- Usuário: warduser
- Senha: wardpass123
- Banco: wardscore

### Redis

- Host: localhost
- Porta: 6379

### Adminer (Gerenciador do Banco)

- URL: http://localhost:8081
- Sistema: PostgreSQL
- Servidor: postgres
- Usuário: warduser
- Senha: wardpass123
- Base de dados: wardscore

## 💻 Desenvolvimento

### Estrutura do Projeto

```
.
├── cmd/
│   └── api/              # Ponto de entrada da aplicação
├── internal/
│   ├── config/          # Configurações
│   ├── controllers/     # Controladores HTTP
│   ├── database/        # Conexões com banco de dados
│   ├── middleware/      # Middlewares
│   ├── models/          # Modelos de dados
│   ├── routes/          # Rotas da API
│   ├── services/        # Lógica de negócio
│   └── utils/           # Utilitários
├── docker-compose.yml   # Configuração Docker
└── dockerfile          # Build da aplicação
```

### Hot Reload

O ambiente de desenvolvimento usa Air para hot reload automático. Qualquer alteração nos arquivos Go será detectada e o servidor será reiniciado automaticamente.

## 🔒 Variáveis de Ambiente

Principais variáveis que devem ser configuradas no arquivo `.env`:

```env
# Server
PORT=8080
HOST=0.0.0.0
API_URL=http://localhost:8080

# Database
DATABASE_URL=postgresql://warduser:wardpass123@postgres:5432/wardscore?sslmode=disable

# Redis
REDIS_URL=redis://redis:6379

# JWT
JWT_SECRET=your_jwt_secret_here
JWT_EXPIRES_IN=24h

# Development
DEBUG=true
GIN_MODE=debug
```

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## ✨ Funcionalidades Planejadas

- [ ] Autenticação com conta Riot
- [ ] Análise automática de replays
- [ ] Geração de heatmaps de ward
- [ ] Sistema de ranking
- [ ] Integração com Discord
- [ ] Análise de tendências e métricas
- [ ] Recomendações personalizadas
- [ ] Compartilhamento de análises

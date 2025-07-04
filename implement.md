# GoCloud Storage - Fluxo de Implementação

## 📋 Checklist de Implementação

### 🔧 1. Setup Inicial
- [ ] Inicializar projeto Go (`go mod init`)
- [ ] Instalar dependências (Chi, PostgreSQL, AWS SDK, OAuth2, JWT)
- [ ] Criar estrutura de pastas (`cmd/`, `internal/`, `pkg/`, `config/`)
- [ ] Configurar `.env` e `.env.example`
- [ ] Setup `.gitignore`

### 🗄️ 2. Database & Config
- [ ] Implementar `config/config.go` (carregar env vars)
- [ ] Criar conexão PostgreSQL (`internal/database/`)
- [ ] Criar tabelas (`users`, `files`)
- [ ] Definir models (`internal/models/`)
- [ ] Testar conexão com database

### 🔐 3. Autenticação OAuth
- [ ] Configurar OAuth providers (Google, GitHub)
- [ ] Implementar `internal/auth/oauth.go`
- [ ] Criar handlers OAuth (`/oauth/{provider}`, `/oauth/{provider}/callback`)
- [ ] Implementar JWT service
- [ ] Criar middleware de autenticação
- [ ] Endpoints: `GET /oauth/{provider}`, `GET /oauth/{provider}/callback`, `POST /logout`

### ☁️ 4. AWS S3 Integration
- [ ] Configurar AWS SDK v2
- [ ] Implementar `internal/storage/s3.go`
- [ ] Criar bucket S3 e configurar permissões
- [ ] Testar upload/download básico

### 📁 5. File Management
- [ ] Implementar `internal/file/service.go`
- [ ] Criar handlers de arquivo (`api/files/`)
- [ ] Endpoints: `POST /files/upload`, `GET /files`, `GET /files/{id}`, `DELETE /files/{id}`
- [ ] Validação de arquivos (tipo, tamanho)
- [ ] Associar arquivos aos usuários

### 🌐 6. API Routes
- [ ] Implementar `cmd/go-cloud-storage/main.go`
- [ ] Configurar Chi router
- [ ] Agrupar rotas (`/api/auth/*`, `/api/files/*`)
- [ ] Adicionar middlewares (CORS, Logger, Auth)
- [ ] Tratamento de erros global

### 🔒 7. Security & Validation
- [ ] Implementar rate limiting
- [ ] Validação de input
- [ ] Sanitização de dados
- [ ] Logs de segurança
- [ ] Headers de segurança

### 🧪 8. Testing & Documentation
- [ ] Testes unitários básicos
- [ ] Testar fluxo OAuth completo
- [ ] Testar upload/download de arquivos
- [ ] Documentar endpoints (README ou Swagger)
- [ ] Criar collection Postman/Insomnia

### 🚀 9. Deploy Preparation
- [ ] Configurar para produção
- [ ] Docker/Dockerfile
- [ ] Configurar HTTPS
- [ ] Configurar domínio para OAuth callbacks
- [ ] Backup strategy

## 📊 Fluxo de Dados

```
1. User → OAuth Provider → Callback → JWT Token
2. JWT Token → API Request → Database Query
3. File Upload → S3 Storage → Database Record
4. File Download → Database Query → S3 Retrieval
```

## 🎯 Ordem de Prioridade

1. **Database connection** (sem isso, nada funciona)
2. **OAuth authentication** (base para tudo)
3. **File upload** (funcionalidade principal)
4. **File management** (CRUD completo)
5. **Security & validation** (produção ready)

## ⚡ MVP Mínimo

- [ ] OAuth login (Google)
- [ ] Upload arquivo
- [ ] Listar arquivos
- [ ] Download arquivo
- [ ] Delete arquivo

## 🔄 Fluxo de Desenvolvimento

```
Config → Database → OAuth → S3 → File API → Security → Tests
```

Cada etapa deve estar funcionando antes de avançar para a próxima!
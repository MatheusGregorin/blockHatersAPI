# myMarket

## ✅ Visão geral
`myMarket` é uma API REST em Go que exemplifica padrões básicos de construção de um backend com autenticação via JWT, persistência via GORM/MySQL e estrutura em camadas (handlers, serviços/repositório, modelos e middleware).

O objetivo principal do projeto é demonstrar:
- Uso de **Gin** como framework HTTP.
- Organização em camadas com **interfaces** e **injeção de dependência**.
- Autenticação com **JWT** e proteção de rotas via middleware.
- Persistência com **GORM** e **MySQL**.
- Validação simples de payloads e limpeza de dados sensíveis (ex: senha).

---

## 🧩 Principais recursos / funcionalidades

### 🔐 Autenticação
- Login via endpoint `/login` (retorna token JWT).
- Rotas protegidas com **middleware JWT** (`internal/middleware/jwt.go`).
- Token expira em 24 horas.

### 👤 Gestão de usuários
- Registro de usuário via `/api/v1/register`.
- Busca de usuário por ID via `/api/v1/user/:id`.
- Senhas são armazenadas como **hash bcrypt**.

### 💾 Persistência
- Banco de dados gerenciado via **GORM**.
- Auto-migração das entidades:
  - `models.User`
  - `models.Merchant`
  - `models.Product`
- Modelo pronto para extensão (produto/merchant, mesmo que não exista endpoint hoje).

### 🧱 Estrutura em camadas / Abstrações
- `handler/` — camada de entrada (HTTP). Recebe requests, valida e chama a camada de domínio.
- `internal/interfaces/` — define contratos (interfaces) para repositórios.
- `internal/repository/` — implementa persistência (MySQL via GORM).
- `internal/database/` — configura conexão ao banco e executa migrações.
- `internal/middleware/` — middleware de autenticação JWT.
- `internal/models/` — definições de entidades do domínio.

---

## 🧠 Conceitos usados / arquiteturas demonstradas

### ✅ Injeção de dependência (Dependency Injection)
O `handler` recebe uma interface (`interfaces.IUser`) e não depende diretamente da implementação MySQL.
Isso permite trocar a implementação (por exemplo, PostgreSQL ou mock para testes).

### ✅ Interface + Implementação (Repository Pattern)
A interface `IUser` define as operações esperadas. `repository.UserMysqlRepository` a implementa.

### ✅ Middleware
O middleware em `internal/middleware/jwt.go` intercepta requisições e garante que apenas tokens válidos passem para rotas privadas.

### ✅ Validação de entrada
A validação básica é feita com as tags `binding` do Gin em `handler/user.go` via `ShouldBindJSON`.

### ✅ Segurança básica
- Senhas nunca são retornadas (campo `Password` é limpo antes de enviar resposta).
- Proteção de rotas via token JWT.
- Uso de hashing bcrypt para senhas.

---

## 🧰 Tecnologias / dependências usadas

- **Go 1.25**
- **Gin** (router HTTP) — `github.com/gin-gonic/gin`
- **GORM** (ORM) — `gorm.io/gorm` + `gorm.io/driver/mysql`
- **JWT** — `github.com/golang-jwt/jwt/v5`
- **bcrypt** — `golang.org/x/crypto/bcrypt`
- **dotenv** — `github.com/joho/godotenv`

> Nota: muitas dependências no `go.mod` são indiretas; o projeto utiliza explicitamente as listadas acima.

---

## ▶️ Como rodar

1. Copie o `.env` (já incluído no projeto) e ajuste as variáveis conforme necessário:
   - `DATABASE_URL` (DSN do MySQL)
   - `TOKEN_JWT` (chave secreta para assinar tokens)
   - `REPOSITORY` ("mysql" ou outro tipo no futuro)

2. Execute:

```bash
go run ./cmd
```

O servidor iniciará em `http://localhost:8083`.

---

## 🧪 Endpoints existentes

### Public (sem token)
- `POST /login` — recebe `{ "username", "password" }` e retorna token JWT.

### Protegido (necessita header `Authorization: Bearer <token>`)
- `POST /api/v1/register` — registra novo usuário.
- `GET /api/v1/user/:id` — busca usuário por ID.

---

## 🛠️ Pontos de extensão / próximos passos

- Implementar suporte a PostgreSQL (`repository.NewUserMysqlRepositoryPostgres`).
- Adicionar endpoints CRUD para `Product` e `Merchant`.
- Implementar testes unitários/integração (mocks de repositório).
- Melhorar validação de payloads com *validator* e mensagens específicas.
- Adicionar logging / tracing.

---

## 📌 Estrutura de pastas

```
cmd/             # ponto de entrada (main.go)
handler/         # camada HTTP (controllers)
internal/
  database/      # conexão e migração GORM
  interfaces/    # contratos (interfaces)
  middleware/    # JWT middleware
  models/        # entidades do domínio
  repository/    # implementação de persistência
```

---

Se quiser, posso também gerar um conjunto de testes unitários/integrados ou adicionar endpoints para `Product` e `Merchant`.

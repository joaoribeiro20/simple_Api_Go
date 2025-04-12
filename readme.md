# 📘 Simple Budget API

API simples para gerenciamento de orçamentos (`budgets`) utilizando Go, Gin e GORM.

## 🚀 Tecnologias

- [Go](https://golang.org/) go1.23.8
- [Gin](https://gin-gonic.com/) v1.10.0
- [GORM](https://gorm.io/) v1.25.12
- [driver MySQL] v1.5.7
- [ENV] godotenv v1.5.1

---
## Get start

```
Dcoker 
docker-compose build
docker-compose up -d

Terminal 

go mod init github.com/joaoribeiro20/simplegoapi
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get golang.org/x/text@v0.21.0

go mod tidy

go run ./cmd/main.go


```

## 🧩 Endpoints

---

### 🔹 Criar um Orçamento

- **URL:** `/budgets`
- **Método:** `POST`
- **Body (JSON):**

```json
{
  "name": "Website institucional",
  "descript": "Criação de site para empresa",
  "valor": 1500.50
}

```

- **Resposta (Sucesso - 201):**

```json
json
CopiarEditar
{
  "id": 1,
  "name": "Website institucional",
  "descript": "Criação de site para empresa",
  "valor": 1500.50,
  "created_at": "2025-04-12T10:00:00Z",
  "updated_at": "2025-04-12T10:00:00Z"
}

```

- **Resposta (Erro - 400):**

```json
json
CopiarEditar
{
  "error": "campo X é obrigatório"
}

```

---

### 🔹 Listar Todos os Orçamentos

- **URL:** `/budgets`
- **Método:** `GET`
- **Resposta (Sucesso - 200):**

```json
json
CopiarEditar
[
  {
    "id": 1,
    "name": "Website institucional",
    "descript": "Criação de site para empresa",
    "valor": 1500.50,
    "created_at": "2025-04-12T10:00:00Z",
    "updated_at": "2025-04-12T10:00:00Z"
  },
  {
    "id": 2,
    "name": "Sistema de vendas",
    "descript": "Desenvolvimento de sistema interno",
    "valor": 3200.00,
    "created_at": "2025-04-12T10:10:00Z",
    "updated_at": "2025-04-12T10:10:00Z"
  }
]

```

---

### 🔹 Buscar Orçamento por ID

- **URL:** `/budgets/:id`
- **Método:** `GET`
- **Resposta (Sucesso - 200):**

```json
json
CopiarEditar
{
  "id": 1,
  "name": "Website institucional",
  "descript": "Criação de site para empresa",
  "valor": 1500.50,
  "created_at": "2025-04-12T10:00:00Z",
  "updated_at": "2025-04-12T10:00:00Z"
}

```

- **Resposta (Erro - 404):**

```json
json
CopiarEditar
{
  "error": "Budget não encontrado"
}

```

---

### 🔹 Atualizar Orçamento

- **URL:** `/budgets/:id`
- **Método:** `PUT`
- **Body (JSON):**

```json
json
CopiarEditar
{
  "name": "Site institucional atualizado",
  "descript": "Alteração no escopo",
  "valor": 1800.00
}

```

- **Resposta (Sucesso - 200):**

```json
json
CopiarEditar
{
  "id": 1,
  "name": "Site institucional atualizado",
  "descript": "Alteração no escopo",
  "valor": 1800.00,
  "created_at": "2025-04-12T10:00:00Z",
  "updated_at": "2025-04-12T10:20:00Z"
}

```

- **Resposta (Erro - 400 | 404):**

```json
json
CopiarEditar
{
  "error": "Budget não encontrado"
}

```

ou

```json
json
CopiarEditar
{
  "error": "Erro de validação no body"
}

```

---

### 🔹 Deletar Orçamento

- **URL:** `/budgets/:id`
- **Método:** `DELETE`
- **Resposta (Sucesso - 204):**

```
http
CopiarEditar
(No content)

```

- **Resposta (Erro - 404):**

```json
json
CopiarEditar
{
  "error": "Budget não encontrado"
}

```

---

## 📌 Observações

- Os campos `created_at` e `updated_at` são gerenciados automaticamente pelo GORM.
- O campo `id` é gerado automaticamente.
- Os valores são enviados e recebidos em **JSON**.

---

## 🚀 Tecnologias

- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/)

---

## 🛠 Autor

**João Ribeiro**

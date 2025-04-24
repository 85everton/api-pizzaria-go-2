# 🍕 Pizzaria API

Projeto desenvolvido durante o curso **"Go: Criando uma API REST"** da [Alura](https://www.alura.com.br/), com o objetivo de aplicar os conceitos de construção de APIs utilizando o framework **Gin** em Go.

[![Go](https://img.shields.io/badge/feito%20com-Go-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![Autor](https://img.shields.io/badge/github-@85everton-181717?style=for-the-badge&logo=github)](https://github.com/85everton)

---

## 📋 Descrição

Essa é uma API simples para cadastro, listagem e consulta de pizzas.  
As informações são armazenadas localmente em um arquivo JSON (`dados/pizza.json`).

---

## 🚀 Como rodar o projeto

1. **Clone o repositório:**

```bash
git clone https://github.com/85everton/pizzaria.git
cd pizzaria
```

2. **Execute o projeto:**

```bash
go run main.go
```

A API estará disponível em: `http://localhost:8080`

---

## 🔌 Endpoints disponíveis

### ✅ Listar todas as pizzas
**GET** `/pizzas`

📎 Retorna uma lista com todas as pizzas cadastradas.

---

### 🍕 Cadastrar uma nova pizza
**POST** `/pizzas`

📎 Exemplo de corpo da requisição (JSON):

```json
{
  "nome": "Calabresa",
  "preco": 39.90
}
```

🔁 Retorna a pizza cadastrada com um `id` gerado automaticamente.

---

### 🔍 Buscar pizza por ID
**GET** `/pizzas/:id`

📎 Exemplo: `/pizzas/2`

🔁 Retorna a pizza correspondente ao ID ou uma mensagem de erro.

---

## 💾 Persistência dos dados

As pizzas cadastradas são armazenadas no arquivo:

```
dados/pizza.json
```

Ao reiniciar o servidor, os dados continuam disponíveis.

---

## 🧪 Testando com Thunder Client

Você pode testar todos os endpoints utilizando a extensão **Thunder Client** no VS Code:

1. Instale a extensão pela aba de extensões
2. Crie requisições `GET` e `POST` com os exemplos acima
3. Veja as respostas diretamente na interface do Thunder Client

---

## 🛠 Tecnologias utilizadas

- [Go](https://golang.org/)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- JSON para persistência local

---

## ✍️ Autor

Feito com 💙 por [Everton](https://github.com/85everton) durante os estudos com a Alura.

---

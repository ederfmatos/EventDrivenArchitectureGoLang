## Arquitetura Orientada a Eventos com Golang

[![FullCycle](https://img.shields.io/badge/Plataforma-FullCycle-blue)](https://fullcycle.com.br/)
[![Curso](https://img.shields.io/badge/Curso-Arquitetura%20Orientada%20a%20Eventos-orange)](https://plataforma.fullcycle.com.br/courses/3b8c4f2c-aff9-4399-a72a-ad879e5689a2)
[![GoLang](https://img.shields.io/badge/GoLang-GoLang)](https://go.dev/)

**Descrição**

Este projeto faz parte do curso de Arquitetura Orientada a Eventos da plataforma FullCycle. Ele exemplifica a implementação de uma arquitetura orientada a eventos em Golang, com o uso do Kafka como plataforma de streaming de eventos. A arquitetura proposta compreende diversos componentes, incluindo endpoints REST para interação com o sistema, emissão de eventos para transações realizadas e consumo de eventos para atualização em tempo real da projeção do saldo.

### Componentes do Projeto:

1. **Endpoints REST:**
    - **Criação de cliente:** Permite registrar novos clientes no sistema.
    - **Criação de conta:** Facilita a criação de novas contas associadas a clientes existentes.
    - **Criação de transação entre contas:** Possibilita a realização de transações entre contas.
    - **Busca de saldo de uma conta:** Permite consultar o saldo de uma conta específica.

2. **Emissões de Eventos:**
    - **Publicação de eventos para cada transação realizada:** A cada transação entre contas, um evento é emitido para ser consumido e processado por outros serviços ou componentes da aplicação.

3. **Consumo de Eventos:**
    - **Atualização da projeção do saldo em tempo real:** Eventos emitidos são consumidos para atualizar em tempo real a projeção do saldo das contas envolvidas.

### Exemplo de Requisições com curl:

**Criação de cliente:**

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Joe Doe",
    "email": "joe@gmail.com"
  }' \
  http://localhost:8080/customers
```

**Criação de conta:**

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "a707f6d0-2e13-4eaf-9836-1191e80f8633"
  }' \
  http://localhost:8080/accounts
```

**Criação de transação:**

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "account_id_from": "b23ea375-a3a6-44ee-bdb2-6c2134e7bc8c",
    "account_id_to": "b6bf1c84-4209-4315-b8db-f9a32212264e",
    "amount": 1.00
  }' \
  http://localhost:8080/transactions
```

**Busca de saldo:**

```bash
curl -X GET \
  http://localhost:8080/balances/{accountId}
```

### Instruções para Execução com Docker Compose:

1. Clone este repositório:

```bash
git clone https://github.com/ederfmatos/EventDrivenArchitectureGoLang.git
```

2. Acesse o diretório do projeto:

```bash
cd EventDrivenArchitectureGoLang
```

3. Execute o comando `docker-compose up` para iniciar os containers:

```bash
docker-compose up
```

4. A aplicação estará disponível na porta `8080`.
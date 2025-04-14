
# Payment Gateway

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)](https://go.dev/) [![NestJS](https://img.shields.io/badge/NestJS-%E2%9D%A4-red?logo=nestjs)](https://nestjs.com/) [![Next.js](https://img.shields.io/badge/Next.js-black?logo=next.js)](https://nextjs.org/) [![TailwindCSS](https://img.shields.io/badge/TailwindCSS-38B2AC?logo=tailwind-css&logoColor=white)](https://tailwindcss.com/) [![MySQL](https://img.shields.io/badge/MySQL-00758F?logo=mysql&logoColor=white)](https://www.mysql.com/) [![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)

---

Esse sistema é baseado no curso **Imersão Full Cycle** e foi desenvolvido com fins educacionais. O projeto simula o funcionamento de um **gateway de pagamentos**.

## Tecnologias utilizadas

| Camada             | Tecnologia                      |
|--------------------|----------------------------------|
| Backend principal  | Golang                      |
| Frontend           | Next.js + Tailwind CSS           |
| Microsserviço Anti-Fraude | Nest.js (TypeScript)       |
| Banco de dados     | MySQL                            |
| Containers         | Docker + Docker Compose          |

## Como rodar o projeto

### Pré-requisitos
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Node.js](https://nodejs.org/)
- [Go](https://go.dev/)
- [Nest CLI](https://docs.nestjs.com/cli/overview)

### Passo a passo

1. Clone o repositório:
```
git clone https://github.com/atjulia/payment-gateway.git
cd payment-gateway
```
2.  Suba os containers:    
```
docker-compose up --build
``` 
3.  Acesse os serviços:

-   Frontend: [http://localhost:3000](http://localhost:3000)
    
-   API Gateway (Go): [http://localhost:8000](http://localhost:8000)
    
-   Microsserviço Anti-Fraude (Nest.js): [http://localhost:3001](http://localhost:3001)
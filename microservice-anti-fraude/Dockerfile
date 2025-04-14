# Use Node 20 LTS com Alpine (mais estável e leve)
FROM node:20-alpine

# Instala dependências necessárias (apk ao invés de apt)
RUN apk add --no-cache openssl procps

# Instala o NestJS CLI globalmente
RUN npm install -g @nestjs/cli@11.0.6

# Cria e define o usuário e diretório de trabalho
WORKDIR /usr/src/app

# Copia os arquivos de dependência primeiro para aproveitar cache Docker
COPY package*.json ./

# Instala dependências do projeto
RUN npm install

# Copia o restante dos arquivos
COPY . .

# Expõe a porta do NestJS
EXPOSE 3000

# Comando para iniciar a aplicação em desenvolvimento
CMD ["npm", "run", "start:dev"]
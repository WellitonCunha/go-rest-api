# Minha API REST em Go

Esta é uma API REST construída em Go. A API fornece várias funcionalidades e utiliza diversas bibliotecas para suportar o roteamento, middleware, conexão com banco de dados e outras utilidades.

## Funcionalidades

- [x] Roteamento com Gorilla Mux
- [x] Middleware para CORS
- [x] Conexão com PostgreSQL usando GORM
- [x] Migrações de banco de dados
- [x] Tratamento de erros e respostas JSON

## Dependências

A API REST foi construída utilizando as seguintes bibliotecas:

- [github.com/felixge/httpsnoop](https://github.com/felixge/httpsnoop) v1.0.3
- [github.com/gorilla/handlers](https://github.com/gorilla/handlers) v1.5.2
- [github.com/gorilla/mux](https://github.com/gorilla/mux) v1.8.1
- [github.com/jackc/pgpassfile](https://github.com/jackc/pgpassfile) v1.0.0
- [github.com/jackc/pgservicefile](https://github.com/jackc/pgservicefile) v0.0.0-20221227161230-091c0ba34f0a
- [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) v5.5.5
- [github.com/jackc/puddle/v2](https://github.com/jackc/puddle) v2.2.1
- [github.com/jinzhu/inflection](https://github.com/jinzhu/inflection) v1.0.0
- [github.com/jinzhu/now](https://github.com/jinzhu/now) v1.1.5
- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) v0.17.0
- [golang.org/x/sync](https://pkg.go.dev/golang.org/x/sync) v0.1.0
- [golang.org/x/text](https://pkg.go.dev/golang.org/x/text) v0.14.0
- [gorm.io/driver/postgres](https://gorm.io/docs/driver_postgres.html) v1.5.9
- [gorm.io/gorm](https://gorm.io) v1.25.10

Essas bibliotecas foram utilizadas para fornecer funcionalidades como roteamento, manipulação de middleware, suporte a PostgreSQL e outras utilidades necessárias para a construção da API.

## Pré-requisitos

- Go 1.18+
- PostgreSQL
- Docker e Docker Compose

## Configuração

1. Clone o repositório:

   ```bash
   git clone https://github.com/WellitonCunha/go-rest-api.git
   cd go-rest-api

2. Instale as dependências:
   go mod download

3. Para rodar a aplicação localmente, use:
   go run main.go

A API estará disponível em http://localhost:8000.

## Observações pessoais

## comando para visualizar o ip local no docker
1. Passo: docker-compose exec postgres sh
    hostname -i
    resultado: 172.18.0.2
2. Passo: docker inpect 461 | grep IPAddress
#bibliotecas utilizadas
1. gorilla mux - para rotas
2. gorm orm - manipulação de banco de dados
#atualizar o gorm orm
comando: go get -u gorm.io/gorm
#drive do postgres
comando: go get gorm.io/driver/postgres
#comando para parar todos os containers
docker stop $(docker ps -q)
#tive um problema ao expor a porta 5432 do postgres que fica exposta 
    ##para fora onde já estava sendo usado então refinir outra porta 5485
Instalação do github.com/gorilla/handlers
#biblioteca para facilitar a implementação de política de CORS

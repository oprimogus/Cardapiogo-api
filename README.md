# Cardapiogo-api
Repositório da API do Cardápio em Golang

## Sobre o projeto
![img](/docs/assets/banner.svg)
- Consiste numa API para delivery de pedidos

## Depêndencias
### 1. Migrate CLI
Configure a CLI do [migrate](https://github.com/golang-migrate/migrate/tree/v4.16.2/cmd/migrate) 

### 2. SQLC
Configure o SQLC através deste [link](https://docs.sqlc.dev/en/stable/overview/install.html)

### 3. swaggo
Configure o swaggo através deste [link](https://github.com/swaggo/swag)

### 4. gin-swagger
Configure o swaggo através deste [link](https://github.com/swaggo/gin-swagger)


## Rodar o app localmente

1. Criar .env
    ```
    cp .env.example .env
    ```

2. Instalar dependências do app
    ```
    make install
    ```

3. Subir banco de dados e demais containers
    ```
    make dev-docker-up
    ```

4. Executar migrations
    ```
    make migration-up
    ```
    
5. Mockar dados no banco local
    ```
    make mock-database
    ```

6. Acessar rotas em  http://localhost/api/v1/swagger/index.html#/

# Cardapio-go
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

1. Instalar dependências do app
    ```
    make install
    ```

2. Subir banco de dados e demais containers
    ```
    make docker
    ```
    
3. Rodar o app
    ```
    make run
    ```

4. Acessar rotas em  http://localhost:8080/api/v1/swagger/index.html#/

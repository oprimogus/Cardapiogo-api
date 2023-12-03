# Cardapio-go
Repositório da API do Cardápio em Golang


## Depêndencias
### 1. Migrate CLI
Configure a CLI do [migrate](https://github.com/golang-migrate/migrate/tree/v4.16.2/cmd/migrate) 

### 2. SQLC
Configure o SQLC através deste [link](https://docs.sqlc.dev/en/stable/overview/install.html)


## Comandos

1. Instalar dependências do app
```
make install
```

2. Criar migration
```
make migration
```
3. Rodar o app
```
make run
```
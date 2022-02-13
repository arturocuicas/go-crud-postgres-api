# go-crud-postgres-api

## UUID library:
google/uuid

`go get github.com/google/uuid`

## HTTP Handler
gorilla/mux

`go get -u github.com/gorilla/mux`

## ORM Handler
gorm.io/gorm

`go get -u gorm.io/driver/postgres`

`go get -u gorm.io/gorm`

## Compiler Handler
Very simple Compile Daemon for Go

`CompileDaemon -command="main.exe"`

## Postgres Docker

```dockerfile
docker run --name gorm -e POSTGRES_PASSWORD=gorm -d gorm --ports=9920
```
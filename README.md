# CRUD Приложение коротких данных о языках программирования

### Стэк
- go 1.17
- postgres 

### Запуск
```go run cmd/main.go```

Для postgres можно использовать Docker

```go
docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres
```

```go
docker exec -it postgres-container createdb --username=root --owner=root crud-languages
```

```go
migrate -path ./schema -database "postgresql://root:root@localhost:5432/crud-languages?sslmode=disable" -verbose up
```

### Swagger UI
```http://localhost:8080/swagger/index.html#/```
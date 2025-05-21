# go-study-backend

go study backend

## TODO

- server init
- mysql connect
- REST API
  - GET
  - POST
  - PUT
  - DELETE

## bash

```bash
go run .
go test ./...
```

## DB Connect

```bash
export MYSQL_DATABASE=go_study && export MYSQL_USER=go_user && export MYSQL_PASSWORD=go_password && export MYSQL_ROOT_PASSWORD=root_password && docker-compose -f docker-compose.test.yml up -d

go run .
```

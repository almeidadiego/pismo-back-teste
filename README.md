# pismo-back-teste

## About

This project implements the endpoints considering the business rules mentioned in the Pismo test instructions.

## To run on your developer machine

### Installing dependencies

- docker and docker-compose

### Start services

In the root directory:

```sh
docker-compose up
```

### To execute unit tests

In the root directory:

```sh
go test ./...
```

### Example of use

#### Creating an account

```sh
curl -X POST -d '{"document":"11111111111"}' -H 'Content-Type: application/json' localhost:8080/api/v1/accounts
```

#### Retrieving an account

```sh
curl -X GET http://localhost:8080/api/v1/accounts/1
```

#### Creating a transaction

```sh
curl -X POST -d '{"account_id": 1, "operation_type_id": 3, "amount": 53.75}' -H 'Content-Type: application/json' localhost:8080/api/v1/transactions
```

### Observações:

The _sql-scripts_ directory is used to create schema and insert operation types in mysql database container.

# E-commerce REST API

This project is a REST API for an e-commerce system developed in Go. The API provides endpoints for user management, product handling, and checkout processes.

## Requirements

Before starting the project, make sure you have installed:

- [Go](https://go.dev/doc/install)
- [Docker](https://www.docker.com/get-started) (to run the database)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) (for database migrations)

## Environment Setup

Create a `.env` file based on the template below and fill it with the correct settings:

```
PUBLIC_HOST=http://localhost
PORT=8080
DB_USER=ecom
DB_PASS=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=ecom
JWT_EXPIRATION_SECONDS=604800 # 1 week
JWT_SECRET=secret
``` 

## Available Commands

### Run the Project

```sh
make run
```

### Build the Project

```sh
make build
```

### Run Tests

```sh
make test
```

## Database Migrations

1. Ensure the database instance is running before executing migrations.
2. Install the `golang-migrate` CLI tool if you haven't already.

### Create a New Migration

```sh
make migration <description>
```

### Apply Migration

```sh
make migrate-up
```

### Revert Migration

```sh
make migrate-down
```

## API Endpoints

### User Registration

**POST** `/api/v1/register`

```json
{
    "email": "user@test.com",
    "password": "test",
    "firstName": "User",
    "lastName": "Test"
}
```

### User Login

**POST** `/api/v1/login`

```json
{
    "email": "user@test.com",
    "password": "test"
}
```

### Create a New Product

**POST** `/api/v1/products`

```json
{
    "name": "Product 2",
    "description": "Description of product 2",
    "image": "https://via.placeholder.com/150",
    "price": 50,
    "quantity": 10
}
```

### List Products

**GET** `/api/v1/products`

### Cart Checkout

**POST** `/api/v1/cart/checkout`

**Headers:**
```
Authorization: Bearer <token>
```

```json
{
    "items": [
        {
            "productId": 2,
            "quantity": 2
        }
    ]
}
```

## Authentication

The API uses JWT (JSON Web Token) authentication. To access protected endpoints, provide a valid token in the `Authorization` header as `Bearer <token>`.

## Contribution

1. Fork the repository.
2. Create a branch for your feature (`git checkout -b my-feature`).
3. Commit your changes (`git commit -m 'My new feature'`).
4. Push to the branch (`git push origin my-feature`).
5. Open a Pull Request.

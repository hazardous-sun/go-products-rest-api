# Go REST API with Gin and PostgreSQL

This project is a simple REST API built using the Gin framework in Go. It interacts with a PostgreSQL database to manage
products. The API allows you to create, list, and retrieve products by their ID. The application runs inside Docker
containers, with one container for the Go application and another for the PostgreSQL database.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
    - [Environment Variables](#environment-variables)
    - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Docker Compose](#docker-compose)
- [Database Initialization](#database-initialization)
- [Contributing](#contributing)
- [License](#license)

## Features

- Create a new product
- List all products
- Retrieve a product by ID
- Dockerized environment for easy setup
- PostgreSQL database for data storage

## Prerequisites

Before you begin, ensure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install) (optional, for local development)

## Getting Started

### Environment Variables

The application requires the following environment variables to be set:

- `DB_HOST`: The hostname of the PostgreSQL database (default: `go-db`).
- `DB_PORT`: The port of the PostgreSQL database (default: `5432`).
- `DB_USER`: The username for the PostgreSQL database (default: `postgres`).
- `DB_PASSWORD`: The password for the PostgreSQL database (default: `1234`).
- `DB_NAME`: The name of the PostgreSQL database (default: `postgres`).

These variables are set in the `run.sh` script.

### Running the Application

```bash
# Clone the repository
git clone https://github.com/hazardous-sun/go-rest-api.git
cd go-rest-api

# Make the run.sh script executable:
chmod +x run.sh

# 3. Run the application:
./run.sh
```

#### The [run.sh](./run.sh) script will:

1. Set the necessary environment variables.
2. Clean up any previously created containers.
3. Start the Docker containers using podman-compose.

Finally, the Go application will be available at http://localhost:8000

## API Endpoints

- Create a new product
    - `POST /products`
    - Request Body:
      ```json
      {
        "product_name": "Sample Product",
        "price": 19.99
      }
      ```
    - Response:
      ```json
      {
        "id": 1,
        "product_name": "Sample Product",
        "price": 19.99
      }
      ```
- List all products
    - `GET /products`
    - Response:
      ```json
      [
        {
          "id": 1,
            "product_name": "Sample Product 1",
            "price": 19.99
        },
        {
          "id": 2,
            "product_name": "Sample Product 2",
            "price": 29.99
        }
      ]
      ```
- Retrieve a product by ID:
    - `GET /products/:id`
    - Response:
      ```json
      {
        "id": 1,
        "product_name": "Sample Product",
        "price": 19.99
      }
      ```

## Project Structure

The project is structured into four layers:

1. Controller: Handles HTTP requests and responses.
2. Use Case: Contains the business logic.
3. Repository: Manages data access and interaction with the database.
4. Model: Defines the data models and database schema.

## Docker Compose

The `docker-compose.yml` file defines two services:

- **go-app**: The Go application server.
- **go-db**: The PostgreSQL database.

The `go-app` service depends on the `go-db` service, ensuring that the database is up and running before the application
starts.

```postgresql
CREATE TABLE products
(
    ID           SERIAL PRIMARY KEY,
    PRODUCT_NAME VARCHAR(255) UNIQUE NOT NULL,
    PRICE        DECIMAL(10, 2)      NOT NULL
);
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
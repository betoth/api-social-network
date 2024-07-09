# API Social Network

## Description

This project is an API for a social network, built with Go and following the hexagonal architecture. The API allows creating users, authentication, making posts, following other users, and more.

## Project Structure

- `cmd/api`: Contains the main entry point for the application.
- `internal/application`: Contains the services and ports (interfaces) of the application.
- `internal/domain`: Contains the domain entities.
- `internal/infrastructure`: Contains configuration, middleware, repositories, and other infrastructure code.
- `internal/docs`: Contains the Swagger documentation for the API.

## Technologies Used

- Go
- GORM
- PostgreSQL
- JWT

## How to Run

### Prerequisites

- Go installed
- PostgreSQL configured and running

### Steps

1. Clone the repository:
    ```sh
    git clone https://github.com/your-username/api-social-network.git
    cd api-social-network
    ```

2. Set up the `.env` file with the required environment variables:
    ```
    DB_HOST=your_database_host
    DB_PORT=your_database_port
    DB_USER=your_database_user
    DB_PASSWORD=your_database_password
    DB_NAME=your_database_name
    JWT_SECRET=your_jwt_secret
    ```

3. Install dependencies:
    ```sh
    go mod download
    ```

4. Generate Swagger documentation:
    ```sh
    swag init -g cmd/api/main.go
    ```

5. Run the application:
    ```sh
    make run
    ```

6. Access the Swagger documentation at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Testing

To run the tests, use the command:
```sh
make test

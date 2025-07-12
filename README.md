# Golang Gin Gorm PORTOFOLIO-RESTAPI

## Introduction 
> Clean Architecture is a software design approach that emphasizes the separation of concerns and the clear organization of dependencies between components. In the context of Go (Golang), it refers to applying these principles to build well-structured, maintainable applications.

This approach organizes an application into distinct layers, each with a specific role and responsibility. These layers are designed to be independent from one another, allowing for greater modularity, testability, and adaptability as the application evolves.

Below are the common layers typically found in a Go application following the Clean Architecture principles:

## Layers in Clean Architecture 
- **Initializers**: Handles the initialization and configuration of environment variables and the database connection for the RESTful API.

- **Constants**: Serves as a constant variable that can be used across packages.

- **Middleware**: Provides reusable components that enforce security and access control in the RESTful API application. These middlewares ensure that only authorized users can access specific resources or perform certain actions.

- **Controller**: Provides handler functions for user-related operations and interact with the database in a RESTful API built using the Gin framework.

- **Service**: Services encapsulate the core business logic of the application. They receive instructions from controllers, perform necessary operations—such as calculations or data retrieval—and return results. This separation helps keep controllers thin and focused on request handling, while the services manage the actual functionality behind each operation.

- **Repository**: The repository layer handles all interactions with the data storage system, such as databases or external data sources. It acts as an abstraction layer that shields services from the technical details of how data is stored or retrieved. This separation improves modularity and makes the application easier to test and maintain.

- **Utils**: Utilities consist of helper functions or shared tools used across the application—like input validation, string formatting, error handling, or security checks. Keeping these in a centralized utility layer prevents code duplication and enhances code readability and reusability.

## Web Flow 
The program flow goes through the layers below :

`request` > router > controller > service > repository > service > controller > router > `response`

## Database schema


## Prerequisite 
- Go Version `>= go 1.20`
- PostgreSQL Version `>= version 15.0`

## How To Use
1. Clone the repository
  ```bash
  git clone https://github.com/Shabrinashsf/PORTOFOLIO-RESTAPI.git
  ```
2. Navigate to the project directory:
  ```bash
  cd PORTOFOLIO-RESTAPI
  ```
3. Copy the example environment file and configure it:
  ```bash 
  cp .env.example .env
  ```

There are 2 ways to do running
### With Docker
comming soon

### Without Docker
1. Configure `.env` with your PostgreSQL credentials:
  ```bash
  DB_HOST=localhost
  DB_USER=postgres
  DB_PASS=
  DB_NAME=
  DB_PORT=5432
  ```
2. Open the terminal and follow these steps:
  - If you haven't downloaded PostgreSQL, download it first.
  - Run:
    ```bash
    psql -U postgres
    ```
  - Create the database according to what you put in `.env` => if using uuid-ossp or auto generate (check file **/entity/user_entity.go**):
    ```bash
    CREATE DATABASE your_database;
    \c your_database
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; // remove default:uuid_generate_v4() if you not use you can uncomment code in user_entity.go
    \q
    ``` 
3. Run the application:
  ```bash
  go run main.go
  ```

## Run Migrations and Seeder
To run migrations and seed the database, use the following commands:

```bash
go run main.go migrate seed
```

#### Migrate Database 
To migrate the database schema 
```bash
go run main.go migrate
```
This command will execute all outstanding migrations on your PostgreSQL database as configured in the `.env` file.

#### Seeder Database 
To seed the database with initial data:
```bash
go run main.go seed
```
This command inserts initial data into the database using the seeders configured within your application.

### Postman Documentation
You can explore the available endpoints and their usage in the [Postman Documentation](https://www.postman.com/shabresf/workspace/my-projects/collection/38942886-ed607f3f-5c8c-4cb5-9bf6-2d928b52f023?action=share&creator=38942886). This documentation provides a comprehensive overview of the API endpoints, including request and response examples, making it easier to understand how to interact with the API.
# account-transactions

A dockerized Go web service managing account-related transaction endpoints, created as an interview assignment for Pismo. Focuses on transaction processing, querying, and management, adhering to clean architecture principles and well-structured code.

## Table of contents

1. [Architecture](#%EF%B8%8F-architecture)
2. [Repository Structure](#-repository-structure)
3. [Api-spec](#-api-spec)
4. [Getting Started](#%EF%B8%8F%EF%B8%8F-getting-started)
5. [Docker Image](#-docker-image)
6. [Links](#-links)

## 🏛️ Architecture
The **account-transactions** project follows a clean architecture approach, promoting modularity, maintainability, and separation of concerns. The architecture is designed with flexibility and scalability in mind. Key aspects include:

- **cmd** - Contains the entry point for the application, where the application is initialized and started, setting up routing and server configurations.

- **internal** - Houses the core logic of the application, such as middleware, validation, logging, and system signal handling, which are essential to the internal operations of the project.

- **pkg** - Contains reusable packages for the application, including API handlers and database repositories, enabling modular interaction with business logic and data.

- **Modular and Loosely Coupled Components** - The handler and repository packages are fully modular and loosely coupled, leveraging interfaces in Go to define clear boundaries between the layers of the application. This allows for easy testing, replacement, or extension of components without affecting the overall system.

    - **handler** - Contains the API layer, responsible for handling HTTP requests related to account and transaction operations. Each handler is designed to be independent, and by using interfaces, handlers are easily replaceable or extendable.

    - **repository** - Manages database interactions through repository patterns, isolating database logic from the rest of the application. The use of interfaces in the repository layer ensures flexibility, allowing for different database implementations without changing the business logic.

- **schema** - Contains the database schema and migration files for managing the database structure.

This architecture provides flexibility, making it easy to scale and extend the project while maintaining a clean separation of concerns across the various layers.

## 🗼 Repository structure
```
account-transactions
├───cmd
│   └───account-transactions
├───internal
│   ├───logger
│   ├───middleware
│   ├───migrator
│   ├───mocks
│   ├───signal
│   ├───validator
│   └───writter
├───pkg
│   ├───handler
│   │   ├───accounts
│   │   └───transactions
│   └───repository
├───schema
│   └───migrations
```

## 🚀 Api-spec

This documentation includes everything you need to get started with the API, including detailed endpoint descriptions, parameter definitions, response examples, and error handling information.

You can view the full API specification [here](https://aswinudhayakumar.github.io/account-transactions-swagger/).

## 🏃‍♂️‍➡️ Getting started
### Prerequisites

- Docker
- Docker Compose

### Steps
1. **Clone the Repository**

   ```
   git clone https://github.com/aswinudhayakumar/account-transactions.git
   cd account-transactions
   ```

2. **Start the service**

   ```
   # this builds the docker image and runs the application along with db
   make run
   ```

3. **Run unit tests**
    ```
    make uni-test
    ```

4. **Stopping the Services: To stop and remove the containers, run:**

   ```
   make down
   ```

## ⚓ Docker Image

You can pull the Docker image for this application from Docker Hub. Simply use the following command:

```
docker pull aswin8799/account-transactions
```

For more details and access to the repository, visit the Docker Hub page [here](https://hub.docker.com/repository/docker/aswin8799/account-transactions).

This image contains the fully packaged application, ready to be deployed and run in any Docker-enabled environment.

### Env Variables

The following environment variables are required for configuring the application:

- `APP_PORT` - The port on which the application will run.
- `DB_USER` - The username for connecting to the database.
- `DB_PASSWORD` - The password for the database user.
- `DB_NAME` - The name of the database to use.
- `DB_HOST` - The hostname of the database server.
- `DB_PORT` - The port on which the database is running.
- `SSL_MODE` - SSL connection mode to the database. This can be set to disable, require, etc., depending on your database configuration.
These environment variables must be set in a .env file or configured directly in your system to ensure proper connectivity and behavior of the application.

## 🔗 Links
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/aswin-udhayakumar/)



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
- **Golang with Chi Router** - The application is built using Go (Golang) for its high performance and simplicity. The Chi router is used to handle HTTP routing, providing a lightweight and flexible way to define routes and middleware for the application.

- **Modular Architecture** - The application follows a modular architecture, where components such as handlers, repositories, and middleware are decoupled and interact through well-defined interfaces. This structure ensures flexibility and easy maintenance.

- **Dockerized Application** - The application is containerized using Docker, ensuring portability, ease of deployment, and isolation of dependencies. This allows the application to run consistently across different environments.

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
│   └───writer
├───pkg
│   ├───handler
│   │   ├───accounts
│   │   └───transactions
│   └───repository
├───schema
│   └───migrations
```

- **cmd** - Contains the entry point for the application, where the application is initialized and started, setting up routing and server configurations.

- **internal** - Houses the core logic of the application, such as middleware, validation, logging, and system signal handling, which are essential to the internal operations of the project.

- **pkg** - Contains reusable packages for the application, including API handlers and database repositories, enabling modular interaction with business logic and data.

- **Modular and Loosely Coupled Components** - The handler and repository packages are fully modular and loosely coupled, leveraging interfaces in Go to define clear boundaries between the layers of the application. This allows for easy testing, replacement, or extension of components without affecting the overall system.

    - **handler** - Contains the API layer, responsible for handling HTTP requests related to account and transaction operations. Each handler is designed to be independent, and by using interfaces, handlers are easily replaceable or extendable.

    - **repository** - Manages database interactions through repository patterns, isolating database logic from the rest of the application. The use of interfaces in the repository layer ensures flexibility, allowing for different database implementations without changing the business logic.

- **schema** - Contains the database schema and migration files for managing the database structure.

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
   # This builds the docker image and runs the application along with the db. 
   # By default, application will run at the port 8080.

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
- `SHUTDOWN_TIMEOUT` (Optional): Specifies the duration (in seconds) the application waits before forcefully terminating processes during shutdown. Default is 5 seconds.

These environment variables must be set in a .env file or configured directly in your system to ensure proper connectivity and behavior of the application.

## 🔗 Links
[![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)](https://aswinudhayakumar.github.io/account-transactions-swagger/)
[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/repository/docker/aswin8799/account-transactions)
[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://www.postman.com/aswinudhayakumar/aswin-general/collection/lc45chu/account-transactions)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/aswin-udhayakumar/)



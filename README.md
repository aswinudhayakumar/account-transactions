# account-transactions

This repository contains a project for managing account-related transaction endpoints. Built as part of an interview assignment for PISMO, the project focuses on handling transaction processing, querying, and management with clean architecture and well-structured code.

## Table of contents

1. [Architecture](#%EF%B8%8F-architecture)
2. [Repository structure](#-repository-structure)
3. [Api-spec](#-api-spec)
4. [Getting started](#%EF%B8%8F%EF%B8%8F-getting-started)
5. [Docker image](#-docker-image)
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

API-spec will be added.

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

## ⚓ Docker image

- Can get the docker image of this application here - https://hub.docker.com/repository/docker/aswin8799/account-transactions

## 🔗 Links
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/aswin-udhayakumar/)



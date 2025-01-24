# account-transactions

Result - [screen recording](https://drive.google.com/file/d/1fHCDtGmiDtNpqSoydimtLE1cjbQUV9JD/view?usp=sharing)

[![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)](https://aswinudhayakumar.github.io/account-transactions-swagger/)
[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/repository/docker/aswin8799/account-transactions)
[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://www.postman.com/aswinudhayakumar/aswin-general/collection/lc45chu/account-transactions)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/aswin-udhayakumar/)

A dockerized Go web service managing account-related transaction endpoints, created as an interview assignment for Pismo. Focuses on transaction processing, querying, and management, adhering to clean architecture principles and well-structured code.

## Table of contents

1. [Architecture](#%EF%B8%8F-architecture)
2. [Repository Structure](#-repository-structure)
3. [Api-spec](#-api-spec)
4. [Getting Started](#%EF%B8%8F%EF%B8%8F-getting-started)

## 🏛️ Architecture
- Built with Go (Golang) for high performance and simplicity, using the Chi router for lightweight, flexible HTTP routing and middleware handling.
- Follows a modular architecture with decoupled components (handlers, repositories, middleware) for flexibility and easy maintenance.
- Containerized with Docker for portability, consistent deployment, and isolation of dependencies across environments.

## 🗼 Repository structure
```account-transactions
├───cmd                       # Service entry point
│   └───account-transactions
├───internal                  # Core application logic
│   ├───logger                # Logging utilities
│   ├───middleware            # HTTP middleware
│   ├───migrator              # DB migrations
│   ├───mocks                 # Test mocks
│   ├───signal                # Signal handling
│   ├───validator             # Input validation
│   └───writer                # Data writing logic
├───pkg
│   ├───handler               # HTTP handlers
│   │   ├───accounts          # Account handlers
│   │   └───transactions      # Transaction handlers
│   └───repository            # DB access
├───schema
│   └───migrations            # DB migrations
```
## 🚀 Api-spec

Documentation includes everything you need to get started with the API, including detailed endpoint descriptions, parameter definitions, response examples, and error handling information.

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

### Another method by pulling the docker image

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

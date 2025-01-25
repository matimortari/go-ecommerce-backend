# Go E-Commerce Backend 🛒⚡

Backend service for an e-commerce platform using Golang and PostgreSQL. It provides APIs for user authentication, product management, cart functionality, and order processing. <img src="https://static.velvetcache.org/pages/2018/06/13/party-gopher/dancing-gopher.gif" alt="Alt text" width="25" height="25">

## 📦 Key Features

- **Modular Architecture**: Clean and scalable codebase for easy maintenance.
- **Authentication**: Secure user authentication with JWT and password hashing.
- **Products**: Get, create and update products.
- **Cart**: Get and create orders with multiple products and secure validation.
- **Database Migrations**: Easy setup and migration of database schemas.
- **Automated Testing**: Easy-to-run tests for ensuring code reliability, using Go's built-in testing framework.
- **Makefile Automation**: Simplified commands for building, running, testing, and managing migrations.

## 📂 Directory Structure

```bash
  bin/          - Compiled binaries
  cmd/          - Main application entry points
  config/       - Configuration files
  services/     - Business logic and API handlers
  types/        - Shared data types
  utils/        - Utility functions
```

## 🏁 Getting Started

- Clone this repository:

  ```bash
  git clone https://github.com/matimortari/go-ecom-backend .
  ```

- Install dependencies:

  ```bash
  go mod tidy
  ```

- Create a `.env` file in the project root with the following environment variables (modify as needed):

  ```bash
   # Server configuration
  PUBLIC_HOST=http://localhost/
  PORT=8080

   # PostgreSQL database configuration
  DB_USER=postgres
  DB_PASSWORD=postgres
  DB_HOST=localhost
  DB_PORT=5432
  DB_NAME=ecom
  ```

- Use the `Makefile` to automate common tasks.

## ⚙️🧪 Makefile Commands

- Build the application to a binary at `bin/ecom`:

  ```bash
  make build
  ```

- Run the compiled binary:

  ```bash
  make run
  ```

- Run tests:

  ```bash
  make test
  ```

- Create a new migration file at `cmd/migrate/migrations`:

  ```bash
  make migration-create name=<migration_name>
  ```

- Run migrations "up" to apply all pending migrations:

  ```bash
  make migrate-up
  ```

- Run migrations "down" to roll back the last applied migration:

  ```bash
  make migrate-down
  ```

- Check migration status and version:
  ```bash
  make migrate-status
  ```

## 📬 Contact

Feel free to reach out to discuss collaboration opportunities or to say hello!

- [**My Email**](mailto:matheus.felipe.19rt@gmail.com)
- [**My LinkedIn Profile**](https://www.linkedin.com/in/matheus-mortari-19rt)
- [**My GitHub Profile**](https://github.com/matimortari)

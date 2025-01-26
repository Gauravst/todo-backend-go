# Todo List REST API in Go

![Go](https://img.shields.io/badge/Go-1.20-blue)
![REST API](https://img.shields.io/badge/REST-API-brightgreen)
![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-blue)

A simple REST API for managing a to-do list application, built with Go (Golang). This project is designed for beginners to learn and practice backend development concepts.

---

## Features

- **CRUD Operations**:
  - Create, read, update, and delete tasks.
- **HTTP Handling**:
  - Uses Go's `net/http` package to handle HTTP requests and responses.
- **Database Integration**:
  - Stores tasks in a PostgreSQL database.
- **JSON Support**:
  - Uses JSON for request and response payloads.
- **Structured Code**:
  - Follows clean and modular project structure.

---

## Skills Practiced

- Handling HTTP requests and responses in Go.
- Structuring a Go project for scalability and maintainability.
- Working with JSON for data serialization and deserialization.
- Integrating and interacting with a PostgreSQL database in Go.

---

## API Endpoints

| Method | Endpoint     | Description                    |
| ------ | ------------ | ------------------------------ |
| GET    | `/task`      | Get all tasks.                 |
| GET    | `/task/{id}` | Get a specific task by ID.     |
| POST   | `/task`      | Create a new task.             |
| PUT    | `/task/{id}` | Update an existing task by ID. |
| DELETE | `/task/{id}` | Delete a task by ID.           |

---

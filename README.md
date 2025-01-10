# Messaging Blog Service

Welcome to the **Messaging Blog Service**, a Go-based application designed using the **Hexagonal Architecture** (Ports and Adapters) pattern. This architecture ensures a clean separation of concerns, making the application modular, testable, and easy to maintain.

## Overview

The Messaging Blog Service is a platform that allows users to create and manage blogs, write messages, and interact with others. The application integrates multiple layers and interfaces while adhering to the principles of Hexagonal Architecture to support scalability and adaptability to future changes.

---

## Features

- **User Management**: Sign-up, login, and user profile handling.
- **Blog Management**: Create, update, delete, and view blogs.
- **Messaging System**: Write, edit, delete, and interact with messages.
- **Persistence**: Supports both SQL and NoSQL databases through repository adapters.
- **APIs**: RESTful APIs for client interaction.
- **Testing**: Comprehensive unit and integration test coverage.

---

## Project Structure

The project follows the Hexagonal Architecture, ensuring a clear separation of concerns:

messaging-blog/
│
├── cmd/                  # Application entry points
│   └── server/           # Main server setup
│
├── internal/             # Core application logic
│   ├── app/              # Application orchestrators (Use Cases)
│   ├── domain/           # Business logic (Entities and Aggregates)
│   ├── ports/            # Interfaces for inbound and outbound adapters
│   │   ├── inbound/      # Interfaces for API, CLI, etc.
│   │   └── outbound/     # Interfaces for repository, message brokers, etc.
│   ├── adapters/         # Adapters implementing the ports
│   │   ├── inbound/      # REST API implementation
│   │   └── outbound/     # Database repositories, brokers, etc.
│   └── config/           # Configuration files
│
├── pkg/                  # Shared utilities and helpers
├── tests/                # Unit and integration tests
└── README.md             # Documentation (this file)



---

## Getting Started

### Prerequisites

- [Go](https://golang.org/) (1.20 or newer)
- [Docker](https://www.docker.com/) (optional, for running dependencies)
- A database (e.g., PostgreSQL or Redis)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/messaging-blog.git
   cd messaging-blog

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
    Create a .env file in the project root and configure the following:

   ```bash
    APP_PORT=8080
    DATABASE_URL=postgres://username:password@localhost:5432/messaging_blog
   ```
4. Run the application:
    ```bash
    go run cmd/server/main.go

    ```

# Hexagonal Design
## Domain Layer
Defines the core business logic:
- Entities: Represent business objects such as User, Blog, and Message.
- Aggregates: Group related entities into cohesive units.
- Services: Contain domain-specific business rules.

## Application Layer
Implements use cases:
- Orchestrates the interaction between domain entities and adapters.
- Ensures that business rules are applied consistently.
## Ports
Interfaces that define contracts for inbound and outbound interactions:

- Inbound Ports: Define application services (e.g., BlogService, MessageService).
- Outbound Ports: Define interfaces for persistence, external APIs, etc.

## Adapters
Concrete implementations of ports:

- Inbound Adapters: REST API, CLI, or message handlers.
- Outbound Adapters: Database repositories, third-party APIs, or event brokers.

# Testing
The application includes comprehensive testing:

Unit Tests: Focus on individual use cases and domain logic.
Integration Tests: Test the interaction between components (e.g., database and services).
Run tests using:

```bash
go test ./..
```

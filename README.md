# Task Management Service

A task management backend service built with Go.

This project focuses on applying software engineering principles such as Domain-Driven Design (DDD), Clean Architecture, SOLID principles, and separation of concerns while building a complete backend system.

The application models business concepts such as users, projects, and tasks, with emphasis on maintainable domain logic, clear architectural boundaries, and thoughtful design decisions.

---

# Project Goals

The goal of this project is to build a backend service while practicing:

- Domain-Driven Design (DDD)
- Clean Architecture
- SOLID Principles
- Domain Modeling
- Separation of Concerns
- Design Patterns
- Testable code design
- Containerization
- CI/CD workflows
- Cloud-native deployment concepts

---

# Domain

The system focuses on task management with the following core concepts:

- User
- Project
- Task
- Task Priority
- Task Status
- Due Date

The domain model is designed around business rules rather than simple CRUD operations.

Examples of domain rules:

- Task title cannot be empty
- Priority must be valid
- Task status transitions must follow business rules
- Completed tasks cannot be reopened

---

# Architecture Approach

The project follows:

- Domain-Driven Design (DDD)
- Clean Architecture
- SOLID Principles
- Repository Pattern
- Dependency Inversion

The system is organized around business capabilities.

High-level architecture:

```
                 Client
                   |
                   |
                   v

          Interface Layer
              HTTP API

                   |
                   |

          Application Layer
             Use Cases

                   |
                   |

             Domain Layer
        Business Rules & Models

                   |
                   |

        Infrastructure Layer
 Database, External Services, Messaging
```

The domain layer remains independent from external technologies.

Business rules should not depend on:

- HTTP frameworks
- Databases
- Cloud providers
- Infrastructure libraries

---

# Project Structure

```
task-management-service/

├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │
│   ├── application/
│   │
│   ├── infrastructure/
│   │
│   └── interfaces/
│
├── docs/
│   ├── adr/
│   │   └── 0001-personal-task-management.md
│   │
│   ├── 01-domain-model.md
│   └── 02-architecture.md
│
├── migrations/
│
├── deployments/
│
├── scripts/
│
├── Dockerfile
│
├── docker-compose.yml
│
├── go.mod
│
└── README.md
```

---

# Documentation

Architecture and design decisions are documented in the `docs` directory.

Current documentation:

## Domain Model

Documents:

- Entities
- Value Objects
- Aggregates
- Business Rules
- Domain Events

## Architecture Decision Records (ADRs)

ADRs capture important technical decisions:

- Context behind a decision
- Alternatives considered
- Consequences

Example:

```
docs/adr/

0001-personal-task-management.md
```

---

# Technology Stack

## Backend

- Go

## Database

- PostgreSQL

## Containerization

- Docker

## Deployment

- Kubernetes

## Cloud

- AWS
- AWS ECR

## CI/CD

- GitHub Actions

---

# Engineering Practices

## Domain Modeling

This project explores:

- Entities
- Value Objects
- Aggregates
- Aggregate Roots
- Domain Events
- Domain Services

---

## Architecture

Practices applied:

- Clean Architecture
- Dependency Injection
- Repository Pattern
- Separation of Concerns
- Dependency Inversion Principle

---

## Backend Development

Topics covered:

- REST API design
- Request validation
- Error handling
- Database integration
- Testing strategies
- Configuration management

---

## Operations

Topics covered:

- Docker images
- Docker containers
- Docker networking
- Kubernetes deployments
- CI/CD automation
- Observability concepts

---

# Development Workflow

The project follows an incremental development approach:

```
Requirement

     |

Domain Discovery

     |

Architecture Decision Record

     |

Domain Model

     |

Implementation

     |

Testing

     |

Containerization

     |

Deployment
```

Important architectural decisions are documented before implementation.

---

# Current Status

Under development.

The project is being built incrementally with focus on:

- Understanding business requirements
- Designing the domain model
- Making intentional architectural decisions
- Writing maintainable Go code
- Applying backend engineering practices

---

# Future Enhancements

Potential future capabilities:

- Authentication and authorization
- User collaboration
- Task assignment
- Notifications
- Comments
- Attachments
- Activity history
- Audit logging
- Background workers
- Event-driven workflows
- Metrics and monitoring

---

# Learning Outcomes

Through this project, the focus is on understanding:

- How to model business domains
- How to design maintainable backend systems
- How to separate business logic from infrastructure
- How experienced engineers make architectural decisions
- How to evolve a service over time
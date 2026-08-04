# ADR-0001: Adopt Personal Task Management as the Core Domain

**Status:** Accepted

**Date:** 2026-08-04

---

## Context

The purpose of this project is to build a backend service using Go while following modern software engineering practices.

The application requires a domain that is:

- Simple enough to understand.
- Rich enough to demonstrate real business rules.
- Easily extensible as new requirements emerge.
- Suitable for showcasing backend engineering best practices.

A Personal Task Management domain satisfies these requirements.

---

## Decision

The application will model the **Personal Task Management** domain.

The primary business concepts are:

- User
- Project
- Task
- Task Title
- Priority
- Task Status
- Due Date

The project will prioritize modeling business behavior over implementing CRUD operations.

Business rules will be expressed through domain objects rather than being scattered across HTTP handlers, services, or database logic.

---

## Rationale

The Personal Task Management domain provides a realistic environment for demonstrating:

- Domain-Driven Design (DDD)
- Clean Architecture
- SOLID Principles
- Rich domain models
- Aggregate design
- Value Objects
- Repository Pattern
- Domain Events
- Testing business rules

It also provides a natural path for future enhancements without requiring significant architectural changes.

Possible future features include:

- Project sharing
- Team collaboration
- Comments
- Attachments
- Labels
- Recurring tasks
- Notifications
- Activity history
- Audit logging

---

## Consequences

### Positive

- The project reflects real business concepts instead of technical artifacts.
- The domain can evolve without major redesign.
- Business rules remain centralized inside the domain layer.
- The codebase becomes easier to understand and maintain.
- The project serves as a strong portfolio demonstrating backend architecture and design skills.

### Negative

- Initial development requires more design effort.
- Additional domain modeling introduces more types and abstractions than a simple CRUD application.
- The learning curve is higher than framework-driven development.

---

## Alternatives Considered

### Generic CRUD Application

Pros

- Faster to build.
- Minimal design effort.

Cons

- Encourages an anemic domain model.
- Business rules become scattered across the application.
- Demonstrates limited software design skills.

---

### Issue Tracking System

Pros

- Rich business domain.
- Real-world complexity.

Cons

- Significantly larger scope.
- Less suitable for incremental learning.

---

## Principles Applied

- Domain-Driven Design (DDD)
- Ubiquitous Language
- Rich Domain Model
- Separation of Concerns
- Clean Architecture
- SOLID Principles

---

## Related Documents

- `docs/01-domain-model.md`
- `docs/02-architecture.md` *(planned)*
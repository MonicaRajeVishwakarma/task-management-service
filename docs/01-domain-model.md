# Domain Model

## Task Management Service

---

## Purpose

The Task Management Service manages users, projects, and tasks.

The system is designed using Domain-Driven Design (DDD) principles, where business concepts and rules are modeled explicitly inside the domain layer.

The focus is not only exposing APIs, but building a backend service with a clear domain model and well-defined architectural boundaries.

---

# Bounded Context

## Task Management Context

The initial bounded context of this system is:

```
Task Management

├── User
├── Project
├── Task
├── Priority
├── Status
└── Due Date
```

Future bounded contexts may include:

```
Identity Context

Notification Context

Billing Context

Reporting Context
```

---

# Ubiquitous Language

These terms represent the shared language between developers and business stakeholders.

| Term | Description |
|------|-------------|
| User | Person who owns and manages projects |
| Project | A collection of related tasks |
| Task | A unit of work that needs completion |
| Task Title | Name describing the task |
| Priority | Importance level of a task |
| Status | Current lifecycle state of a task |
| Due Date | Expected completion date |

---

# Domain Overview

```
User

  |
  | owns
  v

Project

  |
  | contains
  v

Task Aggregate

  |
  ├── Task ID
  ├── Task Title
  ├── Description
  ├── Priority
  ├── Status
  └── Due Date
```

---

# Aggregates

## Task Aggregate

Aggregate Root:

```
Task
```

The Task entity controls all changes within the aggregate.

External code should not directly modify internal state.

Correct:

```
task.Complete()
```

Incorrect:

```
task.Status = COMPLETED
```

The aggregate protects business rules and ensures the object cannot enter an invalid state.

---

# Entities

## User

Represents a person using the system.

Identity:

```
UserID
```

---

## Project

Represents a collection of tasks.

Identity:

```
ProjectID
```

Responsibilities:

- Organize tasks
- Manage project lifecycle

---

## Task

Represents a business unit of work.

Identity:

```
TaskID
```

Responsibilities:

- Manage task lifecycle
- Enforce task rules
- Raise domain events

---

# Value Objects

Value Objects represent concepts where identity is not important.

They are defined by their value.

---

## TaskID

Represents the unique identifier of a task.

Example:

```
task-123
```

---

## TaskTitle

Represents a valid task title.

Rules:

- Cannot be empty
- Cannot contain only whitespace
- Maximum length rules will be defined later

---

## Priority

Represents task importance.

Allowed values:

```
LOW
MEDIUM
HIGH
```

---

## TaskStatus

Represents the lifecycle state.

Allowed values:

```
TODO
IN_PROGRESS
COMPLETED
```

---

## DueDate

Represents when a task should be completed.

Rules:

- Must represent a valid date
- Future business rules may restrict invalid deadlines

---

# Business Rules

The domain enforces the following rules:

## Task Title

A task cannot exist without a valid title.

---

## Priority

Only supported priority values are allowed.

---

## Task Lifecycle

Valid transitions:

```
TODO
 |
 v
IN_PROGRESS
 |
 v
COMPLETED
```

Invalid transitions:

```
COMPLETED -> TODO

COMPLETED -> IN_PROGRESS
```

A completed task cannot be reopened.

---

# Domain Events

Domain events represent important business facts that happened.

Potential events:

```
TaskCreated

TaskCompleted

TaskRenamed

TaskPriorityChanged

TaskDueDateChanged
```

Example:

When a task is completed:

```
Task.Complete()

        |
        v

TaskCompleted Event
```

The domain does not decide what happens next.

Other parts of the system may react.

Examples:

- Send notification
- Update activity history
- Trigger analytics

---

# Domain Independence

The domain layer does not depend on:

- HTTP
- PostgreSQL
- Docker
- Kubernetes
- AWS services

The domain only contains business concepts and rules.

---

# Design Principles Applied

## Domain-Driven Design

Business concepts are represented explicitly.

## Single Responsibility Principle

Each object owns a focused responsibility.

## Encapsulation

Objects protect their own rules.

## Dependency Inversion

The domain does not depend on infrastructure details.

## Separation of Concerns

Business logic, application flow, and infrastructure are separated.
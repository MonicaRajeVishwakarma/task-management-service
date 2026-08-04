# Domain Model

## Project

Todo API

---

# Purpose

The Todo API is a backend service for managing personal tasks organized into projects.

The primary goal is to model the business domain rather than simply exposing CRUD operations.

This project follows:

- Domain-Driven Design (DDD)
- Clean Architecture
- SOLID Principles
- Separation of Concerns

---

# Domain

Personal Task Management

---

# Ubiquitous Language

The following terms are shared between developers, product owners and stakeholders.

| Term | Description |
|------|-------------|
| User | Person who owns projects |
| Project | Collection of related tasks |
| Task | Work item owned by a project |
| Task Title | Human readable title of a task |
| Priority | Importance of a task |
| Status | Current lifecycle state of a task |
| Due Date | Expected completion date |
| Reminder | Notification before due date |

---

# Domain Overview

```

User
│
│ owns
▼
Project
│
│ contains
▼
Task
│
├── Title
├── Description
├── Priority
├── Status
└── Due Date

```

---

# Entities

## User

Represents a person using the system.

Identity:

- UserID

---

## Project

Represents a logical grouping of tasks.

Identity:

- ProjectID

---

## Task

Represents a unit of work.

Identity:

- TaskID

---

# Value Objects

## TaskTitle

Represents a valid task title.

Rules:

- Cannot be empty
- Cannot contain only whitespace
- Maximum length (to be defined)

---

## Priority

Represents task importance.

Allowed values:

- LOW
- MEDIUM
- HIGH

---

## TaskStatus

Represents the lifecycle of a task.

Allowed values:

- TODO
- IN_PROGRESS
- COMPLETED

---

## DueDate

Represents a valid task deadline.

Rules:

- Cannot be invalid date
- Future business rules may restrict past dates

---

# Aggregate

## Task Aggregate

Aggregate Root:

- Task

Contains:

- TaskTitle
- Priority
- TaskStatus
- DueDate
- Description

The Task Aggregate is responsible for enforcing all task-related business rules.

---

# Aggregate Root Responsibilities

The Task Aggregate Root controls:

- Completing a task
- Renaming a task
- Updating priority
- Updating due date
- Protecting lifecycle transitions

External code must never modify internal state directly.

Correct:

```
task.Complete()
```

Incorrect:

```
task.Status = COMPLETED
```

---

# Business Rules

Current rules:

1. Task title cannot be empty.
2. Priority must be valid.
3. Task status must be valid.
4. Completed tasks cannot be reopened.
5. Task lifecycle follows defined transitions.

---

# Allowed Status Transitions

```

TODO
│
├──────────────► COMPLETED
│
▼
IN_PROGRESS
│
▼
COMPLETED

```

Not allowed:

```

COMPLETED → TODO

COMPLETED → IN_PROGRESS

```

---

# Candidate Domain Events

The following domain events may be introduced later.

- TaskCreated
- TaskCompleted
- TaskRenamed
- TaskPriorityChanged
- TaskDueDateChanged

---

# Repository Contracts

The domain requires repositories but does not define persistence technology.

Examples:

- TaskRepository
- ProjectRepository
- UserRepository

Implementations may use:

- PostgreSQL
- MySQL
- MongoDB
- In-memory repository for testing

The domain remains independent of infrastructure.

---

# Design Goals

The domain should:

- Protect business rules
- Prevent invalid states
- Express business language
- Remain independent of frameworks
- Remain independent of databases
- Be easily testable
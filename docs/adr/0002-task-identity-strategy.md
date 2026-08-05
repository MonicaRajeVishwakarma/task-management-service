# ADR-0002: Use UUID as Task Identity

**Status:** Accepted

**Date:** 2026-08-04

---

## Context

The Task Management Service requires a way to uniquely identify domain entities.

The primary entity in the system is `Task`.

A task needs an identity that:

- Remains stable throughout its lifecycle
- Is independent from persistence concerns
- Can safely exist across distributed systems
- Can be exposed through APIs
- Does not require a database-generated sequence

The identity of a task is a domain concern and should not be represented only as a primitive value.

---

## Decision

The system will use UUIDs as the identity mechanism for domain entities.

The `TaskID` will be modeled as a Value Object inside the domain layer.

Example:

```
TaskID
 |
 +-- UUID value
```

The domain will work with:

```go
TaskID
```

instead of:

```go
string
```

or:

```go
int
```

---

## Rationale

UUIDs provide:

### Independence from Database

The database does not generate the identity.

The application can create an entity before persistence.

Example:

```
Create Task
      |
      v
Generate TaskID
      |
      v
Save Task
```

---

### Support for Distributed Systems

Multiple services can generate IDs without coordination.

Example:

```
Service A -> UUID-1

Service B -> UUID-2
```

No central sequence generator is required.

---

### Avoid Predictable IDs

Database sequences:

```
1
2
3
4
```

can expose internal information.

UUIDs:

```
550e8400-e29b-41d4-a716-446655440000
```

are harder to guess.

---

## Alternatives Considered

---

## Option 1: Database Generated Integer IDs

Example:

```
1
2
3
```

### Advantages

- Simple
- Efficient database indexes
- Easy to understand

### Disadvantages

- Database owns identity creation
- Harder to create entities before persistence
- Less suitable for distributed systems

Decision:

Rejected.

---

## Option 2: String Based IDs

Example:

```
TASK-10001
```

### Advantages

- Human readable
- Friendly for users

### Disadvantages

- Requires additional generation logic
- Collision handling
- More complexity

Decision:

Rejected for initial implementation.

---

## Option 3: UUID

Example:

```
550e8400-e29b-41d4-a716-446655440000
```

### Advantages

- Globally unique
- Database independent
- Distributed-system friendly
- Well supported in Go ecosystem

Decision:

Accepted.

---

## Consequences

### Positive

- Domain entities have explicit identity types.
- Business logic is independent from database implementation.
- Future services can exchange entity IDs safely.
- Code becomes more expressive.

Example:

Before:

```go
func FindTask(id string)
```

After:

```go
func FindTask(id TaskID)
```

The code communicates the domain.

---

### Negative

- UUIDs are larger than integers.
- Database indexes consume more storage.
- IDs are less human readable.

---

## Principles Applied

- Domain-Driven Design
- Value Objects
- Explicit Domain Modeling
- Dependency Inversion
- Separation of Concerns

---

## Related Documents

- `docs/01-domain-model.md`
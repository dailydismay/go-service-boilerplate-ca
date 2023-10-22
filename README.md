# Getting started

To run this app, execute `docker-compose -f ./deploy/dev/docker-compose.dev.yaml up -d`

Then run the cmd/service entrypoint.

# Service Modules

### Application

The UseCases and their implementations are placed here. For example, `/auth/login/login.go.`

Paths are built following the Screaming Architecture.

### Domain

This layer encapsulates the Entity, VO, Aggregate Root, and Domain Service Logic. It provides the definition of the Repository (Gateway/Provider) interfaces.

### Infrastructure

This layer contains the implementation of the repositories from the domain (using Postgres in this example), as well as crosscutting (shared) client packages.

### Interface

This is the delivery layer that can include Message Broker consumers, gRPC, GraphQL servers, etc.

### Core

This includes base definitions such as domain errors, usecase or entity base, validation utils, shared DTOs (pagination parameters/results), constants, and interfaces for cross-module dependencies (such as logger, cryptography, config properties, etc).

## Infrastructure interface matching

The Go way usually tells us to define the dependency interface above the unit constructor and mocking in each module which depends on it. _It causes code duplication and inconsistency._

**Since the repository interfaces are defined once in the domain layer, the problem is solved.**

**The infrastructure implementation constructor signature must define the return type as the domain repository interface.** This ensures that the code will not compile if there are changes in the domain interface signatures, and it clearly indicates the location of the implementation of the domain interface. This approach makes it easier for new code contributors to get started and reduces the time required to understand the architecture.

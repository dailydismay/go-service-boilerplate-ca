# Getting started

To run this app, execute `docker-compose -f ./deploy/dev/docker-compose.dev.yaml up -d`

Then execute `cmd/service` entrypoint.

# Service Modules

### Application

UseCases are placed here. For example `/auth/login/login.go`

Paths are built following Screamin Architecture.

### Domain

Entity, VO, Aggregate Root and Domain Service Logic encapsulated layer. Provides Repository(Gateway/Provider) interface definition

### Infrastructure

Includes repository implementation (postgres in this example), crosscutting(shared) client constructors.

### Interface

delivery layer (http in this example) which can be Message Broker consumer, rpc servcr implementation etc.

### Core

Root constructors as domain errors, usecase generic definitions, validation utils, shared DTOs like pagination options, interfaces for cross-module dependencies (like logger, cryptografy etc.)

## Infrastructure interface matching

1. As repository interfaces are defined once in `domain layer` (NOT in go-way like define dependency interface above the constructor and mock it, which causes code duplication and inconsistency)
2. Infrastructure implementation constructor signature defines return type as `domain repository interface`. This won't allow you to compile code when you change domain intefaces signature and will certainly show, where is implementation of domain interface. (This gives low treshold to new code contributers and takes less time to understand architecture)

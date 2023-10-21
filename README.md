# Getting started

To run this app, execute `docker-compose -f docker-compose.dev.yaml up -d`

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

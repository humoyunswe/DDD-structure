Copy the full vertical slice from `internal/**/v1/users/**` and rename the module.

Required files:
- domain: entity.go, repository.go, errors.go
- application: ports.go, usecase.go
- models + pg repository
- dto, handler, router
- wire in cmd/server/main.go + Dependencies + engineRouter


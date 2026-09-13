package users_use_case

import domain "ddd-structure/internal/domain/v1/users"

// RepositoryPort aliases the domain repository so application stays decoupled
// from concrete infrastructure packages (same pattern as szpt_new).
type RepositoryPort = domain.Repository

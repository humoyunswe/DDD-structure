package users

import "time"

// User is a domain entity — no JSON/DB tags here.
type User struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateUserCmd is an application command carried through domain.
type CreateUserCmd struct {
	Email     string
	FirstName string
	LastName  string
}

// UpdateUserCmd updates mutable profile fields.
type UpdateUserCmd struct {
	FirstName string
	LastName  string
	IsActive  bool
}

// ListParams is pagination for list queries.
type ListParams struct {
	Limit  int
	Offset int
}

package users

import "context"

// Repository is the persistence port owned by the domain layer.
// Infrastructure implements this interface; application depends on it via ports alias.
type Repository interface {
	Create(ctx context.Context, cmd CreateUserCmd) (User, error)
	GetByID(ctx context.Context, id int64) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	List(ctx context.Context, params ListParams) ([]User, int64, error)
	Update(ctx context.Context, id int64, cmd UpdateUserCmd) (User, error)
	Delete(ctx context.Context, id int64) error
}

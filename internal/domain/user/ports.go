package user
package user

import (







































}	ListUsers(ctx context.Context, limit, offset int) ([]*User, error)	GetUserByUsername(ctx context.Context, username Username) (*User, error)	GetUserByEmail(ctx context.Context, email Email) (*User, error)	GetUserByID(ctx context.Context, userID UserID) (*User, error)	// Query operations		ActivateUser(ctx context.Context, userID UserID) error	SuspendUser(ctx context.Context, userID UserID) error	ChangePassword(ctx context.Context, userID UserID, oldPassword, newPassword string) error	UpdateUserProfile(ctx context.Context, userID UserID, firstName, lastName, bio string) error	CreateUser(ctx context.Context, email, username, password string) (*User, error)	// User management operations		ValidateCredentials(ctx context.Context, email Email, password string) (*User, error)	// Authentication operationstype UserService interface {// UserService defines domain services for user operations}	CountByStatus(ctx context.Context, status UserStatus) (int64, error)	Count(ctx context.Context) (int64, error)	ExistsByUsername(ctx context.Context, username Username) (bool, error)	ExistsByEmail(ctx context.Context, email Email) (bool, error)	FindByStatus(ctx context.Context, status UserStatus, limit, offset int) ([]*User, error)	FindAll(ctx context.Context, limit, offset int) ([]*User, error)	// Query operations	Delete(ctx context.Context, id UserID) error	Update(ctx context.Context, user *User) error	FindByUsername(ctx context.Context, username Username) (*User, error)	FindByEmail(ctx context.Context, email Email) (*User, error)	FindByID(ctx context.Context, id UserID) (*User, error)	Save(ctx context.Context, user *User) error	// Basic CRUD operationstype UserRepository interface {// UserRepository defines the contract for user persistence)	"context"
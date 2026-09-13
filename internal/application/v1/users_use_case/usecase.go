package users_use_case

import (
	"context"
	"errors"
	"strings"

	domain "ddd-structure/internal/domain/v1/users"
)

type Service struct {
	repo RepositoryPort
}

func NewService(repo RepositoryPort) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, cmd domain.CreateUserCmd) (domain.User, error) {
	cmd.Email = strings.TrimSpace(strings.ToLower(cmd.Email))
	cmd.FirstName = strings.TrimSpace(cmd.FirstName)
	cmd.LastName = strings.TrimSpace(cmd.LastName)

	if cmd.Email == "" {
		return domain.User{}, domain.ErrInvalidInput{Message: "email is required"}
	}
	if cmd.FirstName == "" {
		return domain.User{}, domain.ErrInvalidInput{Message: "first_name is required"}
	}

	_, err := s.repo.GetByEmail(ctx, cmd.Email)
	if err == nil {
		return domain.User{}, domain.ErrAlreadyExists{}
	}
	var notFound domain.ErrNotFound
	if !errors.As(err, &notFound) {
		return domain.User{}, err
	}

	return s.repo.Create(ctx, cmd)
}

func (s *Service) GetByID(ctx context.Context, id int64) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, domain.ErrInvalidInput{Message: "id must be positive"}
	}
	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context, params domain.ListParams) ([]domain.User, int64, error) {
	if params.Limit <= 0 {
		params.Limit = 20
	}
	if params.Limit > 100 {
		params.Limit = 100
	}
	if params.Offset < 0 {
		params.Offset = 0
	}
	return s.repo.List(ctx, params)
}

func (s *Service) Update(ctx context.Context, id int64, cmd domain.UpdateUserCmd) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, domain.ErrInvalidInput{Message: "id must be positive"}
	}
	cmd.FirstName = strings.TrimSpace(cmd.FirstName)
	cmd.LastName = strings.TrimSpace(cmd.LastName)
	return s.repo.Update(ctx, id, cmd)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return domain.ErrInvalidInput{Message: "id must be positive"}
	}
	return s.repo.Delete(ctx, id)
}

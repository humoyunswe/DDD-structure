package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	domain "ddd-structure/internal/domain/v1/users"
	models "ddd-structure/internal/infrastructure/persistence/models/v1/users"

	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) Create(ctx context.Context, cmd domain.CreateUserCmd) (domain.User, error) {
	now := time.Now().UTC()
	row := models.UserModel{
		Email:     cmd.Email,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := r.db.WithContext(ctx).Create(&row).Error; err != nil {
		return domain.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return toDomain(row), nil
}

func (r *RepositoryImpl) GetByID(ctx context.Context, id int64) (domain.User, error) {
	var row models.UserModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, domain.ErrNotFound{}
	}
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user by id: %w", err)
	}
	return toDomain(row), nil
}

func (r *RepositoryImpl) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var row models.UserModel
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, domain.ErrNotFound{}
	}
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user by email: %w", err)
	}
	return toDomain(row), nil
}

func (r *RepositoryImpl) List(ctx context.Context, params domain.ListParams) ([]domain.User, int64, error) {
	var total int64
	q := r.db.WithContext(ctx).Model(&models.UserModel{})
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	var rows []models.UserModel
	if err := q.Order("id ASC").Limit(params.Limit).Offset(params.Offset).Find(&rows).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}

	result := make([]domain.User, len(rows))
	for i, row := range rows {
		result[i] = toDomain(row)
	}
	return result, total, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, id int64, cmd domain.UpdateUserCmd) (domain.User, error) {
	var row models.UserModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, domain.ErrNotFound{}
	}
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to load user for update: %w", err)
	}

	row.FirstName = cmd.FirstName
	row.LastName = cmd.LastName
	row.IsActive = cmd.IsActive
	row.UpdatedAt = time.Now().UTC()

	if err := r.db.WithContext(ctx).Save(&row).Error; err != nil {
		return domain.User{}, fmt.Errorf("failed to update user: %w", err)
	}
	return toDomain(row), nil
}

func (r *RepositoryImpl) Delete(ctx context.Context, id int64) error {
	res := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.UserModel{})
	if res.Error != nil {
		return fmt.Errorf("failed to delete user: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return domain.ErrNotFound{}
	}
	return nil
}

func toDomain(row models.UserModel) domain.User {
	return domain.User{
		ID:        row.ID,
		Email:     row.Email,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		IsActive:  row.IsActive,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

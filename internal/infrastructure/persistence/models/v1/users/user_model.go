package users

import "time"

// UserModel is a GORM persistence model.
// Keep mapping to/from domain.User inside the repository — not in handlers.
type UserModel struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Email     string    `gorm:"column:email;uniqueIndex;size:255;not null"`
	FirstName string    `gorm:"column:first_name;size:100;not null"`
	LastName  string    `gorm:"column:last_name;size:100;not null"`
	IsActive  bool      `gorm:"column:is_active;not null;default:true"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (UserModel) TableName() string { return "app.users" }

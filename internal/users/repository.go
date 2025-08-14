package users

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) Create(data UserModel) error {
	return u.db.Create(&data).Error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

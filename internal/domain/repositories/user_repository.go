package repository

import (
	"errors"
	"fx-bank/internal/domain/models"
	"gorm.io/gorm"
)

var (
	ErrUserAccountFailedToCreate = errors.New("account failed to create")
	ErrUserAccountDoesNotExist   = errors.New("account does not exist")
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) CreateUserAccount(user *models.User) error {
	if u.DB.Create(&user).Error != nil {
		return ErrUserAccountFailedToCreate
	}
	return nil
}

func (u *UserRepository) FindUserAccountByUsername(username string) (*models.User, error, int64) {
	var user *models.User
	db := u.DB.Where("username = ?", username).Find(&user)
	if db.Error != nil {
		return nil, db.Error, 0
	}
	return user, nil, db.RowsAffected
}

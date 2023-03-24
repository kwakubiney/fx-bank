package repository

import (
	"errors"
	"fx-bank/internal/domain/models"
	"gorm.io/gorm"
)

var (
	ErrAccountNotFound        = errors.New("account not found")
	ErrAccountCannotCreated   = errors.New("account cannot be created")
	ErrAccountNameAlreadyUsed = errors.New("account name already used")
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (a *AccountRepository) CreateAccount(account *models.Account) error {
	if a.DB.Create(&account).Error != nil {
		return ErrAccountCannotCreated
	}
	return nil
}

func (a *AccountRepository) FindAccountByNameAndUserID(name string, id string) error {
	var account *models.Account
	db := a.DB.Where("user_id = ? AND name = ?", id, name).Find(&account)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return nil
	}
	return ErrAccountNameAlreadyUsed
}

func (a *AccountRepository) FindAllAccountsByUserID(id string) ([]*models.Account, error) {
	var accounts []*models.Account
	db := a.DB.Where("user_id = ?", id).Find(&accounts)
	if db.Error != nil {
		return nil, db.Error
	}
	return accounts, nil
}

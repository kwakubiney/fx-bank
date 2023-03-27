package repository

import (
	"errors"
	"fx-bank/internal/domain/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	ErrAccountNotFound         = errors.New("account not found")
	ErrAccountSenderNotFound   = errors.New("account used for sending not found")
	ErrAccountReceiverNotFound = errors.New("account used for receiving not found")
	ErrAccountCannotCreated    = errors.New("account cannot be created")
	ErrAccountNameAlreadyUsed  = errors.New("account name already used")
)

type AccountRepository struct {
	DB *gorm.DB
}

func (a *AccountRepository) WithTrx(trxHandle *gorm.DB) *AccountRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return &AccountRepository{}
	}
	a.DB = trxHandle
	return a
}

func (t *TransactionRepository) WithTrx(trxHandle *gorm.DB) *TransactionRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return &TransactionRepository{}
	}

	t.DB = trxHandle
	return t
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

func (a *AccountRepository) FindAccountByID(originID string, destinationID string) (*models.Account, *models.Account, error) {
	var originAccount *models.Account
	var destinationAccount *models.Account
	db := a.DB.Where("id = ?", originID).Find(&originAccount)
	if db.RowsAffected == 0 {
		return nil, nil, ErrAccountSenderNotFound
	}
	db = a.DB.Where("id = ?", destinationID).Find(&destinationAccount)
	if db.RowsAffected == 0 {
		return nil, nil, ErrAccountReceiverNotFound
	}
	return originAccount, destinationAccount, nil
}

func (a *AccountRepository) Transfer(sender models.Account, receiver models.Account,
	amount decimal.Decimal, rate decimal.Decimal) error {
	err := sender.Withdraw(amount, rate)
	if err != nil {
		return err
	}
	err = a.DB.Model(&sender).Update("balance", sender.Balance).Error
	if err != nil {
		log.Println(err)
		return err
	}

	valueToAdd := receiver.Deposit(amount)
	receiver.Balance = valueToAdd
	err = a.DB.Model(&receiver).Update("balance", receiver.Balance).Error
	if err != nil {
		log.Println(err)
		return err
	}

	err = a.DB.Model(&sender).Update("last_modified", time.Now()).Error
	if err != nil {
		log.Println(err)
		return err
	}

	err = a.DB.Model(&receiver).Update("last_modified", time.Now()).Error
	if err != nil {
		log.Println(err)
		return err
	}

	//TODO: Update last modified on both accounts

	return nil
}

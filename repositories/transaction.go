package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindTransaction() ([]models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	// UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetOrderByID(orderID int) ([]models.Order, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Order").Create(&transaction).Error
	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Order").First(&transaction, ID).Error
	return transaction, err
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Find(&transaction).Error
	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error
	return transaction, err
}

func (r *repository) GetOrderByID(ID int) ([]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("Product").Preload("Topping").Preload("Buyyer").Where("buyyer_id = ?", ID).Find(&order).Error
	return order, err
}

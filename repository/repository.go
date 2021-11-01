package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyeeego/flowershop/domain"
)

type Repository struct {
	Customer CustomerRepository
	Bouquet  BouquetRepository
	Seller   SellerRepository
	Purchase PurchaseRepository
}

type CustomerRepository interface {
	Insert(model *domain.CustomerModel) error
	FindById(id uint) (*domain.CustomerModel, error)
	Update(model *domain.CustomerModel) error
	Delete(id uint) error
}

type BouquetRepository interface {
	Insert(model *domain.BouquetModel) error
	FindById(id uint) (*domain.BouquetModel, error)
	FindAll() ([]domain.BouquetModel, error)
	Update(model *domain.BouquetModel) error
	Delete(id uint) error
}

type SellerRepository interface {
	Insert(model *domain.SellerModel) error
	FindById(id uint) (*domain.SellerModel, error)
	FindAll() ([]domain.SellerModel, error)
	Update(model *domain.SellerModel) error
	Delete(id uint) error
}

type PurchaseRepository interface {
	Insert(model *domain.PurchaseModel) error
	FindByCustomerId(customerId uint) ([]domain.PurchaseModel, error)
}

func New(db *gorm.DB) *Repository {
	db.AutoMigrate(domain.BouquetModel{},
		domain.CustomerModel{},
		domain.PurchaseModel{},
		domain.SellerModel{})

	return &Repository{
		Customer: newCustomerRepo(db),
		Bouquet:  newBouquetRepo(db),
		Purchase: newPurchaseRepo(db),
		Seller:   newSellerRepo(db),
	}
}

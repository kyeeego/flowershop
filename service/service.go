package service

import (
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/repository"
)

type CustomerService interface {
	GetById(id uint) (domain.CustomerDto, error)
	Create(dto domain.CreateCustomerDto) (domain.CustomerDto, error)
	Update(dto domain.UpdateCustomerDto) (domain.CustomerDto, error)
	Delete(id uint) error
}

type BouquetService interface {
	GetById(id uint) (domain.BouquetDto, error)
	GetAll() ([]domain.BouquetDto, error)
	Create(dto domain.CreateBouquetDto) (domain.BouquetDto, error)
	Update(dto domain.UpdateBouquetDto) (domain.BouquetDto, error)
	Delete(id uint) error
}

type SellerService interface {
	GetById(id uint) (domain.SellerDto, error)
	GetAll() ([]domain.SellerDto, error)
	Create(dto domain.CreateSellerDto) (domain.SellerDto, error)
	Update(dto domain.UpdateSellerDto) (domain.SellerDto, error)
	Delete(id uint) error
}

type PurchaseService interface {
	GetAll(customerId uint) ([]domain.PurchaseDto, error)
	Do(customerId, bouquetId uint) error
}

type Service struct {
	Customer CustomerService
	Seller   SellerService
	Bouquet  BouquetService
	Purchase PurchaseService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Customer: newCustomerService(repository),
		Seller:   newSellerService(repository),
		Bouquet:  newBouquetService(repository),
		Purchase: newPurchaseService(repository),
	}
}

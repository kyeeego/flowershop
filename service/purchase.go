package service

import (
	"github.com/jinzhu/copier"
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/repository"
)

type PurchaseServ struct {
	repository *repository.Repository
}

func newPurchaseService(r *repository.Repository) *PurchaseServ {
	return &PurchaseServ{r}
}

func (s PurchaseServ) GetAll(customerId uint) ([]domain.PurchaseDto, error) {
	var dtos []domain.PurchaseDto
	models, err := s.repository.Purchase.FindByCustomerId(customerId)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&dtos, &models)

	return dtos, err
}

func (s PurchaseServ) Do(customerId, bouquetId uint) error {
	bouquet, err := s.repository.Bouquet.FindById(bouquetId)
	if err != nil {
		return err
	}

	model := domain.PurchaseModel{
		BouquetId:  bouquetId,
		CustomerId: customerId,
		Price:      bouquet.Price,
		Profit:     bouquet.Price * 0.3,
	}

	return s.repository.Purchase.Insert(&model)
}

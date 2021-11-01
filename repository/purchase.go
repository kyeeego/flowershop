package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyeeego/flowershop/domain"
)

type PurchaseRepo struct {
	db *gorm.DB
}

func newPurchaseRepo(db *gorm.DB) *PurchaseRepo {
	return &PurchaseRepo{db}
}

func (r *PurchaseRepo) FindByCustomerId(customerId uint) ([]domain.PurchaseModel, error) {
	var models []domain.PurchaseModel
	res := r.db.Find(&models, domain.PurchaseModel{CustomerId: customerId})

	return models, res.Error
}

func (r *PurchaseRepo) Insert(model *domain.PurchaseModel) error {
	res := r.db.Create(model)
	return res.Error
}

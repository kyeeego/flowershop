package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyeeego/flowershop/domain"
)

type SellerRepo struct {
	db *gorm.DB
}

func newSellerRepo(db *gorm.DB) *SellerRepo {
	return &SellerRepo{db}
}

func (r SellerRepo) FindById(id uint) (*domain.SellerModel, error) {
	model := domain.SellerModel{}
	res := r.db.First(&model, id)

	return &model, res.Error
}

func (r SellerRepo) FindAll() ([]domain.SellerModel, error) {
	var models []domain.SellerModel
	res := r.db.Find(&models)

	return models, res.Error
}

func (r *SellerRepo) Insert(model *domain.SellerModel) error {
	res := r.db.Create(model)
	return res.Error
}

func (r SellerRepo) Delete(id uint) error {
	res := r.db.Delete(&domain.SellerModel{}, id)
	return res.Error
}

func (r SellerRepo) Update(model *domain.SellerModel) error {
	res := r.db.Save(model)
	return res.Error
}

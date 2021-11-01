package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyeeego/flowershop/domain"
)

type CustomerRepo struct {
	db *gorm.DB
}

func newCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db}
}

func (r CustomerRepo) FindById(id uint) (*domain.CustomerModel, error) {
	model := domain.CustomerModel{}
	res := r.db.First(&model, id)

	return &model, res.Error
}

func (r *CustomerRepo) Insert(model *domain.CustomerModel) error {
	res := r.db.Create(model)
	return res.Error
}

func (r CustomerRepo) Delete(id uint) error {
	res := r.db.Delete(&domain.CustomerModel{}, id)
	return res.Error
}

func (r CustomerRepo) Update(model *domain.CustomerModel) error {
	res := r.db.Save(model)
	return res.Error
}


package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyeeego/flowershop/domain"
)

type BouquetRepo struct {
	db *gorm.DB
}

func newBouquetRepo(db *gorm.DB) *BouquetRepo {
	return &BouquetRepo{db}
}

func (r BouquetRepo) FindById(id uint) (*domain.BouquetModel, error) {
	model := domain.BouquetModel{}
	res := r.db.First(&model, id)

	return &model, res.Error
}

func (r BouquetRepo) FindAll() ([]domain.BouquetModel, error) {
	var models []domain.BouquetModel
	res := r.db.Find(&models)

	return models, res.Error
}

func (r *BouquetRepo) Insert(model *domain.BouquetModel) error {
	res := r.db.Create(model)
	return res.Error
}

func (r BouquetRepo) Delete(id uint) error {
	res := r.db.Delete(&domain.BouquetModel{}, id)
	return res.Error
}

func (r BouquetRepo) Update(model *domain.BouquetModel) error {
	res := r.db.Save(model)
	return res.Error
}

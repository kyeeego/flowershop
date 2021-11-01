package service

import (
	"github.com/jinzhu/copier"
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/repository"
)

type BouquetServ struct {
	r *repository.Repository
}

func newBouquetService(r *repository.Repository) *BouquetServ {
	return &BouquetServ{r}
}

func (s BouquetServ) GetById(id uint) (domain.BouquetDto, error) {
	model, err := s.r.Bouquet.FindById(id)
	if err != nil {
		return domain.BouquetDto{}, err
	}

	dto := domain.BouquetDto{}
	err = copier.Copy(&dto, &model)
	return dto, err
}

func (s BouquetServ) GetAll() ([]domain.BouquetDto, error) {
	var dtos []domain.BouquetDto
	models, err := s.r.Bouquet.FindAll()
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&dtos, &models)

	return dtos, err
}

func (s BouquetServ) Create(dto domain.CreateBouquetDto) (domain.BouquetDto, error) {
	model := domain.BouquetModel{
		Name:     dto.Name,
		Price:    dto.Price,
		PhotoUrl: dto.PhotoUrl,
		SellerId: dto.SellerId,
	}

	err := s.r.Bouquet.Insert(&model)
	if err != nil {
		return domain.BouquetDto{}, nil
	}

	resDto := domain.BouquetDto{}
	err = copier.Copy(&resDto, &model)
	return resDto, err
}

func (s BouquetServ) Update(dto domain.UpdateBouquetDto) (domain.BouquetDto, error) {
	model, err := s.r.Bouquet.FindById(dto.ID)
	if err != nil {
		return domain.BouquetDto{}, err
	}

	err = copier.Copy(&model, &dto)
	if err != nil {
		return domain.BouquetDto{}, err
	}

	res := domain.BouquetDto{}
	err = copier.Copy(&res, &model)
	if err != nil {
		return domain.BouquetDto{}, err
	}

	return res, s.r.Bouquet.Update(model)
}

func (s BouquetServ) Delete(id uint) error {
	return s.r.Seller.Delete(id)
}

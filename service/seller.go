package service

import (
	"github.com/jinzhu/copier"
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/repository"
)

type SellerServ struct {
	repository *repository.Repository
}

func newSellerService(r *repository.Repository) *SellerServ {
	return &SellerServ{r}
}

func (s SellerServ) GetById(id uint) (domain.SellerDto, error) {
	model, err := s.repository.Seller.FindById(id)
	if err != nil {
		return domain.SellerDto{}, err
	}

	dto := domain.SellerDto{}
	err = copier.Copy(&dto, &model)
	return dto, err
}

func (s SellerServ) GetAll() ([]domain.SellerDto, error) {
	var dtos []domain.SellerDto
	models, err := s.repository.Seller.FindAll()
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&dtos, &models)

	return dtos, err
}

func (s SellerServ) Create(dto domain.CreateSellerDto) (domain.SellerDto, error) {
	model := domain.SellerModel{
		ShopName:     dto.ShopName,
		PhotoUrl:     dto.PhotoUrl,
		Bouquets:     []domain.BouquetModel{},
		SoldBouquets: 0,
	}

	err := s.repository.Seller.Insert(&model)
	if err != nil {
		return domain.SellerDto{}, nil
	}

	resDto := domain.SellerDto{}
	err = copier.Copy(&resDto, &model)
	return resDto, err
}

func (s SellerServ) Update(dto domain.UpdateSellerDto) (domain.SellerDto, error) {
	model, err := s.repository.Seller.FindById(dto.ID)
	if err != nil {
		return domain.SellerDto{}, err
	}

	err = copier.Copy(&model, &dto)
	if err != nil {
		return domain.SellerDto{}, err
	}

	res := domain.SellerDto{}
	err = copier.Copy(&res, &model)
	if err != nil {
		return domain.SellerDto{}, err
	}

	return res, s.repository.Seller.Update(model)
}

func (s SellerServ) Delete(id uint) error {
	return s.repository.Seller.Delete(id)
}

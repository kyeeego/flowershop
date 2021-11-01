package service

import (
	"github.com/jinzhu/copier"
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/repository"
)

type CustomerServ struct {
	r *repository.Repository
}

func newCustomerService(r *repository.Repository) *CustomerServ {
	return &CustomerServ{r}
}

func (c CustomerServ) GetById(id uint) (domain.CustomerDto, error) {
	model, err := c.r.Customer.FindById(id)
	if err != nil {
		return domain.CustomerDto{}, err
	}

	dto := domain.CustomerDto{}
	err = copier.Copy(&dto, &model)
	return dto, err
}

func (c CustomerServ) Create(dto domain.CreateCustomerDto) (domain.CustomerDto, error) {
	model := domain.CustomerModel{
		Name:      dto.Name,
		Email:     dto.Email,
		Purchases: []domain.PurchaseModel{},
	}

	err := c.r.Customer.Insert(&model)
	if err != nil {
		return domain.CustomerDto{}, nil
	}

	resDto := domain.CustomerDto{}
	err = copier.Copy(&resDto, &model)
	return resDto, err
}

func (c CustomerServ) Update(dto domain.UpdateCustomerDto) (domain.CustomerDto, error) {
	model, err := c.r.Customer.FindById(dto.ID)
	if err != nil {
		return domain.CustomerDto{}, err
	}

	err = copier.Copy(&model, &dto)
	if err != nil {
		return domain.CustomerDto{}, err
	}

	res := domain.CustomerDto{}
	err = copier.Copy(&res, &model)
	if err != nil {
		return domain.CustomerDto{}, err
	}

	return res, c.r.Customer.Update(model)
}

func (c CustomerServ) Delete(id uint) error {
	return c.r.Customer.Delete(id)
}


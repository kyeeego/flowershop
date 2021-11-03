package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type CustomerModel struct {
	gorm.Model
	Name      string
	Email     string
	Purchases []PurchaseModel `gorm:"foreignKey:CustomerId"`
}

var GqlCustomer = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Customer",
		Fields: graphql.Fields{
			"id":        {Type: graphql.Int},
			"name":      {Type: graphql.String},
			"email":     {Type: graphql.String},
			"purchases": {Type: graphql.NewList(GqlPurchase)},
		},
	},
)

type CustomerDto struct {
	ID        uint          `json:"id" copier:"must"`
	Name      string        `json:"name" copier:"must"`
	Email     string        `json:"email" copier:"must"`
	Purchases []PurchaseDto `json:"purchases" copier:"must"`
}

type CreateCustomerDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateCustomerDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

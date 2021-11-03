package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type PurchaseModel struct {
	gorm.Model
	BouquetId  uint
	CustomerId uint
	Price      float64
	Profit     float64
}

var GqlPurchase = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Purchase",
		Fields: graphql.Fields{
			"id":         {Type: graphql.Int},
			"bouquetId":  {Type: graphql.Int},
			"customerId": {Type: graphql.Int},
			"price":      {Type: graphql.Float},
			"profit":     {Type: graphql.Float},
		},
	},
)

type PurchaseDto struct {
	ID         uint    `json:"id"`
	BouquetID  uint    `json:"bouquet"`
	CustomerID uint    `json:"customer"`
	Price      float64 `json:"price"`
	Profit     float64 `json:"profit"`
}

type DoPurchaseDto struct {
	CustomerID uint `json:"customerId"`
	BouquetID  uint `json:"bouquetId"`
}

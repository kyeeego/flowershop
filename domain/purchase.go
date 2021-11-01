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
			"id":         {Type: graphql.ID},
			"bouquetId":  {Type: graphql.ID},
			"customerId": {Type: graphql.ID},
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
	Profit     float64 `json:"service_income"`
}

var GqlDoPurchaseDto = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DoPurchaseDto",
		Fields: graphql.Fields{
			"customerId": {Type: graphql.ID},
			"bouquetId":  {Type: graphql.ID},
		},
	},
)

type DoPurchaseDto struct {
	CustomerID uint `json:"customer_id"`
	BouquetID  uint `json:"bouquet_id"`
}

package domain

import "github.com/graphql-go/graphql"

var GqlPurchase = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Purchase",
		Fields: graphql.Fields{
			"id":            {Type: graphql.ID},
			"bouquetId":     {Type: graphql.ID},
			"customerId":    {Type: graphql.ID},
			"price":         {Type: graphql.Float},
			"serviceIncome": {Type: graphql.Float},
		},
	},
)

type PurchaseDto struct {
	Id            uint32       `json:"id"`
	Bouquet       *BouquetDto  `json:"bouquet"`
	Customer      *CustomerDto `json:"customer"`
	Price         float64      `json:"price"`
	ServiceIncome float64      `json:"service_income"`
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
	CustomerId uint32 `json:"customer_id"`
	BouquetId  uint32 `json:"bouquet_id"`
}

package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type BouquetModel struct {
	gorm.Model
	Name     string
	Price    float64
	PhotoUrl string
	SellerId uint
}

var GqlBouquet = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Bouquet",
		Fields: graphql.Fields{
			"id":       {Type: graphql.Int},
			"name":     {Type: graphql.String},
			"price":    {Type: graphql.Float},
			"photoUrl": {Type: graphql.String},
			"sellerId": {Type: graphql.Int},
		},
	},
)

type BouquetDto struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	Price    float64    `json:"price"`
	PhotoUrl string     `json:"photoUrl"`
	Seller   *SellerDto `json:"seller"`
}

type CreateBouquetDto struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	PhotoUrl string  `json:"photoUrl"`
	SellerId uint    `json:"sellerId"`
}

type UpdateBouquetDto struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	PhotoUrl string  `json:"photoUrl,omitempty"`
}

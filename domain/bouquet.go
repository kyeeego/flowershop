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
			"id":       {Type: graphql.ID},
			"name":     {Type: graphql.String},
			"price":    {Type: graphql.Float},
			"photoUrl": {Type: graphql.String},
			"sellerId": {Type: graphql.ID},
		},
	},
)

type BouquetDto struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	Price    float64    `json:"price"`
	PhotoUrl string     `json:"photo_url"`
	Seller   *SellerDto `json:"seller"`
}

var GqlCreateBouquetDto = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreateBouquetDto",
		Fields: graphql.Fields{
			"name":     {Type: graphql.String},
			"price":    {Type: graphql.Float},
			"photoUrl": {Type: graphql.String},
			"sellerId": {Type: graphql.ID},
		},
	},
)

type CreateBouquetDto struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	PhotoUrl string  `json:"photo_url"`
	SellerId uint    `json:"seller_id"`
}

var GqlUpdateBouquetDto = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UpdateBouquetDto",
		Fields: graphql.Fields{
			"id":       {Type: graphql.ID},
			"name":     {Type: graphql.String},
			"price":    {Type: graphql.Float},
			"photoUrl": {Type: graphql.String},
		},
	},
)

type UpdateBouquetDto struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	PhotoUrl string  `json:"photo_url,omitempty"`
}

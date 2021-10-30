package domain

import "github.com/graphql-go/graphql"

var GqlBouquet = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Bouquet",
		Fields: graphql.Fields{
			"id":       {Type: graphql.ID},
			"name":     {Type: graphql.String},
			"price":    {Type: graphql.Float},
			"photoUrl": {Type: graphql.String},
			"seller":   {Type: GqlSeller},
		},
	},
)

type BouquetDto struct {
	Id       uint32     `json:"id"`
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
	SellerId uint32  `json:"seller_id"`
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
	Id       uint32  `json:"id"`
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	PhotoUrl string  `json:"photo_url,omitempty"`
}

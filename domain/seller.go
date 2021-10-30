package domain

import (
	"github.com/graphql-go/graphql"
)

var GqlSeller = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Seller",
		Fields: graphql.Fields{
			"id":           {Type: graphql.ID},
			"shopName":     {Type: graphql.String},
			"photoUrl":     {Type: graphql.String},
			"createdDate":  {Type: graphql.Int},
			"bouquets":     {Type: graphql.NewList(GqlBouquet)},
			"soldBouquets": {Type: graphql.Int},
		},
	},
)

type SellerDto struct {
	Id           uint32        `json:"id"`
	ShopName     string        `json:"shop_name"`
	PhotoUrl     string        `json:"photo_url"`
	CreatedDate  int32         `json:"created_date"`
	Bouquets     []*BouquetDto `json:"bouquets"`
	SoldBouquets int           `json:"sold_bouquets"`
}

var GqlCreateSellerDto = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreateSellerDto",
		Fields: graphql.Fields{
			"shopName": {Type: graphql.String},
			"photoUrl": {Type: graphql.String},
		},
	},
)

type CreateSellerDto struct {
	ShopName string `json:"shop_name"`
	PhotoUrl string `json:"photo_url"`
}

var GqlUpdateSellerDto = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UpdateSellerDto",
		Fields: graphql.Fields{
			"id":       {Type: graphql.ID},
			"shopName": {Type: graphql.String},
			"photoUrl": {Type: graphql.String},
		},
	},
)

type UpdateSellerDto struct {
	Id       uint32 `json:"id"`
	ShopName string `json:"shop_name,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
}

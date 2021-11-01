package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"time"
)

type SellerModel struct {
	gorm.Model
	ShopName     string
	PhotoUrl     string
	Bouquets     []BouquetModel `gorm:"foreignKey:SellerId"`
	SoldBouquets int
}

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
	ID           uint          `json:"id"`
	ShopName     string        `json:"shop_name"`
	PhotoUrl     string        `json:"photo_url"`
	CreatedDate  *time.Time    `json:"created_date"`
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
	ID       uint   `json:"id"`
	ShopName string `json:"shop_name,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
}

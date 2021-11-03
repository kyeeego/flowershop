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
			"id":           {Type: graphql.Int},
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
	ShopName     string        `json:"shopName"`
	PhotoUrl     string        `json:"photoUrl"`
	CreatedDate  *time.Time    `json:"createdDate"`
	Bouquets     []*BouquetDto `json:"bouquets"`
	SoldBouquets int           `json:"soldBouquets"`
}

type CreateSellerDto struct {
	ShopName string `json:"shopName"`
	PhotoUrl string `json:"photoUrl"`
}

type UpdateSellerDto struct {
	ID       uint   `json:"id"`
	ShopName string `json:"shopName,omitempty"`
	PhotoUrl string `json:"photoUrl,omitempty"`
}

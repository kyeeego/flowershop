package domain

import "github.com/graphql-go/graphql"


var GqlCustomer = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Customer",
		Fields: graphql.Fields{
			"id":        {Type: graphql.ID},
			"name":      {Type: graphql.String},
			"email":     {Type: graphql.String},
			"purchases": {Type: graphql.NewList(GqlPurchase)},
		},
	},
)

type CustomerDto struct {
	Id        uint32         `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Purchases []*PurchaseDto `json:"purchases"`
}

var GqlCreateCustomerDto = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreateCustomerDto",
		Fields: graphql.Fields{
			"name":  {Type: graphql.String},
			"email": {Type: graphql.String},
		},
	},
)

type CreateCustomerDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var GqlUpdateCustomerDto = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UpdateCustomerDto",
		Fields: graphql.Fields{
			"id":    {Type: graphql.ID},
			"name":  {Type: graphql.String},
			"email": {Type: graphql.String},
		},
	},
)

type UpdateCustomerDto struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

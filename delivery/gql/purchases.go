package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyeeego/flowershop/domain"
)

type purchase struct{}

func (purchase) initQueries() graphql.Fields {
	return graphql.Fields{
		"purchases": &graphql.Field{
			Type: graphql.NewList(domain.GqlPurchase),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
	}
}

func (purchase) initMutations() graphql.Fields {
	return graphql.Fields{
		"purchaseBouquet": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlDoPurchaseDto},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
	}
}

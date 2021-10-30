package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyeeego/flowershop/domain"
)

type Seller struct{}

func (Seller) initQueries() graphql.Fields {
	return graphql.Fields{
		"seller": &graphql.Field{
			Type: domain.GqlSeller,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
		"sellers": &graphql.Field{
			Type: graphql.NewList(domain.GqlSeller),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("not yet impl")
			},
		},
	}
}

func (Seller) initMutations() graphql.Fields {
	return graphql.Fields{
		"createSeller": &graphql.Field{
			Type: domain.GqlSeller,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlCreateSellerDto},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("not yet impl")
			},
		},
		"updateSeller": &graphql.Field{
			Type: domain.GqlSeller,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlUpdateSellerDto},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("not yet impl")
			},
		},
		"deleteSeller": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("not yet impl")
			},
		},
	}
}

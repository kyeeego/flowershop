package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyeeego/flowershop/domain"
)

type bouquet struct{}

func (bouquet) initQueries() graphql.Fields {
	return graphql.Fields{
		"bouquet": &graphql.Field{
			Type: domain.GqlBouquet,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
		"bouquets": &graphql.Field{
			Type: graphql.NewList(domain.GqlBouquet),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("not yet impl")
			},
		},
	}
}

func (bouquet) initMutations() graphql.Fields {
	return graphql.Fields{
		"createBouquet": &graphql.Field{
			Type: domain.GqlBouquet,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlCreateBouquetDto},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
		"updateBouquet": &graphql.Field{
			Type: domain.GqlSeller,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlUpdateBouquetDto},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
		"deleteBouquet": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
	}
}

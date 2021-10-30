package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyeeego/flowershop/domain"
)

type Customer struct{}

func (Customer) initQueries() graphql.Fields {
	return graphql.Fields{
		"customer": &graphql.Field{
			Type: domain.GqlCustomer,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
	}
}

func (Customer) initMutations() graphql.Fields {
	return graphql.Fields{
		"createCustomer": &graphql.Field{
			Type: domain.GqlCustomer,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlCreateCustomerDto},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet implemented")
			},
		},
		"updateCustomer": &graphql.Field{
			Type: domain.GqlCustomer,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlUpdateCustomerDto},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("not yet impl")
			},
		},
		"deleteCustomer": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				panic("Not yet impl")
			},
		},
	}
}

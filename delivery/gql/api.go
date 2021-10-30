package gql

import (
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/kyeeego/flowershop/utils"
)

type Endpoint interface {
	initQueries() graphql.Fields
	initMutations() graphql.Fields
}

type Api struct {
	Handler *gqlhandler.Handler
}

func New(endpoints []Endpoint) (*Api, error) {

	queries := graphql.Fields{}
	mutations := graphql.Fields{}
	for _, e := range endpoints {
		eq := e.initQueries()
		em := e.initMutations()

		queries = utils.MergeGqlFieldMaps(queries, eq)
		mutations = utils.MergeGqlFieldMaps(mutations, em)
	}

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(graphql.ObjectConfig{
				Fields: queries,
			}),
			Mutation: graphql.NewObject(graphql.ObjectConfig{
				Fields: mutations,
			}),
		},
	)
	if err != nil {
		return nil, err
	}

	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
		Pretty: true,
	})

	return &Api{
		Handler: h,
	}, nil
}

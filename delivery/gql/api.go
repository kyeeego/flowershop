package gql

import (
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/kyeeego/flowershop/service"
	"github.com/kyeeego/flowershop/utils"
	"net/http"
)

type Endpoint interface {
	initQueries() graphql.Fields
	initMutations() graphql.Fields
	initServices(serv *service.Service)
}

type Api struct {
	Handler http.Handler
}

func New(service *service.Service, endpoints ...Endpoint) (*Api, error) {

	queries := graphql.Fields{}
	mutations := graphql.Fields{}
	for _, e := range endpoints {
		eq := e.initQueries()
		em := e.initMutations()

		e.initServices(service)

		queries = utils.MergeGqlFieldMaps(queries, eq)
		mutations = utils.MergeGqlFieldMaps(mutations, em)
	}

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(graphql.ObjectConfig{
				Name:   "query",
				Fields: queries,
			}),
			Mutation: graphql.NewObject(graphql.ObjectConfig{
				Name:   "mutation",
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

	handler := http.NewServeMux()
	handler.Handle("/gql", h)

	return &Api{
		Handler: handler,
	}, nil
}

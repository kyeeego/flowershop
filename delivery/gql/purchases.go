package gql

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/service"
)

type Purchase struct{ service *service.Service }

func (p *Purchase) initServices(serv *service.Service) {
	p.service = serv
}

func (p Purchase) ResolveAll(params graphql.ResolveParams) (interface{}, error) {
	customerId, ok := params.Args["id"].(int)
	if !ok {
		return nil, errors.New("invalid id argument")
	}

	return p.service.Purchase.GetAll(uint(customerId))
}

func (p Purchase) ResolvePurchaseBouquet(params graphql.ResolveParams) (interface{}, error) {
	dto := domain.DoPurchaseDto{
		CustomerID: uint(params.Args["customerId"].(int)),
		BouquetID:  uint(params.Args["bouquetId"].(int)),
	}

	err := p.service.Purchase.Do(dto.CustomerID, dto.BouquetID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p Purchase) initQueries() graphql.Fields {
	return graphql.Fields{
		"purchases": &graphql.Field{
			Type: graphql.NewList(domain.GqlPurchase),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: p.ResolveAll,
		},
	}
}

func (p Purchase) initMutations() graphql.Fields {
	return graphql.Fields{
		"purchaseBouquet": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"customerId": {Type: graphql.Int},
				"bouquetId":  {Type: graphql.Int},
			},
			Resolve: p.ResolvePurchaseBouquet,
		},
	}
}

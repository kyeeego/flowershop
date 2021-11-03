package gql

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/service"
)

type Bouquet struct{ service *service.Service }

func (b *Bouquet) initServices(serv *service.Service) {
	b.service = serv
}

func (b Bouquet) ResolveOne(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, errors.New("invalid id argument")
	}

	return b.service.Bouquet.GetById(uint(id))
}

func (b Bouquet) ResolveAll(graphql.ResolveParams) (interface{}, error) {
	return b.service.Bouquet.GetAll()
}

func (b Bouquet) ResolveCreate(p graphql.ResolveParams) (interface{}, error) {
	dto := domain.CreateBouquetDto{
		Name:     p.Args["name"].(string),
		Price:    p.Args["price"].(float64),
		PhotoUrl: p.Args["photoUrl"].(string),
		SellerId: uint(p.Args["sellerId"].(int)),
	}

	return b.service.Bouquet.Create(dto)
}

func (b Bouquet) ResolveUpdate(p graphql.ResolveParams) (interface{}, error) {
	dto := domain.UpdateBouquetDto{
		ID:       uint(p.Args["id"].(int)),
		Name:     p.Args["name"].(string),
		Price:    p.Args["price"].(float64),
		PhotoUrl: p.Args["photoUrl"].(string),
	}

	return b.service.Bouquet.Update(dto)
}

func (b Bouquet) ResolveDelete(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, errors.New("invalid id argument")
	}

	err := b.service.Bouquet.Delete(uint(id))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (b Bouquet) initQueries() graphql.Fields {
	return graphql.Fields{
		"bouquet": &graphql.Field{
			Type: domain.GqlBouquet,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: b.ResolveOne,
		},
		"bouquets": &graphql.Field{
			Type:    graphql.NewList(domain.GqlBouquet),
			Resolve: b.ResolveAll,
		},
	}
}

func (b Bouquet) initMutations() graphql.Fields {
	return graphql.Fields{
		"createBouquet": &graphql.Field{
			Type: domain.GqlBouquet,
			Args: graphql.FieldConfigArgument{
				"name":     {Type: graphql.String},
				"price":    {Type: graphql.Float},
				"photoUrl": {Type: graphql.String},
				"sellerId": {Type: graphql.Int},
			},
			Resolve: b.ResolveCreate,
		},
		"updateBouquet": &graphql.Field{
			Type: domain.GqlSeller,
			Args: graphql.FieldConfigArgument{
				"id":       {Type: graphql.Int},
				"name":     {Type: graphql.String},
				"price":    {Type: graphql.Float},
				"photoUrl": {Type: graphql.String},
			},
			Resolve: b.ResolveUpdate,
		},
		"deleteBouquet": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: b.ResolveDelete,
		},
	}
}

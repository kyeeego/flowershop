package gql

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/service"
)

type Seller struct{ service *service.Service }

func (s *Seller) initServices(serv *service.Service) {
	s.service = serv
}

func (s Seller) ResolveOne(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(uint)
	if !ok {
		return nil, errors.New("invalid id argument")
	}

	return s.service.Seller.GetById(id)
}

func (s Seller) ResolveAll(graphql.ResolveParams) (interface{}, error) {
	return s.service.Seller.GetAll()
}

func (s Seller) ResolveCreate(p graphql.ResolveParams) (interface{}, error) {
	dto, ok := p.Args["in"].(domain.CreateSellerDto)
	if !ok {
		return nil, errors.New("invalid in argument")
	}

	return s.service.Seller.Create(dto)
}

func (s Seller) ResolveUpdate(p graphql.ResolveParams) (interface{}, error) {
	dto, ok := p.Args["in"].(domain.UpdateSellerDto)
	if !ok {
		return nil, errors.New("invalid in argument")
	}

	return s.service.Seller.Update(dto)
}

func (s Seller) ResolveDelete(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(uint)
	if !ok {
		return nil, errors.New("invalid id argument")
	}

	err := s.service.Seller.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s Seller) initQueries() graphql.Fields {
	return graphql.Fields{
		"seller": &graphql.Field{
			Type: domain.GqlSeller,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: s.ResolveOne,
		},
		"sellers": &graphql.Field{
			Type:    graphql.NewList(domain.GqlSeller),
			Resolve: s.ResolveAll,
		},
	}
}

func (s Seller) initMutations() graphql.Fields {
	return graphql.Fields{
		"createSeller": &graphql.Field{
			Type: domain.GqlSeller,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlCreateSellerDto},
			},
			Resolve: s.ResolveCreate,
		},
		"updateSeller": &graphql.Field{
			Type: domain.GqlSeller,
			Args: graphql.FieldConfigArgument{
				"in": &graphql.ArgumentConfig{Type: domain.GqlUpdateSellerDto},
			},
			Resolve: s.ResolveUpdate,
		},
		"deleteSeller": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: s.ResolveDelete,
		},
	}
}

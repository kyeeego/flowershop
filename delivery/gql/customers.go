package gql

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/kyeeego/flowershop/domain"
	"github.com/kyeeego/flowershop/service"
)

type Customer struct{ service *service.Service }

func (c *Customer) initServices(serv *service.Service) {
	c.service = serv
}

func (c Customer) ResolveOne(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, errors.New("invalid id argument")
	}

	return c.service.Customer.GetById(uint(id))
}

func (c Customer) ResolveCreate(p graphql.ResolveParams) (interface{}, error) {
	dto := domain.CreateCustomerDto{
		Name:  p.Args["name"].(string),
		Email: p.Args["email"].(string),
	}

	return c.service.Customer.Create(dto)
}

func (c Customer) ResolveUpdate(p graphql.ResolveParams) (interface{}, error) {
	dto := domain.UpdateCustomerDto{
		ID:    uint(p.Args["id"].(int)),
		Name:  p.Args["name"].(string),
		Email: p.Args["email"].(string),
	}

	return c.service.Customer.Update(dto)
}

func (c Customer) ResolveDelete(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, errors.New("invalid id argument")
	}

	err := c.service.Customer.Delete(uint(id))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c Customer) initQueries() graphql.Fields {
	return graphql.Fields{
		"customer": &graphql.Field{
			Type: domain.GqlCustomer,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: c.ResolveOne,
		},
	}
}

func (c Customer) initMutations() graphql.Fields {
	return graphql.Fields{
		"createCustomer": &graphql.Field{
			Type: domain.GqlCustomer,
			Args: graphql.FieldConfigArgument{
				"name":  {Type: graphql.String},
				"email": {Type: graphql.String},
			},
			Resolve: c.ResolveCreate,
		},
		"updateCustomer": &graphql.Field{
			Type: domain.GqlCustomer,
			Args: graphql.FieldConfigArgument{
				"id":    {Type: graphql.Int},
				"name":  {Type: graphql.String},
				"email": {Type: graphql.String},
			},
			Resolve: c.ResolveUpdate,
		},
		"deleteCustomer": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: c.ResolveDelete,
		},
	}
}

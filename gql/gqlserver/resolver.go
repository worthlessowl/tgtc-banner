package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/worthlessowl/tgtc-banner/backend/dictionary"
	"github.com/worthlessowl/tgtc-banner/backend/service"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// id, _ := p.Args["product_id"].(int)

		// update to use Usecase from previous session
		return dictionary.Product{}, nil
	}
}

func (r *Resolver) GetBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)

		result := service.GetBanner(id)
		return result, nil
	}
}

func (r *Resolver) GetBanners() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {

		result := service.GetBanners()
		return result, nil
	}
}

func (r *Resolver) CreateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		

		result := service.CreateBanner(p.Args)
		return result, nil
	}
}

func (r *Resolver) UpdateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		
		result := service.UpdateBanner(p.Args)
		return result, nil
	}
}

func (r *Resolver) DeleteBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)

		result := service.GetBanner(id)
		return result, nil
	}
}

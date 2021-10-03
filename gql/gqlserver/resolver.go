package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/worthlessowl/tgtc-banner/backend/dictionary"
	"github.com/worthlessowl/tgtc-banner/backend/service"
	"github.com/mitchellh/mapstructure"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

// func (r *Resolver) GetProduct() graphql.FieldResolveFn {
// 	return func(p graphql.ResolveParams) (interface{}, error) {
// 		// id, _ := p.Args["product_id"].(int)

// 		// update to use Usecase from previous session
// 		return dictionary.Product{}, err
// 	}
// }

func (r *Resolver) GetBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)

		result, err := service.GetBanner(id)
		return result, err
	}
}

func (r *Resolver) GetBanners() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {

		result, err := service.GetBanners()
		return result, err
	}
}

func (r *Resolver) CreateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		
		var banner dictionary.Banner
		decodeErr := mapstructure.Decode(p.Args, &banner)
		if decodeErr != nil {
			panic(decodeErr)
		}
		result, err := service.CreateBanner(banner)
		return result, err
	}
}

func (r *Resolver) UpdateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		var banner dictionary.Banner
		decodeErr := mapstructure.Decode(p.Args, &banner)
		if decodeErr != nil {
			panic(decodeErr)
		}
		result, err := service.UpdateBanner(banner)
		return result, err
	}
}

func (r *Resolver) DeleteBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)

		err := service.DeleteBanner(id)
		return nil, err
	}
}

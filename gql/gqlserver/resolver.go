package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/radityaqb/tgtc/backend/dictionary"
	"github.com/radityaqb/tgtc/backend/service"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["banner_id"].(int)

		// update to use Usecase from previous session
		return service.GetBanner(id)
	}
}

func (r *Resolver) GetBanners() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		
		return service.GetBanners()
	}
}
func (r *Resolver) GetUsers() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetUsers()
	}
}
// CREATE TABLE IF NOT EXISTS banner (banner_id SERIAL PRIMARY KEY, banner_name TEXT, banner_url TEXT, banner_imagesrc TEXT, 
//banner_startdate INT,banner_enddate INT, banner_maxtier INT, banner_mintier INT,banner_minbalance INT,banner_maxbalance INT,
//banner_mintokopoint INT,banner_maxtokopoint INT,banner_isactive BIT);
func (r *Resolver) CreateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["banner_name"].(string)
		url, _ := p.Args["banner_url"].(string)
		imagesrc, _ := p.Args["banner_imagesrc"].(string)
		startdate, _ := p.Args["banner_startdate"].(int64)
		enddate, _ := p.Args["banner_enddate"].(int64)
		maxtier, _ := p.Args["banner_maxtier"].(int64)
		mintier, _ := p.Args["banner_mintier"].(int64)
		minbalance, _ := p.Args["banner_minbalance"].(int64)
		maxbalance, _ := p.Args["banner_maxbalance"].(int64)
		mintokopoint, _ := p.Args["banner_mintokopoint"].(int64)
		maxtokopoint, _ := p.Args["banner_maxtokopoint"].(int64)
		isavtive, _ := p.Args["banner_isactive"].(bool)

		req := dictionary.Banner{
			Name:     name,
			Url: url,
			Imagesrc: expired,
			StartDate: startdate,
			EndDate: enddate,
			MinTier: mintier,
			MaxTier: maxtier,
			MinBalance: minbalance,
			MaxBalance: maxbalance,
			MinTokopoint: mintokopoint,
			MaxTokopoint: maxtokopoint,
			isActive: isavtive,
		}
		_, err := service.UpdateBanner(req)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (r *Resolver) UpdateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["banner_name"].(string)
		url, _ := p.Args["banner_url"].(string)
		imagesrc, _ := p.Args["banner_imagesrc"].(string)
		sartdate, _ := p.Args["banner_startdate"].(int64)
		enddate, _ := p.Args["banner_enddate"].(int64)
		maxtier, _ := p.Args["banner_maxtier"].(int64)
		mintier, _ := p.Args["banner_mintier"].(int64)
		minbalance, _ := p.Args["banner_minbalance"].(int64)
		maxbalance, _ := p.Args["banner_maxbalance"].(int64)
		mintokopoint, _ := p.Args["banner_mintokopoint"].(int64)
		maxtokopoint, _ := p.Args["banner_maxtokopoint"].(int64)
		isavtive, _ := p.Args["banner_isactive"].(bool)

		req := dictionary.Banner{
			Name:     name,
			Url: url,
			Imagesrc: expired,
			StartDate: startdate,
			EndDate: enddate,
			MinTier: mintier,
			MaxTier: maxtier,
			MinBalance: minbalance,
			MaxBalance: maxbalance,
			MinTokopoint: mintokopoint,
			MaxTokopoint: maxtokopoint,
			isActive: isavtive,
		}
		_, err := service.CreateBanner(req)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}
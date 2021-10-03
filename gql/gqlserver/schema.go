package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	dataResolver *Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(dr *Resolver) *SchemaWrapper {
	s.dataResolver = dr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Get Banner",
			Description: "All query related to getting Banner data",
			Fields: graphql.Fields{
				"BannerDetail": &graphql.Field{
					Type:        BannerType,
					Description: "Get Banner by ID",
					Args: graphql.FieldConfigArgument{
						"banner_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.dataResolver.GetBanner(),
				},
				"Users": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "Get Users",
					Resolve:     s.dataResolver.GetUsers(),
				},
				"Banners": &graphql.Field{
					Type:        graphql.NewList(BannerType),
					Description: "Get banners",
					Resolve:     s.dataResolver.GetBanners(),
				},
			},
		}),
		// uncomment this and add objects for mutation
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Mutation Banner",
			Description: "All query related to modify banner data",
			Fields: graphql.Fields{
				"CreateBanner": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Create banner",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"imgSrc": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"startDate": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"endDate": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"minTier": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"maxTier": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"minBalance": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"maxBalance": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"minTokopoint": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"maxTokopoint": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"isActive": &graphql.ArgumentConfig{
							Type: graphql.Boolean,
						},
					},
					Resolve: s.dataResolver.CreateBanner(),
				},
				"UpdateBanner": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Update banner",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"imgSrc": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"startDate": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"endDate": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"minTier": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"maxTier": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"minBalance": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"maxBalance": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"minTokopoint": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"maxTokopoint": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"isActive": &graphql.ArgumentConfig{
							Type: graphql.Boolean,
						},
					},
					Resolve: s.dataResolver.UpdateBanner(),
				},
			},	
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
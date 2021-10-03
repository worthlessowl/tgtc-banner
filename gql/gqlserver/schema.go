package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	resolver *Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithResolver(pr *Resolver) *SchemaWrapper {
	s.resolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		// Query: graphql.NewObject(graphql.ObjectConfig{
		// 	Name:        "ProductGetter",
		// 	Description: "All query related to getting product data",
		// 	Fields: graphql.Fields{
		// 		"ProductDetail": &graphql.Field{
		// 			Type:        ProductType,
		// 			Description: "Get product by ID",
		// 			Args: graphql.FieldConfigArgument{
		// 				"product_id": &graphql.ArgumentConfig{
		// 					Type: graphql.Int,
		// 				},
		// 			},
		// 			Resolve: s.productResolver.GetProduct(),
		// 		},
		// 	},
		// }),

		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:		"BannerGetter",
			Description: "All query related to getting banner data",
			Fields: graphql.Fields{
				"BannerDetail": &graphql.Field{
					Type: BannerType,
					Description: "Get Banner by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NonNull(graphql.Int),
						},
					},
					Resolve: s.resolver.GetBanner(),
				},
				"AllBannerDetail": &graphql.Field{
					Type: graphql.NewList(BannerType),
					Description: "Get All Banner",
					Resolve: s.resolver.GetBanners(),
				},
				"AllUser": &graphql.Field{
					Type: graphql.NewList(UserType),
					Description: "Get All User",
					Resolve: s.resolver.GetUsers(),
				},
				"UserDetail": &graphql.Field{
					Type: UserType,
					Description: "Get User by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.resolver.GetUser(),
				},
			},
		}),

		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "BannerMutation",
			Description: "All query related to editing banner data",
			Fields: graphql.Fields{
				"CreateBanner": &graphql.Field{
					Type: BannerType,
					Description: "Create Banner",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NonNull(graphql.String),
						},
						"url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"imgSrc": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"startDate": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"endDate": &graphql.ArgumentConfig{
							Type: graphql.Int,
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
							Type: graphql.Int,
						},
					},
					Resolve: s.resolver.CreateBanner(),
				},
				"EditBanner": &graphql.Field{
					Type: BannerType,
					Description: "Edit Banner",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NonNull(graphql.String),
						},
						"id": &graphql.ArgumentConfig{
							Type: graphql.NonNull(graphql.Int),
						},
						"url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"imgSrc": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"startDate": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"endDate": &graphql.ArgumentConfig{
							Type: graphql.Int,
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
							Type: graphql.Int,
						},
					},
					Resolve: s.resolver.UpdateBanner(),
				},
				"DeleteBanner": &graphql.Field{
					Type: BannerType,
					Description: "Delete Banner by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NonNull(graphql.Int),
						},
					},
					Resolve: s.resolver.DeleteBanner(),
				},
				"CreateUser": &graphql.Field{
					Type: UserType,
					Description: "Create User",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"balance": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"tokopoint": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"tier": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"location": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"bannerList": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.resolver.CreateUser(),
				},
				"EditUser": &graphql.Field{
					Type: UserType,
					Description: "Edit User",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"balance": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"tokopoint": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"tier": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"location": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"bannerList": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.resolver.UpdateUser(),
				},
				"DeleteUser": &graphql.Field{
					Type: UserType,
					Description: "Delete User by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.resolver.DeleteUser(),
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

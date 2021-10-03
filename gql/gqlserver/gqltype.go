package gqlserver

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Product",
		Description: "Detail of the product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var BannerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:		"Banner",
		Description:"Banner for user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"imgSrc": &graphql.Field{
				Type: graphql.String,
			},
			"startDate": &graphql.Field{
				Type: graphql.String,
			},
			"endDate": &graphql.Field{
				Type: graphql.String,
			},
			"minTier": &graphql.Field{
				Type: graphql.String,
			},
			"maxTier": &graphql.Field{
				Type: graphql.String,
			},
			"minBalance": &graphql.Field{
				Type: graphql.Int,
			},
			"maxBalance": &graphql.Field{
				Type: graphql.Int,
			},
			"minTokopoint": &graphql.Field{
				Type: graphql.Int,
			},
			"maxTokopoint": &graphql.Field{
				Type: graphql.Int,
			},
			"isactive": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

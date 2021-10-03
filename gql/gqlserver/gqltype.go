package gqlserver

import "github.com/graphql-go/graphql"

var BannerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:		"Banner",
		Description:"Banner Detail"
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
		}
	}
)
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "User",
		Description: "Detail of the user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"balance": &graphql.Field{
				Type: graphql.String,
			},
			"tokopoint": &graphql.Field{
				Type: graphql.String,
			},
			"tier": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
			"bannerList": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
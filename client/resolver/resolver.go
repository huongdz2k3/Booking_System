package resolver

import (
	directives "customer/directive"
	"customer/ent"
	generated "customer/graphql"
	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	c := generated.Config{
		Resolvers: &Resolver{client},
	}
	c.Directives.Auth = directives.Auth
	c.Directives.HasRole = directives.HasRole
	return generated.NewExecutableSchema(c)
}

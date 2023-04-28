package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"customer/ent"
	graphql1 "customer/graphql"
	"fmt"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input ent.RegisterInput) (*ent.Jwt, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input ent.LoginInput) (*ent.Jwt, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context) ([]*ent.Customer, error) {
	panic(fmt.Errorf("not implemented: Customers - customers"))
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

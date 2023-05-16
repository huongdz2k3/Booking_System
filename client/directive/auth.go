package directives

import (
	"context"
	"customer/ent/customer"
	"customer/internal/utils"
	"customer/middleware"
	"github.com/99designs/gqlgen/graphql"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := middleware.CtxValue(ctx)
	if tokenData == nil {
		return nil, utils.WrapGQLUnauthorizedError(ctx)
	}
	return next(ctx)
}

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role []customer.Role) (interface{}, error) {
	tokenData := middleware.CtxValue(ctx)
	if tokenData == nil {
		return nil, utils.WrapGQLUnauthorizedError(ctx)
	}
	if !Contains(tokenData.Role, role) {
		return nil, utils.WrapGQLUnauthorizedError(ctx)
	}
	return next(ctx)

}

func Contains(s string, e []customer.Role) bool {
	for _, a := range e {
		if a.String() == s {
			return true
		}
	}
	return false
}

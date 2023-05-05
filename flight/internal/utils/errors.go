package utils

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapGQLError(ctx context.Context, message string, code string) error {
	// Define a map of gRPC error codes that correspond to the GraphQL error codes
	errorCodeMap := map[string]codes.Code{
		"NOT_FOUND":             codes.NotFound,
		"BAD_REQUEST":           codes.InvalidArgument,
		"INTERNAL_SERVER_ERROR": codes.Internal,
		"UNAUTHORIZED":          codes.Unauthenticated,
		// Add more mappings as needed
	}

	// Check if the provided code exists in the mapping
	gRPCCode, ok := errorCodeMap[code]
	if !ok {
		// Default to an unknown error code
		gRPCCode = codes.Unknown
	}

	// Create a new gRPC status with the provided message and code
	st := status.New(gRPCCode, message)

	//// Add the path to the error details if it's available
	//if ctx != nil {
	//	path := graphql.GetPath(ctx)
	//	st, _ = st.WithDetails(&errorDetails{
	//		Path: path,
	//	})
	//}

	return st.Err()
}

//// Define a custom error details struct that includes the GraphQL error path
//type errorDetails struct {
//	Path graphql.ResponsePath `json:"path,omitempty"`
//}

//// Implement the grpc details method to return the custom error details
//func (e *errorDetails) GRPCDetails() interface{} {
//	return e
//}

func WrapInternalError(ctx context.Context) error {
	return WrapGQLError(ctx, "Internal Server Error", "INTERNAL_SERVER_ERROR")
}

func WrapBadRequestError(ctx context.Context, format string, args ...interface{}) error {
	return WrapGQLError(ctx, fmt.Sprintf(format, args...), "BAD_REQUEST")
}

func WrapUnauthorizedError(ctx context.Context) error {
	return WrapGQLError(ctx, "Unauthorized Request", "UNAUTHORIZED")
}

func WrapNotFoundError(ctx context.Context) error {
	return WrapGQLError(ctx, "Not Found", "NOT_FOUND")
}

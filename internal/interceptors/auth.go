package interceptors

import (
	"connectrpc.com/connect"
	"context"
	"encoding/base64"
	"errors"
	"github.com/BugattiBoys/bazaar/internal/server"
	"github.com/jackc/pgx/v5/pgtype"
	"strings"
)

type store interface {
	FindServer(ctx context.Context, arg server.FindServerParams) (server.Server, error)
}

func NewAuthInterceptor(store store) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			// Ignore authorization on any requests to create a server
			if req.Spec().Procedure == "/api.v1.Servers/CreateServer" {
				return next(ctx, req)
			}
			authHeader := req.Header().Get("Authorization")
			if authHeader == "" {
				return unauthenticated(errors.New("no authentication provided"))
			}

			creds, found := strings.CutPrefix(authHeader, "Basic ")
			if !found {
				return unauthenticated(errors.New("invalid authentication"))
			}

			sDec, err := base64.StdEncoding.DecodeString(creds)
			if err != nil {
				return unauthenticated(errors.New("invalid authentication"))
			}

			chunks := strings.Split(string(sDec), ":")
			if len(chunks) != 2 {
				return unauthenticated(errors.New("invalid authentication"))
			}

			serverId := pgtype.UUID{}
			if err := serverId.Scan(chunks[0]); err != nil {
				return unauthenticated(errors.New("invalid authentication"))
			}
			apiKey := pgtype.UUID{}
			if err := apiKey.Scan(chunks[1]); err != nil {
				return unauthenticated(errors.New("invalid authentication"))
			}

			_, err = store.FindServer(ctx, server.FindServerParams{
				ID:     serverId,
				ApiKey: apiKey,
			})
			if err != nil {
				return unauthenticated(errors.New("invalid authentication"))
			}

			ctx = context.WithValue(ctx, "server_id", serverId)
			return next(ctx, req)
		}
	}
	return interceptor
}

func unauthenticated(err error) (connect.AnyResponse, error) {
	return nil, connect.NewError(
		connect.CodeUnauthenticated,
		err,
	)
}

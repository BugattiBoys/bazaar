package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"context"
	"github.com/BugattiBoys/bazaar/gen/api/v1/apiv1connect"
	"github.com/BugattiBoys/bazaar/internal/interceptors"
	"github.com/BugattiBoys/bazaar/internal/server"
	"github.com/BugattiBoys/bazaar/internal/trade"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func init() {

}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:password@localhost:5432/bazaar_development")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	reflector := grpcreflect.NewStaticReflector(
		apiv1connect.ServersName,
		apiv1connect.TradesName,
	)
	middleware := connect.WithInterceptors(
		interceptors.NewAuthInterceptor(server.New(conn)),
	)
	mux := http.NewServeMux()
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(apiv1connect.NewServersHandler(server.NewServerServer(conn), middleware))
	mux.Handle(apiv1connect.NewTradesHandler(trade.NewServer(conn)))
	if err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}

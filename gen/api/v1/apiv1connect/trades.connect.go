// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/trades.proto

package apiv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/BugattiBoys/bazaar/gen/api/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TradesName is the fully-qualified name of the Trades service.
	TradesName = "api.v1.Trades"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TradesCreateTradeProcedure is the fully-qualified name of the Trades's CreateTrade RPC.
	TradesCreateTradeProcedure = "/api.v1.Trades/CreateTrade"
	// TradesWatchTradeProcedure is the fully-qualified name of the Trades's WatchTrade RPC.
	TradesWatchTradeProcedure = "/api.v1.Trades/WatchTrade"
	// TradesCancelTradeProcedure is the fully-qualified name of the Trades's CancelTrade RPC.
	TradesCancelTradeProcedure = "/api.v1.Trades/CancelTrade"
	// TradesFulfillTradeProcedure is the fully-qualified name of the Trades's FulfillTrade RPC.
	TradesFulfillTradeProcedure = "/api.v1.Trades/FulfillTrade"
	// TradesGetOpenTradesProcedure is the fully-qualified name of the Trades's GetOpenTrades RPC.
	TradesGetOpenTradesProcedure = "/api.v1.Trades/GetOpenTrades"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	tradesServiceDescriptor             = v1.File_api_v1_trades_proto.Services().ByName("Trades")
	tradesCreateTradeMethodDescriptor   = tradesServiceDescriptor.Methods().ByName("CreateTrade")
	tradesWatchTradeMethodDescriptor    = tradesServiceDescriptor.Methods().ByName("WatchTrade")
	tradesCancelTradeMethodDescriptor   = tradesServiceDescriptor.Methods().ByName("CancelTrade")
	tradesFulfillTradeMethodDescriptor  = tradesServiceDescriptor.Methods().ByName("FulfillTrade")
	tradesGetOpenTradesMethodDescriptor = tradesServiceDescriptor.Methods().ByName("GetOpenTrades")
)

// TradesClient is a client for the api.v1.Trades service.
type TradesClient interface {
	// Create a new trade (either buy, or sell), registered to the server that requested it
	CreateTrade(context.Context, *connect.Request[v1.CreateTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error)
	// Stream events for an open trade
	WatchTrade(context.Context, *connect.Request[v1.WatchTradeRequest]) (*connect.ServerStreamForClient[v1.WatchTradeResponse], error)
	// Cancel an open trade
	CancelTrade(context.Context, *connect.Request[v1.CancelTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error)
	// Fulfill an open trade. This will trigger both servers to
	// At the moment, this is only buying with coins, or selling for coins
	FulfillTrade(context.Context, *connect.Request[v1.FulfillTradeRequest]) (*connect.Response[v1.FulfillTradeResponse], error)
	GetOpenTrades(context.Context, *connect.Request[v1.GetOpenTradesRequest]) (*connect.Response[v1.GetOpenTradesResponse], error)
}

// NewTradesClient constructs a client for the api.v1.Trades service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTradesClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TradesClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tradesClient{
		createTrade: connect.NewClient[v1.CreateTradeRequest, v1.CreateTradeResponse](
			httpClient,
			baseURL+TradesCreateTradeProcedure,
			connect.WithSchema(tradesCreateTradeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		watchTrade: connect.NewClient[v1.WatchTradeRequest, v1.WatchTradeResponse](
			httpClient,
			baseURL+TradesWatchTradeProcedure,
			connect.WithSchema(tradesWatchTradeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		cancelTrade: connect.NewClient[v1.CancelTradeRequest, v1.CreateTradeResponse](
			httpClient,
			baseURL+TradesCancelTradeProcedure,
			connect.WithSchema(tradesCancelTradeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		fulfillTrade: connect.NewClient[v1.FulfillTradeRequest, v1.FulfillTradeResponse](
			httpClient,
			baseURL+TradesFulfillTradeProcedure,
			connect.WithSchema(tradesFulfillTradeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getOpenTrades: connect.NewClient[v1.GetOpenTradesRequest, v1.GetOpenTradesResponse](
			httpClient,
			baseURL+TradesGetOpenTradesProcedure,
			connect.WithSchema(tradesGetOpenTradesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// tradesClient implements TradesClient.
type tradesClient struct {
	createTrade   *connect.Client[v1.CreateTradeRequest, v1.CreateTradeResponse]
	watchTrade    *connect.Client[v1.WatchTradeRequest, v1.WatchTradeResponse]
	cancelTrade   *connect.Client[v1.CancelTradeRequest, v1.CreateTradeResponse]
	fulfillTrade  *connect.Client[v1.FulfillTradeRequest, v1.FulfillTradeResponse]
	getOpenTrades *connect.Client[v1.GetOpenTradesRequest, v1.GetOpenTradesResponse]
}

// CreateTrade calls api.v1.Trades.CreateTrade.
func (c *tradesClient) CreateTrade(ctx context.Context, req *connect.Request[v1.CreateTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error) {
	return c.createTrade.CallUnary(ctx, req)
}

// WatchTrade calls api.v1.Trades.WatchTrade.
func (c *tradesClient) WatchTrade(ctx context.Context, req *connect.Request[v1.WatchTradeRequest]) (*connect.ServerStreamForClient[v1.WatchTradeResponse], error) {
	return c.watchTrade.CallServerStream(ctx, req)
}

// CancelTrade calls api.v1.Trades.CancelTrade.
func (c *tradesClient) CancelTrade(ctx context.Context, req *connect.Request[v1.CancelTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error) {
	return c.cancelTrade.CallUnary(ctx, req)
}

// FulfillTrade calls api.v1.Trades.FulfillTrade.
func (c *tradesClient) FulfillTrade(ctx context.Context, req *connect.Request[v1.FulfillTradeRequest]) (*connect.Response[v1.FulfillTradeResponse], error) {
	return c.fulfillTrade.CallUnary(ctx, req)
}

// GetOpenTrades calls api.v1.Trades.GetOpenTrades.
func (c *tradesClient) GetOpenTrades(ctx context.Context, req *connect.Request[v1.GetOpenTradesRequest]) (*connect.Response[v1.GetOpenTradesResponse], error) {
	return c.getOpenTrades.CallUnary(ctx, req)
}

// TradesHandler is an implementation of the api.v1.Trades service.
type TradesHandler interface {
	// Create a new trade (either buy, or sell), registered to the server that requested it
	CreateTrade(context.Context, *connect.Request[v1.CreateTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error)
	// Stream events for an open trade
	WatchTrade(context.Context, *connect.Request[v1.WatchTradeRequest], *connect.ServerStream[v1.WatchTradeResponse]) error
	// Cancel an open trade
	CancelTrade(context.Context, *connect.Request[v1.CancelTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error)
	// Fulfill an open trade. This will trigger both servers to
	// At the moment, this is only buying with coins, or selling for coins
	FulfillTrade(context.Context, *connect.Request[v1.FulfillTradeRequest]) (*connect.Response[v1.FulfillTradeResponse], error)
	GetOpenTrades(context.Context, *connect.Request[v1.GetOpenTradesRequest]) (*connect.Response[v1.GetOpenTradesResponse], error)
}

// NewTradesHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTradesHandler(svc TradesHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tradesCreateTradeHandler := connect.NewUnaryHandler(
		TradesCreateTradeProcedure,
		svc.CreateTrade,
		connect.WithSchema(tradesCreateTradeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tradesWatchTradeHandler := connect.NewServerStreamHandler(
		TradesWatchTradeProcedure,
		svc.WatchTrade,
		connect.WithSchema(tradesWatchTradeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tradesCancelTradeHandler := connect.NewUnaryHandler(
		TradesCancelTradeProcedure,
		svc.CancelTrade,
		connect.WithSchema(tradesCancelTradeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tradesFulfillTradeHandler := connect.NewUnaryHandler(
		TradesFulfillTradeProcedure,
		svc.FulfillTrade,
		connect.WithSchema(tradesFulfillTradeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tradesGetOpenTradesHandler := connect.NewUnaryHandler(
		TradesGetOpenTradesProcedure,
		svc.GetOpenTrades,
		connect.WithSchema(tradesGetOpenTradesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.Trades/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TradesCreateTradeProcedure:
			tradesCreateTradeHandler.ServeHTTP(w, r)
		case TradesWatchTradeProcedure:
			tradesWatchTradeHandler.ServeHTTP(w, r)
		case TradesCancelTradeProcedure:
			tradesCancelTradeHandler.ServeHTTP(w, r)
		case TradesFulfillTradeProcedure:
			tradesFulfillTradeHandler.ServeHTTP(w, r)
		case TradesGetOpenTradesProcedure:
			tradesGetOpenTradesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTradesHandler returns CodeUnimplemented from all methods.
type UnimplementedTradesHandler struct{}

func (UnimplementedTradesHandler) CreateTrade(context.Context, *connect.Request[v1.CreateTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.Trades.CreateTrade is not implemented"))
}

func (UnimplementedTradesHandler) WatchTrade(context.Context, *connect.Request[v1.WatchTradeRequest], *connect.ServerStream[v1.WatchTradeResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.Trades.WatchTrade is not implemented"))
}

func (UnimplementedTradesHandler) CancelTrade(context.Context, *connect.Request[v1.CancelTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.Trades.CancelTrade is not implemented"))
}

func (UnimplementedTradesHandler) FulfillTrade(context.Context, *connect.Request[v1.FulfillTradeRequest]) (*connect.Response[v1.FulfillTradeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.Trades.FulfillTrade is not implemented"))
}

func (UnimplementedTradesHandler) GetOpenTrades(context.Context, *connect.Request[v1.GetOpenTradesRequest]) (*connect.Response[v1.GetOpenTradesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.Trades.GetOpenTrades is not implemented"))
}
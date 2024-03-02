// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/servers.proto

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
	// ServersName is the fully-qualified name of the Servers service.
	ServersName = "api.v1.Servers"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServersCreateServerProcedure is the fully-qualified name of the Servers's CreateServer RPC.
	ServersCreateServerProcedure = "/api.v1.Servers/CreateServer"
	// ServersGetServerProcedure is the fully-qualified name of the Servers's GetServer RPC.
	ServersGetServerProcedure = "/api.v1.Servers/GetServer"
	// ServersPingProcedure is the fully-qualified name of the Servers's Ping RPC.
	ServersPingProcedure = "/api.v1.Servers/Ping"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serversServiceDescriptor            = v1.File_api_v1_servers_proto.Services().ByName("Servers")
	serversCreateServerMethodDescriptor = serversServiceDescriptor.Methods().ByName("CreateServer")
	serversGetServerMethodDescriptor    = serversServiceDescriptor.Methods().ByName("GetServer")
	serversPingMethodDescriptor         = serversServiceDescriptor.Methods().ByName("Ping")
)

// ServersClient is a client for the api.v1.Servers service.
type ServersClient interface {
	// Servers generate a unique UUID, and request a new slot. Bazaar then
	// generates an API key, and returns both to the server. The server
	// is responsible for saving both of these for further requests
	CreateServer(context.Context, *connect.Request[v1.CreateServerRequest]) (*connect.Response[v1.CreateServerResponse], error)
	// Get the currently authenticated server
	GetServer(context.Context, *connect.Request[v1.GetServerRequest]) (*connect.Response[v1.GetServerResponse], error)
	// Ping endpoint. If a server hasn't pinged for X amount of times,
	// we disable its trades until it comes back online
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
}

// NewServersClient constructs a client for the api.v1.Servers service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServersClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServersClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serversClient{
		createServer: connect.NewClient[v1.CreateServerRequest, v1.CreateServerResponse](
			httpClient,
			baseURL+ServersCreateServerProcedure,
			connect.WithSchema(serversCreateServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getServer: connect.NewClient[v1.GetServerRequest, v1.GetServerResponse](
			httpClient,
			baseURL+ServersGetServerProcedure,
			connect.WithSchema(serversGetServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ping: connect.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+ServersPingProcedure,
			connect.WithSchema(serversPingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serversClient implements ServersClient.
type serversClient struct {
	createServer *connect.Client[v1.CreateServerRequest, v1.CreateServerResponse]
	getServer    *connect.Client[v1.GetServerRequest, v1.GetServerResponse]
	ping         *connect.Client[v1.PingRequest, v1.PingResponse]
}

// CreateServer calls api.v1.Servers.CreateServer.
func (c *serversClient) CreateServer(ctx context.Context, req *connect.Request[v1.CreateServerRequest]) (*connect.Response[v1.CreateServerResponse], error) {
	return c.createServer.CallUnary(ctx, req)
}

// GetServer calls api.v1.Servers.GetServer.
func (c *serversClient) GetServer(ctx context.Context, req *connect.Request[v1.GetServerRequest]) (*connect.Response[v1.GetServerResponse], error) {
	return c.getServer.CallUnary(ctx, req)
}

// Ping calls api.v1.Servers.Ping.
func (c *serversClient) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// ServersHandler is an implementation of the api.v1.Servers service.
type ServersHandler interface {
	// Servers generate a unique UUID, and request a new slot. Bazaar then
	// generates an API key, and returns both to the server. The server
	// is responsible for saving both of these for further requests
	CreateServer(context.Context, *connect.Request[v1.CreateServerRequest]) (*connect.Response[v1.CreateServerResponse], error)
	// Get the currently authenticated server
	GetServer(context.Context, *connect.Request[v1.GetServerRequest]) (*connect.Response[v1.GetServerResponse], error)
	// Ping endpoint. If a server hasn't pinged for X amount of times,
	// we disable its trades until it comes back online
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
}

// NewServersHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServersHandler(svc ServersHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serversCreateServerHandler := connect.NewUnaryHandler(
		ServersCreateServerProcedure,
		svc.CreateServer,
		connect.WithSchema(serversCreateServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serversGetServerHandler := connect.NewUnaryHandler(
		ServersGetServerProcedure,
		svc.GetServer,
		connect.WithSchema(serversGetServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	serversPingHandler := connect.NewUnaryHandler(
		ServersPingProcedure,
		svc.Ping,
		connect.WithSchema(serversPingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.Servers/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServersCreateServerProcedure:
			serversCreateServerHandler.ServeHTTP(w, r)
		case ServersGetServerProcedure:
			serversGetServerHandler.ServeHTTP(w, r)
		case ServersPingProcedure:
			serversPingHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServersHandler returns CodeUnimplemented from all methods.
type UnimplementedServersHandler struct{}

func (UnimplementedServersHandler) CreateServer(context.Context, *connect.Request[v1.CreateServerRequest]) (*connect.Response[v1.CreateServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.Servers.CreateServer is not implemented"))
}

func (UnimplementedServersHandler) GetServer(context.Context, *connect.Request[v1.GetServerRequest]) (*connect.Response[v1.GetServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.Servers.GetServer is not implemented"))
}

func (UnimplementedServersHandler) Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.Servers.Ping is not implemented"))
}
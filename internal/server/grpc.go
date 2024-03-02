package server

import (
	"connectrpc.com/connect"
	"context"
	apiv1 "github.com/BugattiBoys/bazaar/gen/api/v1"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ServerServer struct {
	Q *Queries
}

func NewServerServer(db *pgx.Conn) *ServerServer {
	return &ServerServer{Q: New(db)}
}

func (s *ServerServer) CreateServer(
	ctx context.Context,
	req *connect.Request[apiv1.CreateServerRequest],
) (*connect.Response[apiv1.CreateServerResponse], error) {
	serverId := pgtype.UUID{}
	if err := serverId.Scan(req.Msg.ServerId); err != nil {
		return nil, err
	}

	server, err := s.Q.CreateServer(ctx, serverId)
	if err != nil {
		return nil, err
	}
	apiKey, _ := server.ApiKey.Value()
	return connect.NewResponse(&apiv1.CreateServerResponse{
		ApiKey: apiKey.(string),
	}), nil
}

func (s *ServerServer) GetServer(
	ctx context.Context,
	req *connect.Request[apiv1.GetServerRequest],
) (*connect.Response[apiv1.GetServerResponse], error) {
	server, err := s.Q.FindServer(ctx, FindServerParams{
		ID: ctx.Value("server_id").(pgtype.UUID),
	})
	if err != nil {
		return nil, err
	}
	id, _ := server.ID.Value()
	return connect.NewResponse(&apiv1.GetServerResponse{
		Id:        id.(string),
		IpAddress: server.IpAddress.String(),
	}), nil
}

func (s *ServerServer) Ping(
	ctx context.Context,
	req *connect.Request[apiv1.PingRequest],
) (*connect.Response[apiv1.PingResponse], error) {
	serverId := ctx.Value("server_id").(pgtype.UUID)
	err := s.Q.PingServer(ctx, serverId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&apiv1.PingResponse{}), nil
}

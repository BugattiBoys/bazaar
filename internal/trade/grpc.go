package trade

import (
	"connectrpc.com/connect"
	"context"
	v1 "github.com/BugattiBoys/bazaar/gen/api/v1"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Server struct {
	Q *Queries
}

func NewServer(db *pgx.Conn) *Server {
	return &Server{Q: New(db)}
}

// Create a new trade (either buy, or sell), registered to the server that requested it
func (s Server) CreateTrade(ctx context.Context, req *connect.Request[v1.CreateTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error) {
	return nil, nil
}

// Stream events for an open trade
func (s *Server) WatchTrade(ctx context.Context, req *connect.Request[v1.WatchTradeRequest], stream *connect.ServerStream[v1.WatchTradeResponse]) error {
	return nil
}

// Cancel an open trade
func (s *Server) CancelTrade(ctx context.Context, req *connect.Request[v1.CancelTradeRequest]) (*connect.Response[v1.CreateTradeResponse], error) {
	return nil, nil
}

// Fulfill an open trade. This will trigger both servers to
// At the moment, this is only buying with coins, or selling for coins
func (s *Server) FulfillTrade(ctx context.Context, req *connect.Request[v1.FulfillTradeRequest]) (*connect.Response[v1.FulfillTradeResponse], error) {
	return nil, nil
}

func (s *Server) GetOpenTrades(ctx context.Context, req *connect.Request[v1.GetOpenTradesRequest]) (*connect.Response[v1.GetOpenTradesResponse], error) {
	var initiatingServerId [16]byte
	copy(initiatingServerId[:], req.Msg.InitiatingServerId)
	trades, err := s.Q.GetServerOpenTradesWithStatus(context.TODO(), GetServerOpenTradesWithStatusParams{
		InitiatingServer: pgtype.UUID{Bytes: initiatingServerId},
		Status:           pgtype.Text{String: req.Msg.Status},
	})

	if err != nil {
		return nil, err
	}

	var result []*v1.Trade
	for _, trade := range trades {
		result = append(result, buildTradeResponse(trade))
	}

	return connect.NewResponse(&v1.GetOpenTradesResponse{
		Trades: result,
	}), nil
}

func buildTradeResponse(input Trade) *v1.Trade {
	return &v1.Trade{
		Status:           input.Status.String,
		Type:             input.Type.String,
		InitiatingServer: string(input.InitiatingServer.Bytes[:]),
		Item:             input.Item.String,
		Quantity:         input.TradeQuantity.Int32,
		Price:            input.Price.Int32,
		Fulfilled:        input.FulfilledQuantity.Int32,
	}
}

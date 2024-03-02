// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package trade

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getServerOpenTradesWithStatus = `-- name: GetServerOpenTradesWithStatus :many
SELECT id, status, type, initiating_server, item, trade_quantity, price, fulfilled_quantity, created_at, fulfilled_at, cancelled_at, updated_at FROM trades WHERE initiating_server = $1 AND status = $2
`

type GetServerOpenTradesWithStatusParams struct {
	InitiatingServer pgtype.UUID
	Status           pgtype.Text
}

func (q *Queries) GetServerOpenTradesWithStatus(ctx context.Context, arg GetServerOpenTradesWithStatusParams) ([]Trade, error) {
	rows, err := q.db.Query(ctx, getServerOpenTradesWithStatus, arg.InitiatingServer, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Trade
	for rows.Next() {
		var i Trade
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.Type,
			&i.InitiatingServer,
			&i.Item,
			&i.TradeQuantity,
			&i.Price,
			&i.FulfilledQuantity,
			&i.CreatedAt,
			&i.FulfilledAt,
			&i.CancelledAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
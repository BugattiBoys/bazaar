// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package trade

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Trade struct {
	ID                pgtype.UUID
	Status            pgtype.Text
	Type              pgtype.Text
	InitiatingServer  pgtype.UUID
	Item              pgtype.Text
	TradeQuantity     pgtype.Int4
	Price             pgtype.Int4
	FulfilledQuantity pgtype.Int4
	CreatedAt         pgtype.Timestamp
	FulfilledAt       pgtype.Timestamp
	CancelledAt       pgtype.Timestamp
	UpdatedAt         pgtype.Timestamp
}

type TradeFulfillment struct {
	ID                pgtype.UUID
	TradeID           pgtype.UUID
	FulfillingServer  pgtype.UUID
	FulfilledQuantity pgtype.Int4
	CreatedAt         pgtype.Timestamp
}

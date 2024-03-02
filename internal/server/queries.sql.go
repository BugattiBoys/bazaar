// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package server

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createServer = `-- name: CreateServer :one
INSERT INTO servers (id) VALUES ($1) RETURNING id, ip_address, api_key, created_at, last_ping
`

func (q *Queries) CreateServer(ctx context.Context, id pgtype.UUID) (Server, error) {
	row := q.db.QueryRow(ctx, createServer, id)
	var i Server
	err := row.Scan(
		&i.ID,
		&i.IpAddress,
		&i.ApiKey,
		&i.CreatedAt,
		&i.LastPing,
	)
	return i, err
}

const findServer = `-- name: FindServer :one
SELECT id, ip_address, api_key, created_at, last_ping FROM servers WHERE id = $1 AND api_key = $2 LIMIT 1
`

type FindServerParams struct {
	ID     pgtype.UUID
	ApiKey pgtype.UUID
}

func (q *Queries) FindServer(ctx context.Context, arg FindServerParams) (Server, error) {
	row := q.db.QueryRow(ctx, findServer, arg.ID, arg.ApiKey)
	var i Server
	err := row.Scan(
		&i.ID,
		&i.IpAddress,
		&i.ApiKey,
		&i.CreatedAt,
		&i.LastPing,
	)
	return i, err
}

const getServer = `-- name: GetServer :one
SELECT id, ip_address, api_key, created_at, last_ping FROM servers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetServer(ctx context.Context, id pgtype.UUID) (Server, error) {
	row := q.db.QueryRow(ctx, getServer, id)
	var i Server
	err := row.Scan(
		&i.ID,
		&i.IpAddress,
		&i.ApiKey,
		&i.CreatedAt,
		&i.LastPing,
	)
	return i, err
}

const pingServer = `-- name: PingServer :exec
UPDATE servers SET last_ping = CURRENT_TIMESTAMP WHERE id = $1
`

func (q *Queries) PingServer(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, pingServer, id)
	return err
}

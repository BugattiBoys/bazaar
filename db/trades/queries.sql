-- name: GetServerOpenTradesWithStatus :many
SELECT * FROM trades WHERE initiating_server = $1 AND status = $2;
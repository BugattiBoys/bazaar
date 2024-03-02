-- name: GetServer :one
SELECT * FROM servers
WHERE id = $1 LIMIT 1;

-- name: CreateServer :one
INSERT INTO servers (id) VALUES ($1) RETURNING *;

-- name: PingServer :exec
UPDATE servers SET last_ping = CURRENT_TIMESTAMP WHERE id = $1;

-- name: FindServer :one
SELECT * FROM servers WHERE id = $1 AND api_key = $2 LIMIT 1;
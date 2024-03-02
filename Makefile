migrate:
	migrate -database $(DATABASE_URL) -path db/migrations up
rollback:
	migrate -database $(DATABASE_URL) -path db/migrations down $(STEPS)
migrate-force:
	migrate -database $(DATABASE_URL) -path db/migrations force $(VERSION)
compile:
	buf generate proto
	sqlc generate
dev:
	go run main.go
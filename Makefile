MIGRATE=migrate -path ./migrations -database "postgres://admin:2311@localhost:5432/postgres?sslmode=disable"

.PHONY: migrate migrate-down

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

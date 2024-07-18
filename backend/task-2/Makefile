# Create a set of database up/down migrations with a specific name.
migrate-create:
# -ext adalah file extension, artinya kita membuat file.sql || -dir adalah folder tempat simpan || usahakan jangan menggunakan spasi dalam penamaan file migration
	migrate create -ext sql -dir db/migrations ${APP_DB_NAME}

# Migrate database up
migrate-up:
	@echo "Migrating database up..."
	migrate -path db/migrations -database "postgresql://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable" -verbose up
	@echo "Migration up completed successfully."

# Migrate database down
migrate-down:
	@echo "Migrating database down..."
	migrate -path db/migrations -database "postgresql://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable" -verbose down
	@echo "Migration down completed successfully."

scan-vuln-general:
	osv-scanner -r go.mod

scan-vuln-with-json-output:
	osv-scanner --lockfile=go.mod --config=osv-scanner.toml --format json > results.json

scan-vuln-general-with-config:
		osv-scanner --lockfile=go.mod --config=osv-scanner.toml --format=table

run:
	go run main.go

.PHONY: migrate-create migrate-up migrate-down scan-vuln-dir scan-vuln-specify scan-vuln-std run
include env.sh
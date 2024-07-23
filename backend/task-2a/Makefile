# create migrate database 
migrate-create:
	migrate create -ext sql -dir db/migration $(DB_NAME)

# migrate up database
migrate-up:
	migrate -path db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose up

# migrate down database
migrate-down:
	migrate -path db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose down

# run project
run:
	go run main.go

test:
	go test -v ./test/...

.PHONY: migrate-create migrate-up migrate-down run test
include env.sh
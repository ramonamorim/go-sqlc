include .env

migrate-local-create:
	create -ext=sql -dir=sql/migrations -seq init

migrate-local-up:
	migrate -path=sql/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable&options=-c search_path=$(DB_SCHEMA)" -verbose up

migrate-local-down:
	migrate -path=sql/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable&options=-c search_path=$(DB_SCHEMA)" -verbose down


.PHONY: migrate
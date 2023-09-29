TOPDIR := $(shell pwd)

postgres:
	docker run --name postgres14 --network postr-network -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=postgres --owner=postgres postr

dropdb:
	docker exec -it postgres14 dropdb postr

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/postr?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/postr?sslmode=disable" -verbose down

migrateforce:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/postr?sslmode=disable" force ${VERSION}

sqlcinit:
	docker run --rm -v $(TOPDIR):/src -w /src kjconroy/sqlc init

sqlcgenerate:
	docker run --rm -v $(TOPDIR):/src -w /src kjconroy/sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlcinit sqlcgenerate server
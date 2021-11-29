postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root ecom

dropdb:
	docker exec -it postgres13 dropdb ecom

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ecom?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ecom?sslmode=disable" -verbose down

sqlcwindows:
	docker run --rm -v $(CURDIR):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/ecom-go/main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlcwindows test
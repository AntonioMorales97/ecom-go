postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root ecom

dropdb:
	docker exec -it postgres13 dropdb ecom

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ecom?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ecom?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ecom?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ecom?sslmode=disable" -verbose down 1

sqlcwindows:
	docker run --rm -v $(CURDIR):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/ecom-go/main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/AntonioMorales97/ecom-go/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlcwindows test server mock
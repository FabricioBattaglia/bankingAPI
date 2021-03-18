postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root dbname

dropdb:
	docker exec -it postgres12 dropdb dbname

psql:
	docker exec -it postgres12 psql -U root dbname

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/dbname?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/dbname?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb psql migrateup migratedown sqlc test server
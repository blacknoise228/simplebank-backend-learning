postgresrm:
	docker stop postgres2
	docker rm postgres2
postgres:
	docker run --name postgres2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:latest

createdb: 
	docker exec -it postgres2 createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgres2 dropdb simple_bank

gooseup:
	goose up -dir db/migration

goosedown:
	goose down -dir db/migration

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/blacknoise228/simplebank-backend-learning/db/sqlc Store

.PHONY: postgres postgresrm createdb dropdb gooseup goosedown sqlc test server mock

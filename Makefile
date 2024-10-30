postgresrm:
	docker stop postgres2
	docker rm postgres2
postgres:
	docker run --name postgres2 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17-alpine3.20

createdb: 
	docker exec -it postgres2 createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgres2 dropdb simple_bank

gooseup:
	goose up -dir db/migration

gooseup1:
	goose up-to 00001 -dir db/migration

goosedown:
	goose down -dir db/migration

goosedown1:
	goose down-to 00001 -dir db/migration

dbdock: 
	dbdocs build doc/db.dbml

dbschema: 
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/blacknoise228/simplebank-backend-learning/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl


.PHONY: postgres postgresrm createdb dropdb gooseup gooseup1 goosedown goosedown1 dbdock dbschema sqlc test server mock proto evans

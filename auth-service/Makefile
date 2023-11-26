ifndef AUTH_BINARY_FILE
	AUTH_BINARY_FILE=authApp
endif

ifndef DB_URL
	DB_URL=postgresql://root:secret@localhost:5432/auth?sslmode=disable
endif

SERVER_CMD_FILE=cmd/main.go

protobuf:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

server_run:
	go mod tidy
	go run ${SERVER_CMD_FILE}

server_build:
	CGO_ENABLED=0 go build -o ${AUTH_BINARY_FILE} ${SERVER_CMD_FILE}

create_db:
	docker exec -it postgres createdb --username=root --owner=root auth

drop_db:
	docker exec -it postgres dropdb auth

migrate_up:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

queries:
	sqlc generate
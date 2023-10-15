DB_URL=postgresql://root:secret@localhost:5432/backend?sslmode=disable

reset_docker:
	-docker rm -vf backend
	-docker rmi -f backend
	-docker rm -vf postgres
	-docker rmi -f postgres
	-docker rm -vf migrate

backend_build:
	make network;
	make postgres;
	-docker rm -vf backend;
	-docker rmi -f backend:latest;
	-docker rmi -f docker.io/library/golang:1.21-alpine3.18
	docker build -t backend:latest -f bff/Dockerfile
	docker exec -it postgres createdb --username=root --owner=root backend
	docker run --name migrateup --rm --privileged=true -v $(PWD)/bff/db/migration:/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_URL) up

rebuild:
	-docker stop backend-bff_api_1
	-docker stop backend-bff_postgres_1
	-docker rm backend-bff_api_1
	-docker rmi backend-bff_api
	make backend

backend:
	docker-compose -f ./bff/docker-compose.yaml -p backend-bff up -d

stop-backend:
	docker-compose -f ./bff/docker-compose.yaml -p backend-bff down

dev:
	make network
	make postgres
	make createdb
	make migrateup

network:
	-docker network create backend-network

postgres:
	docker start backend-bff_postgres_1 || docker run --name backend-bff_postgres_1 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret --network backend-network -d postgres:15-alpine

migratenew:
	docker run --name migratenew --privileged=true --rm -v $(PWD)/bff/db/migration:/migrations --network host migrate/migrate -path=/migrations/ create -ext sql -dir migrations -seq $(name)

migrateup:
	docker run --name migrateup --privileged=true --rm -v $(PWD)/bff/db/migration:/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_URL) up

migratedown:
	docker run --name migratedown --privileged=true --rm -v $(PWD)/bff/db/migration:/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_URL) down

createdb:
	docker exec -it postgres createdb --username=root --owner=root backend

dropdb:
	docker exec -it postgres dropdb backend

sqlc:
	cd bff && \
	sqlc generate && \
	cd ..

sqlcinit:
	cd bff && \
	sqlc init

test:
	cd bff && \
	go test -v -cover -short -count=1 ./... \
	&& cd ..

coverage:
	cd bff && \
	go test -coverprofile=coverage.out ./... && \
	go tool cover -html=coverage.out && \
	cd ..

server:
	cd bff && \
	go run main.go && \
	cd ..

mock:
	cd bff && mockgen -package mockdb -destination db/mock/store.go github.com/itsscb/backend-template/bff/db/sqlc Store && cd ..

proto:
	cd bff && \
	rm -f doc/swagger/*.swagger.json && \
	rm -f pb/*.go && \
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=backend \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	./proto/*.proto
	cd ..

evans:
	evans --host localhost --port 9090 --package pb -r repl

count_lines:
	cloc --exclude-dir=.history,.git .

.PHONY: reset_docker backend_build rebuild backend backend dev network postgres migratenew migrateup migratedown createdb dropdb sqlc sqlcinit test coverage server mock proto evans count_lines
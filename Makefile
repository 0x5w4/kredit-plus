.PHONY:

run_api_gateway:
	go run api-gateway-service/cmd/main.go -config=./api-gateway-service/config/config.yaml

run_writer:
	go run writer-service/cmd/main.go -config=./writer-service/config/config.yaml

run_reader:
	go run reader-service/cmd/main.go -config=./reader-service/config/config.yaml

# ==============================================================================
# Docker

docker_dev:
	@echo Starting local docker dev compose
	docker-compose -f docker-compose.yaml up --build

local:
	@echo Starting local docker compose
	docker-compose -f docker-compose.local.yaml up -d --build


# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)


# ==============================================================================
# Modules support

tidy:
	go mod tidy

deps-reset:
	git checkout -- go.mod
	go mod tidy

deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache


# ==============================================================================
# Linters https://golangci-lint.run/usage/install/

run-linter:
	@echo Starting linters
	golangci-lint run ./...

# ==============================================================================
# PPROF

pprof_heap:
	go tool pprof -http :8006 http://localhost:6060/debug/pprof/heap?seconds=10

pprof_cpu:
	go tool pprof -http :8006 http://localhost:6060/debug/pprof/profile?seconds=10

pprof_allocs:
	go tool pprof -http :8006 http://localhost:6060/debug/pprof/allocs?seconds=10



# ==============================================================================
# Go migrate postgresql https://github.com/golang-migrate/migrate

DB_NAME = products
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable

force_db:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations force 1

version_db:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations down 1


# ==============================================================================
# MongoDB

mongo:
	cd ./scripts && mongo admin -u admin -p admin < init.js


# ==============================================================================
# Swagger

swagger:
	@echo Starting swagger generating
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g **/**/*.go

# ==============================================================================
# Proto

proto_kafka:
	@echo Generating kafka proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	cd proto/kafka && \
	protoc \
    --plugin=$(GOPATH)/bin/protoc-gen-go \
    --plugin=$(GOPATH)/bin/protoc-gen-go-grpc \
	--go_out=. \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=. \
	kafka.proto

proto_writer:
	@echo Generating product writer microservice proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	cd writer-service/proto/writer && \
	protoc \
    --plugin=$(GOPATH)/bin/protoc-gen-go \
    --plugin=$(GOPATH)/bin/protoc-gen-go-grpc \
	--go_out=. \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=. \
	writer.proto

proto_writer_message:
	@echo Generating product writer messages microservice proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	cd writer-service/proto/writer && \
	protoc \
	--plugin=$(GOPATH)/bin/protoc-gen-go \
    --plugin=$(GOPATH)/bin/protoc-gen-go-grpc \
	--go_out=. \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=. \
	writer_messages.proto

proto_reader:
	@echo Generating product reader microservice proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	cd reader-service/proto/reader && \
	protoc \
    --plugin=$(GOPATH)/bin/protoc-gen-go \
    --plugin=$(GOPATH)/bin/protoc-gen-go-grpc \
	--go_out=. \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=. \
	reader.proto

proto_reader_message:
	@echo Generating product reader messages microservice proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	cd reader-service/proto/reader && \
	protoc \
    --plugin=$(GOPATH)/bin/protoc-gen-go \
    --plugin=$(GOPATH)/bin/protoc-gen-go-grpc \
	--go_out=. \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=. \
	reader_messages.proto
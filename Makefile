##################################################
# variables
##################################################
BIN=./bin

# ADOPTIONS
ADOPTIONS_SCHEMA=db/adoptions/schema.sql
ADOPTIONS_MIGRATIONS=db/adoptions/migrations

##################################################
# rules
##################################################

build:
	make build-linux
	make build-windows

build-linux:
	go env -w GOOS=linux && go build -o $(BIN)/ ./...

build-windows:
	go env -w GOOS=windows && go build -o $(BIN)/ ./...

build-adoptions:
	go env -w GOOS=linux && go build -o $(BIN)/ ./cmd/adoptions/...

build-bot:
	go env -w GOOS=linux && go build -o $(BIN)/ ./cmd/bot/...

build-events:
	go env -w GOOS=linux && go build -o $(BIN)/ ./cmd/events/...

build-news:
	go env -w GOOS=linux && go build -o $(BIN)/ ./cmd/news/...

build-shop:
	go env -w GOOS=linux && go build -o $(BIN)/ ./cmd/shop/...

clean:
	make clean_windows
	make clean_linux

clean-linux:
	rm ./$(BIN)/$(APP)

clean-windows:
	rm ./$(BIN)/$(APP).exe

run:
	go run ./...

test:
	go test ./... -v

adoptions-db-status:
	dbmate \
	--env ADOPTIONS_DB_URL \
	--schema-file $(ADOPTIONS_SCHEMA) \
	--migrations-dir $(ADOPTIONS_MIGRATIONS) \
	status 

adoptions-db-up:
	dbmate \
	--env ADOPTIONS_DB_URL \
	--schema-file $(ADOPTIONS_SCHEMA) \
	--migrations-dir $(ADOPTIONS_MIGRATIONS) \
	up

adoptions-db-down:
	dbmate \
	--env ADOPTIONS_DB_URL \
	--schema-file $(ADOPTIONS_SCHEMA) \
	--migrations-dir $(ADOPTIONS_MIGRATIONS) \
	down 

adoptions-db-sqlc:
	sqlc generate

gen-proto:
	cd data && protoc --go_out=.. --go-grpc_out=.. ./*.proto

.PHONY:  build build-linux build-windows \
 	 	 clean clean-linux clean-windows \
 	 	 gen-sql gen-proto \
 	 	 run test \

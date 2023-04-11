##################################################
# variables
##################################################
BIN=./bin


##################################################
# rules
##################################################
build:
	make build-linux
	make build-windows

build-linux:
	go env -w GOOS=linux && go build -o $(BIN)/ ./...
	mv $(BIN)/api $(BIN)/petpark

build-windows:
	go env -w GOOS=windows && go build -o $(BIN)/ ./...
	mv $(BIN)/api.exe $(BIN)/petpark.exe

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

gen-sql:
	sqlc generate

gen-proto:
	cd data && protoc --go_out=.. --go-grpc_out=.. ./*.proto

.PHONY:  build build-linux build-windows \
 	 	 clean clean-linux clean-windows \
 	 	 gen-sql gen-proto \
 	 	 run test \

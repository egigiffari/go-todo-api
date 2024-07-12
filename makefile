BIN=todoapp

build:
	@go build -o ./bin/${BIN} ./cmd/main.go

run: build
	@./bin${BIN} start

clean:
	@rm -rf ./bin${BIN}
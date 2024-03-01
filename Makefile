dev:
	@go build -tags dev -o bin/todo_list_dev cmd/main.go
	@./bin/todo_list_dev
build:
	@go build -o bin/todo_list cmd/main.go
clean:
	@rm -rf ./bin
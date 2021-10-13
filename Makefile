run:
	go run main.go

migrate:
	go run main.go -migrate

test:
	go clean -cache
	go test -v -cover github.com/mrrizal/fiber-example/book

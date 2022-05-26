build:
	go build -o cmd/app cmd/main.go

start:
	go run cmd/main.go route=$(route)

publish:
	go test ./tests/... -v -count=1
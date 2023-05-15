migrate:
	go run ./pkg/migrate/migration.go

run:
	go run main.go

test:
	go test -v ./test/...
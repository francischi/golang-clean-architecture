FROM golang
RUN     mkdir -p /app
WORKDIR /app
COPY . .
RUN    go mod download
RUN    go build -o app
RUN    go run ./pkg/migrate/migrate.go
EXPOSE 8000
ENTRYPOINT  ["./app"]

#> docker build -t 'golang_docker' . 
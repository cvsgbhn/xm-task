
export PORT = 8080
export DB_PORT = 15432
export DB_HOST=localhost
export DB_PORT=15432
export DB_NAME=xm
export DB_USER=dev
export DB_PWD=dev

build-db:
	docker-compose -f ./dockers/docker-compose.yml up -d

run:
	go run ./cmd/main.go

### Migrations ###

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1

migrate-new:
	migrate create -ext sql -dir ./migrations "$(name)"

migrate-up:
	migrate -database="postgres://${DB_USER}:${DB_PWD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations up

# rollback last migration
migrate-down:
	migrate -database="postgres://${DB_USER}:${DB_PWD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations down 1

# curl localhost:8080/add -XPOST -d '{"name": "asdf", "country": "Argentina", "phone": "1231223", "website": "asdafafad"}' | jq
# curl localhost:8080/b
# curl localhost:8080/b -XPUT -d '{"name": "a", "country": "Argentina", "phone": "1", "website": "a"}'
# curl curl localhost:8080/b -XDELETE
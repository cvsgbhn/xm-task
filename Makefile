build-local:
	go build -o ./build/xm ./cmd/

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

# export DB_USER=dev DB_PWD=dev DB_HOST=localhost DB_PORT=15432 DB_NAME=xm
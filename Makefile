### Migrations ###

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1

migrate-new:
	migrate create -ext sql -dir ./migrations "$(name)"
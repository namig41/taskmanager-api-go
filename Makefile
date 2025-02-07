APP_BIN = app/build/app

MIGRATE_BIN = $(shell which migrate)
DB_URL = "${DATABASE_PROVIDER}://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_DB}?sslmode=disable"


build: clean $(APP_BIN)

$(APP_BIN):
	mkdir -p app/build
	cd app && go build -o build/app ./cmd/app/main.go

run: clean $(APP_BIN)
	./$(APP_BIN)

lint:
	golangci-lint run

clean:
	rm -rf ./app/build || true

swagger:
	swag init -g ./app/cmd/app/main.go -o ./app/docs

migrate.up:
	$(MIGRATE_BIN) -path ./migrations -database $(DB_URL) up

migrate.down:
	$(MIGRATE_BIN) -path ./migrations -database $(DB_URL) down 1

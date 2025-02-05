APP_BIN = app/build/app

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

migrate:
	$(APP_BIN) migrate -version $(version)

migrate.down:
	$(APP_BIN) migrate -seq down

migrate.up:
	$(APP_BIN) migrate -seq up

.PHONY: build
build:
	@echo "Building..."
	go build -o bin/main ./cmd/transactions
	@echo "Done."

.PHONY: run
run: build
	@echo "Running..."
	./bin/main
	@echo "Done."


.PHONY: run/db
run/db:
	@echo "Starting database..."
	docker-compose up -d db
	@echo "Done."

.PHONY: migrate/db
migrate/db: run/db
	@echo "Migrating database..."
	atlas migrate up --env dev
	@echo "Done."

build/docker:
	@echo "Building docker image..."
	docker build -t fairnsquare/transactions --target=prod .
	@echo "Done."

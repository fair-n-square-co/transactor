.PHONY: build
build:
	@echo "Building..."
	go build -o bin/main ./cmd/transactions
	@echo "Done."

.PHONY: run
run: run/db build
	@echo "Running..."
	./bin/main
	@echo "Done."

.PHONY: test
test:
	@echo "Running tests..."
	go test -race -covermode=atomic -coverprofile=cover.out -v ./...
	@echo "Done."

.PHONY: test/coverage
test/coverage: test
	@echo "Generating coverage report..."
	go tool cover -html=cover.out
	@echo "Done."

.PHONY: run/db
run/db:
	@echo "Starting database..."
	docker-compose up -d --no-recreate db
	@echo "Done."

.PHONY: migrate/gen
migrate/gen:
	@echo "\nReplace <migration_name> in the command below \nwith an appropriate migration name \nand generate migration files\n"
	@echo "=============================================="
	@echo "atlas migrate diff <migration_name> --env dev"
	@echo "=============================================="

.PHONY: migrate/db
migrate/db:
	@echo "Migrating database..."
	atlas migrate apply --env dev
	@echo "Done."


.PHONY: migrate/prod
migrate/fly:
	@echo "Migrating database in fly.io..."
	@echo "Ensure you have the fly cli installed and logged in"
	@echo "Ensure you update atlas.hcl with the correct DB password for env \"fly\""
	@echo "Ensure you have this command running - \`fly proxy 5454:5432 -a transactions-pg\`"
	# atlas migrate apply --env fly
	@echo "Done."

.PHONY: migrate/rehash
migrate/rehash: run/db
	@echo "Rehashing database..."
	atlas migrate hash --env dev
	@echo "Done."

.PHONY: build/docker
build/docker:
	@echo "Building docker image..."
	docker build -t fairnsquare/transactions --target=prod .
	@echo "Done."

.PHONY: gen/mock
gen/mock:
	@echo "Generating mocks..."
	go generate ./...
	@echo "Done."

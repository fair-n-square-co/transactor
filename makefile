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

build/docker:
	@echo "Building docker image..."
	docker build -t fairnsquare/transactions --target=prod .
	@echo "Done."

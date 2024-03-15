.PHONY: build
build:
	@echo "Building..."
	go build -o bin/main ./cmd/transactor
	@echo "Done."

.PHONY: run
run: build
	@echo "Running..."
	./bin/main
	@echo "Done."

build/docker:
	@echo "Building docker image..."
	docker build -t fairnsquare/transactor --target=prod .
	@echo "Done."

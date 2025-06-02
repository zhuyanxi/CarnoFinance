.PHONY: dev build

dev:
	@echo "Running in development mode..."
	# Insert development mode commands here
	# e.g., start watchers, servers, or run tests
	@cd ui && npm run build && cd ../cmd && go run main.go 
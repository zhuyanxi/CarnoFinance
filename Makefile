.PHONY: dev build import-ashare-yahoo upload-r2

dev:
	@echo "Running in development mode..."
	# Insert development mode commands here
	# e.g., start watchers, servers, or run tests
	@cd ui && npm run build && cd ../cmd && go run main.go

import-ashare-yahoo:
	@go run ./cmd/yahooimport

upload-r2:
	@go run ./cmd/r2upload
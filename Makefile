clean:
	@echo "Cleaning caches..."
	go clean -cache
	go clean -modcache
	go clean -testcache
	go clean -fuzzcache

test:
	go test ./... -count=1

.PHONY: clean test

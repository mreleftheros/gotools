default: clean

clean:
    go clean -cache
    go clean -modcache
    go clean -testcache
    go clean -fuzzcache

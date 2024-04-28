default: test-unit

unit-test:
	go test -v -race -count=1 ./pkg/...

tool-revive:
	go install github.com/mgechev/revive@master

lint: tool-revive
	golangci-lint run
	revive -config ./revive.toml
	go mod tidy -v && git --no-pager diff --quiet go.mod go.sum


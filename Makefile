dependency:
	@go get -v ./...

unit-test: dependency
	@go test -v -short ./...

cover :
	@echo "\x1b[32;1m>>> running unit test and calculate coverage \x1b[0m"
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo "mode: atomic" > coverage.txt

	@go test ./...  -cover -coverprofile=coverage.txt -covermode=count \
		-coverpkg=$$(go list ./...  | grep -v mocks | tr '\n' ',')
	@go tool cover -func=coverage.txt
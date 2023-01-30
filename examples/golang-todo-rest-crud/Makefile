build:
	@go build -o bin/todo ./cmd/todo

run: build
	@./bin/todo

unittest:
	@go test -v ./... -cover

integrationtest:
	@go clean -testcache
	@go test -v ./test/integration --tags=integration

check_swagger_install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger
swagger: check_swagger_install
	swagger generate spec -o ./docs/swagger.yaml --scan-models
	swagger generate spec -o ./third_party/swagger-ui-4.11.1/swagger.json --scan-models

check_mockery_install:
	which mockery || go install github.com/vektra/mockery/v2@latest
mocks: check_mockery_install
	go generate ./...

dcbuild:
	docker compose up --build -d

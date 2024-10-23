TARGET_DIR = cmd/app
DOCS_DIR = docs
ROUTERS_FILE = internal/controller/http/v1/router.go

all: run

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/router.go

run: swag-v1 ### download deps; create swag; run app
	go mod tidy && go mod download && \
	DISABLE_SWAGGER_HTTP_HANDLER='' \
	GIN_MODE=debug \
	CGO_ENABLED=0 \
	go run $(TARGET_DIR)

test:
	go test -v ./...

postgres-up:
	docker compose up -d

postgres-down:
	docker compose down

clean:
	rm -rf $(DOCS_DIR)

.PHONY: all swag-v1 run test clean postgres-up postgres-down


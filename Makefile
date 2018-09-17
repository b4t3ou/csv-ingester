deps:
		go get -u github.com/alecthomas/gometalinter
		gometalinter --install
		glide install

lint:
		gometalinter --config=gometalinter_config.json ./...

test: lint
		go test -cover -short -timeout=10s $$(glide novendor)

test-all: lint
		go test -tags integration -cover -short -failfast -timeout=60s ./...

generate:
		protoc -I proto/ proto/storage.proto --go_out=plugins=grpc:proto

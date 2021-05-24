# Go parameters
GOCMD?=go
GOBUILD?=$(GOCMD) build
GOCLEAN?=$(GOCMD) clean
GOTEST?=$(GOCMD) test
GOGET?=$(GOCMD) get
GOBIN?=$(GOPATH)/bin

# Build parameters
OUT?=./out
DOCKER_TMP?=$(OUT)/docker_temp/
DOCKER_TAG?=docker-reg.devops.xiaohongshu.com/shequ/consul2istio:latest
BINARY_NAME?=$(OUT)/consul2istio
BINARY_NAME_DARWIN?=$(BINARY_NAME)-darwin
MAIN_PATH_CONSUL_MCP=./cmd/consul2istio/main.go

build:
	CGO_ENABLED=0 GOOS=linux  $(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH_CONSUL_MCP)
build-mac:
	CGO_ENABLED=0 GOOS=darwin  $(GOBUILD) -o $(BINARY_NAME_DARWIN) $(MAIN_PATH_CONSUL_MCP)

docker-build: build
	mkdir $(DOCKER_TMP)
	cp ./docker/Dockerfile $(DOCKER_TMP)
	cp $(BINARY_NAME) $(DOCKER_TMP)
	docker build -t $(DOCKER_TAG) $(DOCKER_TMP)
	rm -rf $(DOCKER_TMP)
docker-push:
	docker push $(DOCKER_TAG)

style-check:
	gofmt -l -d ./
	goimports -l -d ./
lint:
	golangci-lint run -v
test:
	go test --race ./...
clean:
	rm -rf $(OUT)

consul:
	helm install -n consul -f k8s/config.yaml consul hashicorp/consul

.PHONY: build docker-build docker-push clean

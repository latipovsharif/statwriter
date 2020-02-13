REGISTRY = reg.techhprof.ru
PROJECT = rtb
APPLICATION = dailystatuploader
TAG_PROD = test
TAG_DEV = dev
BUILD_DIRECTORY = bin

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	# For some reason this install method is not recommended, see https://github.com/golangci/golangci-lint#install
	# However, they suggest `curl ... | bash` which we should not do
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -local gitlab.com/devtizerteam/real-time-bidding/gate2pushprofit ./

lint: ## Run linters. Use make install-linters first.
	#vendorcheck ./...
	PATH=$$PATH:$$HOME/go/bin golangci-lint run -c .golangci.yml ./...
	@# The govet version in golangci-lint is out of date and has spurious warnings, run it separately
	go vet -all ./...

test: ## Run tests
	@mkdir -p coverage/
	go test -coverpkg="gitlab.com/devtizerteam/real-time-bidding/gate2pushprofit/..." -coverprofile=coverage/go-test.coverage.out -timeout=5m ./internal/...

dep: ## Update vendor
	GOPRIVATE=gitlab.com/devtizerteam/* go mod tidy
	GOPRIVATE=gitlab.com/devtizerteam/* go mod vendor -v

clear: ## clears all binaries
	rm -rf ./bin/

build: clear ## build and prepare project without config file
	go build -o ./bin/dailystatuploader
	cp ./Dockerfile ./bin/
	cp ./date_stat.xlsx ./bin/
	cp ./country_stat.xlsx ./bin/
	cp ./device_pc_stat.xlsx ./bin/
	cp ./device_mobile_stat.xlsx ./bin/
	cp ./subscription_stat.xlsx ./bin/

build-prod: build ## build and copy prod config file
	cp ./config.yaml ./bin/config.yaml

build-dev: build ## build and copy dev config file
	cp ./config.yaml ./bin/config.yaml

push-prod: build-prod # build and push prod
	docker build -t ${REGISTRY}/${PROJECT}/${APPLICATION}:${TAG_PROD} ./${BUILD_DIRECTORY}
	docker push ${REGISTRY}/${PROJECT}/${APPLICATION}:${TAG_PROD}

push-dev: build-dev ## build and push dev
	docker build -t ${REGISTRY}/${PROJECT}/${APPLICATION}:${TAG_DEV} ./${BUILD_DIRECTORY}
	docker push ${REGISTRY}/${PROJECT}/${APPLICATION}:${TAG_DEV}

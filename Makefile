PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
UIPATH := $(PWD)/browser/flagr-ui

################################
### Public
################################

all: deps gen build build_ui run

rebuild: gen build

test: verifiers
	@go test -race -covermode=atomic -coverprofile=coverage.txt ./pkg/...

ci: test

build:
	@echo "Building Flagr Server to $(PWD)/flagr ..."
	@CGO_ENABLED=1 go build -o $(PWD)/flagr ./swagger_gen/cmd/flagr-server

build_ui:
	@echo "Building Flagr UI ..."
	@cd ./browser/flagr-ui/; yarn install && yarn run build

run:
	@$(PWD)/flagr --port 18000

gen: api_docs swagger

deps: checks
	@echo "Installing retool" && GO111MODULE=off go get -u github.com/twitchtv/retool
	@retool sync
	@retool build
	@retool do gometalinter --install
	@echo "Sqlite3" && sqlite3 -version

watch:
	@retool do fswatch

serve_docs:
	@yarn global add docsify-cli@4
	@docsify serve $(PWD)/docs

################################
### Private
################################

api_docs:
	@echo "Installing swagger-merger" && yarn global add swagger-merger
	@swagger-merger -i $(PWD)/swagger/index.yaml -o $(PWD)/docs/api_docs/bundle.yaml

checks:
	@echo "Check deps"
	@(env bash $(PWD)/buildscripts/checkdeps.sh)
	# @echo "Checking project is in GOPATH"
	# @(env bash $(PWD)/buildscripts/checkgopath.sh)

verifiers: verify_gometalinter verify_swagger

verify_gometalinter:
	@echo "Running $@"
	@retool do gometalinter --config=.gometalinter.json ./pkg/...

verify_swagger:
	@echo "Running $@"
	@retool do swagger validate $(PWD)/docs/api_docs/bundle.yaml

clean:
	@echo "Cleaning up all the generated files"
	@find . -name '*.test' | xargs rm -fv
	@rm -rf build
	@rm -rf release

swagger: verify_swagger
	@echo "Regenerate swagger files"
	@rm -f /tmp/configure_flagr.go
	@cp $(PWD)/swagger_gen/restapi/configure_flagr.go /tmp/configure_flagr.go 2>/dev/null || :
	@rm -rf $(PWD)/swagger_gen
	@mkdir $(PWD)/swagger_gen
	@retool do swagger generate server -t ./swagger_gen -f $(PWD)/docs/api_docs/bundle.yaml
	@cp /tmp/configure_flagr.go $(PWD)/swagger_gen/restapi/configure_flagr.go 2>/dev/null || :

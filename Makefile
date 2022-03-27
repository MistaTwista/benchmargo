# Build variables
GO_BIN_PATH ?= go
TAGS ?= generics

#BINARY_NAME = jironimo
#BUILD_DIR = build
#VERSION ?= $(shell git tag --points-at HEAD | tail -n 1)
#BUILD_DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%S")
#COMMIT_SHA = $(shell git rev-parse --short HEAD)
#LDFLAGS = -ldflags "-w -X main.version=${VERSION} -X main.buildDate=${BUILD_DATE} -X main.commit=${COMMIT_SHA}"
#USERNAME = $(shell git config user.name)
#GOOS ?= linux

.PHONY: test
test: ## Run unit tests
	${GO_BIN_PATH} test -v ./... --tags=${TAGS}

.DEFAULT_GOAL := help
help: ## Show this message
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)

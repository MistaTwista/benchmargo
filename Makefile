# Build variables
GO_BIN_PATH ?= go
TAGS ?= generics
BENCH ?= ^BenchmarkJust*
CPU ?= 1,12
COUNT ?= 2
path ?= ./...

.PHONY: version
version: ## Shows go version for GO_BIN_PATH
	${GO_BIN_PATH} version

.PHONY: bench
bench: version ## Run benchmarks
	${GO_BIN_PATH} test -bench=. -count=${COUNT} -cpu=${CPU} --tags=${TAGS} -benchmem -v ./...

.PHONY: benchone
benchone: version ## Run benchmark in one package. Example: BENCH=^BenchmarkAwesomeToJSON.* make benchone path=./internal/genechacha
	${GO_BIN_PATH} test -bench=${BENCH} -count=${COUNT} -cpu=${CPU} -benchmem -cpuprofile=cpu.out -memprofile=mem.out -v --tags=${TAGS} ${path}

# use top10
# list <function name>
.PHONY: pprof
pprof: version ## Profile
	${GO_BIN_PATH} tool pprof genechacha.test cpu.out

.DEFAULT_GOAL := help
help: ## Show this message
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)

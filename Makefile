# Build variables
GO_BIN_PATH ?= go
TAGS ?= generics
BENCH ?= ^BenchmarkJust*
CPU ?= 1,4,12
COUNT ?= 5
path ?= ./...

version: ## Shows go version for GO_BIN_PATH
	${GO_BIN_PATH} version

run: ## Run all benchmarks and generate comparison files
	./lister.sh

bench: version ## Run benchmarks
	${GO_BIN_PATH} test -bench=. -count=${COUNT} -cpu=${CPU} --tags=${TAGS} -benchmem -v ./...

gbench: version ## Run generics benchmark
	${GO_BIN_PATH} test -bench=G$$ -run=^$$ -count=${COUNT} -cpu=${CPU} -benchmem -v --tags=${TAGS} ${path}

jbench: version ## Run native benchmark
	${GO_BIN_PATH} test -bench=J$$ -run=^$$ -count=${COUNT} -cpu=${CPU} -benchmem -v ${path}

benchone: version ## Run benchmark in one package with pprof out. Example: BENCH=^BenchmarkAwesomeToJSON.* make benchone path=./internal/genechacha
	${GO_BIN_PATH} test -bench=${BENCH} -count=${COUNT} -cpu=${CPU} -benchmem -cpuprofile=cpu.out -memprofile=mem.out -v --tags=${TAGS} ${path}

# use top10
# list <function name>
pprof: version ## Profile
	${GO_BIN_PATH} tool pprof genechacha.test cpu.out

debug: ## Run delve
	dlv exec build/tricky

.DEFAULT_GOAL := help
help: ## Show this message
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)

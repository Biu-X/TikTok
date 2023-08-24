GO ?= go
AIR ?= air

.DEFAULT_GOAL := default

TAGS ?=

default: help

.PHONY: tidy
tidy: ## go mod tidy
	${GO} mod tidy

.PHONY: build
build: ## build tiktok binary file
	${GO} build -o tiktok .

.PNONY: gen
gen: build ## gen gorm code
	./tiktok gen

.PHONY: watch
watch: ## live reload
	${AIR} server

.PHONY: test
test: tidy ## go test
	${GO} test -v $$(${GO} list ./... | grep -v /dal/query)

.PHONY: clear
clear: ## clear project
	-rm -rf ./tmp tiktok


.PHONY: help
help: Makefile ## print Makefile help information.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' Makefile #$(MAKEFILE_LIST)

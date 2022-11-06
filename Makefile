# Based on that awesome makefile https://github.com/dunglas/symfony-docker/blob/main/docs/makefile.md#the-template

.DEFAULT_GOAL = default

.PHONY: default
default: build

##—— 📚 Help ——————————————————————————————————————————————————————————————
.PHONY: help
help: ## ❓ Dislay this help
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' \
		| sed -e 's/\[32m##——————————/[33m           /'  \
		| sed -e 's/\[32m##——/[33m ——/' \
		| sed -e 's/\[32m####/[34m                                 /' \
		| sed -e 's/\[32m###/[36m                                 /' \
		| sed -e 's/\[32m##\?/[35m /'  \
		| sed -e 's/\[32m##/[33m/'

##—— ️⚙️  Environments ——————————————————————————————————————————————————————
.PHONY: configure-test-env
configure-test-env: ## 🤖 Install required libraries for test environment (golint, staticcheck, etc)
configure-test-env:
	go install golang.org/x/lint/golint@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest


##—— 🐹 Golang —————————————————————————————————————————————————
.PHONY: build
build: ## 🗜️  Build package
#### Use build_o="..." to specify build options
$(eval build_o ?=)
build:
	go build -v $(build_o)

.PHONY: verify
verify: ## 🗜️  Verify dependencies
verify:
	go mod verify


##—— 🧪️ Tests —————————————————————————————————————————————————————————————
.PHONY: test
test: ## 🏃 Launch all tests
test: test-vet test-lint test-staticcheck test-go

test-go: ## 🏃 Launch go test
#### Use gotest_o="..." to specify options
$(eval gotest_o ?=)
test-go:
	go test -v  $(gotest_o) ./...

test-vet: ## 🏃 Launch go vet
#### Use vet_o="..." to specify options
$(eval vet_o ?=)
test-vet:
	go vet $(vet_o) ./...

test-lint: ## 🏃 Launch go lint
#### Use lint_o="..." to specify options (-set_exit_status for instance)
$(eval lint_o ?=)
test-lint:
	golint $(lint_o) ./...

test-staticcheck: ## 🏃 Launch staticcheck
#### Use staticcheck_o="..." to specify options
$(eval staticcheck_o ?=)
test-staticcheck:
	staticcheck $(staticcheck_o) ./...

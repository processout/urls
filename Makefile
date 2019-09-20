GOLANGCI_LINT_DEP=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1
COVERFILE=cover.out

.PHONY: coverage
coverage:
	go test -v -race ./... -coverprofile=$(COVERFILE) -covermode=atomic

.PHONY: coverage-html
coverage-html: coverage
coverage-html: ## Generate test coverage report in html
	go tool cover -html=$(COVERFILE)

.PHONY: deps
deps: ## Install dependencies
	go get
	go get $(GOLANGCI_LINT_DEP)

.PHONY: style
style: deps
	golangci-lint run

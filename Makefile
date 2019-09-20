GOLANGCI_LINT_DEP=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1

.PHONY: coverage
coverage:
	go test -v -race ./... -coverprofile=cover.out -covermode=atomic

.PHONY: deps
deps: ## Install dependencies
	go get
	go get $(GOLANGCI_LINT_DEP)

.PHONY: style
style: deps
	golangci-lint -run

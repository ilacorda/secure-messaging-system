LINT_VERSION ?= 1.42.1

.PHONY: lint-install
lint-install:
	@echo "Installing golangci-lint v$(LINT_VERSION)..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v$(LINT_VERSION)

.PHONY: lint
lint:
	@golangci-lint run ./...
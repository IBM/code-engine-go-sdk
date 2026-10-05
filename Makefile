# Makefile to build the project
GO=go
LINT=golangci-lint

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: test lint

test:
	${GO} test `${GO} list ./...`

test-cov:
	${GO} test `${GO} list ./...` ${COVERAGE}

test-int:
	${GO} test `${GO} list ./...` -tags=integration

test-int-cov:
	${GO} test `${GO} list ./...` -tags=integration ${COVERAGE}

lint:
	${LINT} run

# Makefile to build the project
GO=go
LINT=golangci-lint
GOSEC=gosec

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: tidy test lint scan-gosec
travis-ci: test-cov lint scan-gosec tidy

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

scan-gosec:
	${GOSEC} -conf gosec.json -exclude-dir=example/v1 -exclude-dir=example/v2 ./...

tidy:
	${GO} mod tidy
# Makefile to build the project
GO=go
LINT=golangci-lint
GOSEC=gosec

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: tidy test lint scan-gosec
travis-ci: ssh-config test-cov lint scan-gosec tidy
github-action-ci: test-cov lint scan-gosec tidy

ssh-config:
	git config --global url.ssh://git@github.ibm.com/.insteadOf https://github.ibm.com/
	git config --global user.email "coligo.devops@de.ibm.com"
	git config --global user.name "CodeEngine DevOps"

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
	${GOSEC} -conf gosec.json -exclude-dir=example/v1 -exclude-dir=example/v2 -exclude-dir=codeenginev2 ./...

tidy:
	${GO} mod tidy

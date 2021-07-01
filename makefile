.PHONY: dev
SRC_PATH=$(GOPATH)/src/github.com/phuwn/coinchameleon
ORACLE_CONTAINER?=oraleXE

dev:
	@GO111MODULE=on LD_LIBRARY_PATH=~/instantclient_19_8 RUN_MODE=local go run cmd/*.go
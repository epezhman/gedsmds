-include env
include env

UNAME := $(shell uname)
GO_PATH := $(shell which go)

ifeq ($(UNAME), Linux)
MDS_BUILD_PATH = ${MDS_BUILD_PATH_LINUX}
SOURCE = .
endif

ifeq ($(UNAME), Darwin)
MDS_BUILD_PATH = ${MDS_BUILD_PATH_DARWIN}
SOURCE = source
endif

ifeq ($(GO_PATH),)
	GO_PATH = ${GO_REMOTE_PATH}
endif

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: tidy
## tidy: tidy up go modules
tidy:
	@${GO_PATH} mod tidy

.PHONY: clean
## clean: clean build path for Darwin
clean:
	@echo "Cleaning"
	@${GO_PATH} clean
	@rm -rf ${MDS_BUILD_PATH}

.PHONY: clean-all
## clean-all: clean everything
clean-all:
	@echo "Cleaning"
	@${GO_PATH} clean
	@rm -rf ${MDS_BUILD_PATH_LINUX}
	@rm -rf ${MDS_BUILD_PATH_DARWIN}

.PHONY: build
## build: build the MDS
build: clean build-mds

.PHONY: build-mds
## build-mds: build the MDS component (OS-dependent)
build-mds:
	@echo "Building MDS ..."
	@${GO_PATH} build -o ${MDS_BUILD_PATH}${MDS_BINARY} ./cmd/mds

.PHONY: build-remote-linux
## build-remote-linux: build for remote Linux on Darwin
build-remote-linux:
	@echo "Building the components on the build remote system..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="remote" ./scripts/build_on_linux.sh

.PHONY: build-local-linux
## build-local-linux: build for local Linux on Darwin
build-local-linux:
	@echo "Building the components on the build local system..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="local" ./scripts/build_on_linux.sh

.PHONY: run-mds
## run-mds: run the MDS component
run-mds:
	@echo "Running MDS ..."
	@${GO_PATH} run ./cmd/mds

.PHONY: run-playground
## run-playground: run some experimental codes in ./cmd/playground
run-playground:
	@echo "Running Playground ..."
	@${GO_PATH} run -race ./cmd/playground
	#@${GO_PATH} run  ./cmd/playground

## git-commit: commit all files
git-commit:
	@echo "Commit"
	@git add . ; git commit -m 'auto push';

## git-push: push to main
git-push:
	@echo "Pushing to git main"
	@git add . ; git commit -m 'auto push'; git push origin main

## git-pull: pull from main
git-pull:
	@echo "Pulling from git main"
	@git pull origin main

.PHONY: protos
## protos: generate the protos
protos:
	@echo "Generating the Protos ..."
	@rm -rf ./protos/goprotos;
	@mkdir ./protos/goprotos
	@protoc -I ./protos ./protos/*.proto  --go_out=./protos/goprotos
	@protoc -I ./protos ./protos/*.proto  --go-grpc_out=require_unimplemented_servers=false:./protos/goprotos

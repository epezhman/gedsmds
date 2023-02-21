-include env
include env

UNAME := $(shell uname)
GO_PATH := $(shell which go)

ifeq ($(UNAME), Linux)
MDS_BUILD_PATH = ${MDS_BUILD_PATH_LINUX}
GO_PATH = ${GO_PATH_LINUX}
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

## run-mds: run the MDS component
run-mds:
	@echo "Running MDS ..."
	@${GO_PATH} run ./cmd/mds

## run-mock-client: run the MockClient
run-mock-client:
	@echo "Running Mock-Client ..."
	@${GO_PATH} run ./cmd/mockclient

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
	@rm -rf ./protos/protos;
	@mkdir ./protos/protos
	@protoc -I ./protos ./protos/*.proto  --go_out=./protos/protos
	@protoc -I ./protos ./protos/*.proto  --go-grpc_out=require_unimplemented_servers=false:./protos/protos

.PHONY: create-certificates
## create-certificates: create certificates
create-certificates:
	@echo "Creating the certificates and keys ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} ./scripts/create_certificates.sh

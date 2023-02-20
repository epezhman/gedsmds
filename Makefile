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
run-node:
	@echo "Running MDS ..."
	@${GO_PATH} run ./cmd/mds

.PHONY: run-playground
## run-playground: run some experimental codes in ./cmd/playground
run-playground:
	@echo "Running Playground ..."
#	@${GO_PATH} run -race ./cmd/playground
	@${GO_PATH} run  ./cmd/playground

.PHONY: git-commit
## git-commit: commit all files
git-commit:
	@echo "Commit"
	@git add . ; git commit -m 'auto push';

.PHONY: git-push
## git-push: push to master
git-push:
	@echo "Pushing to git master"
	@git add . ; git commit -m 'auto push'; git push origin master

.PHONY: git-pull
## git-pull: pull from master
git-pull:
	@echo "Pulling from git master"
	@git pull origin master

.PHONY: protos
## protos: generate the protos
protos:
	@echo "Generating the Protos ..."
	@rm -rf ./protos/goprotos;
	@mkdir ./protos/goprotos
	@protoc -I ./protos ./protos/*.proto  --go_out=./protos/goprotos
	@protoc -I ./protos ./protos/*.proto  --go-grpc_out=require_unimplemented_servers=false:./protos/goprotos

.PHONY: terraform-digitalocean-deploy
## terraform-digitalocean-deploy: deploy the current plan of terraform on Digitalocean
terraform-digitalocean-deploy:
	@echo "Terraform Digitalocean deploying ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} ./scripts/terraform_digitalocean_deploy.sh

.PHONY: terraform-digitalocean-destroy
## terraform-digitalocean-destroy: destroy the current plan of terraform on Digitalocean
terraform-digitalocean-destroy:
	@echo "Terraform Digitalocean destroying ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} ./scripts/terraform_digitalocean_destroy.sh

.PHONY: vagrant-deploy
## vagrant-deploy: deploy the vagrant locally
vagrant-deploy:
	@echo "Vagrant deploying locally ..."
	@cd ./deployment/vagrant/; vagrant up

.PHONY: vagrant-destroy
## vagrant-destroy: destroy the vagrant locally
vagrant-destroy:
	@echo "Vagrant destroy locally ..."
	@cd ./deployment/vagrant/; vagrant destroy -f

.PHONY: vagrant-suspend
## vagrant-suspend: suspend the vagrant locally
vagrant-suspend:
	@echo "Vagrant suspend locally ..."
	@cd ./deployment/vagrant/; vagrant suspend

.PHONY: vagrant-resume
## vagrant-resume: resume the vagrant locally
vagrant-resume:
	@echo "Vagrant resume locally ..."
	@cd ./deployment/vagrant/; vagrant resume

.PHONY: build-deploy-all-remote
## build-deploy-all-remote: build and deploy remote components using Ansible
build-deploy-all-remote: build-remote-linux
	@echo "Running the bash script for building and deploying components ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="remote" ./scripts/deploy_components.sh

.PHONY: build-deploy-all-local
## build-deploy-all-local: build and deploy local components using Ansible
build-deploy-all-local:  build-local-linux
	@echo "Running the bash script for building and deploying components ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="local" ./scripts/deploy_components.sh

.PHONY: prepare-remote-linux-env
## prepare-remote-linux-env: Create the remote linux build env
prepare-remote-linux-env:
	@echo "Preparing the remote linux build env ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="remote" ./scripts/prepare_linux_build_env.sh

.PHONY: prepare-local-vms
## prepare-local-vms: Install dependencies on local VMs
prepare-local-vms:
	@echo "Preparing the local linux build env ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="local" ./scripts/prepare_linux_build_env.sh

.PHONY: build-docker-images
## build-docker-images: build docker images
build-docker-images:
	@echo "Building Docker images ..."
	@docker build --force-rm -t ${DOCKER_NS_ENV}/geds-mds -f ./deployment/docker/images/mds.dockerfile .
	@docker tag ${DOCKER_NS_ENV}/geds-node ${DOCKER_NS_ENV}/crdtchain-mds:${DOCKER_TAG_ENV}

.PHONY: push-docker-images
## push-docker-images: push docker images
push-docker-images:
	@echo "Pushing Docker images ..."
	@docker login --username ${DOCKER_USERNAME_ENV} --password ${DOCKER_PASSWORD_ENV}
	@docker push ${DOCKER_NS_ENV}/geds-mds:${DOCKER_TAG_ENV}

.PHONY: install-local-docker
## install-local-docker: Install local docker
install-local-docker:
	@echo "Installing dockers locally ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="local" ./scripts/install_docker.sh

.PHONY: install-remote-docker
## install-remote-docker: Install remote docker
install-remote-docker:
	@echo "Installing dockers remote ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="remote" ./scripts/install_docker.sh

.PHONY: build-deploy-all-remote-docker
## build-deploy-all-remote-docker: build and deploy remote components using Ansible and Docker
build-deploy-all-remote-docker:build-docker-images push-docker-images
	@echo "Running the bash script for building and deploying components ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="remote" ./scripts/deploy_components_with_dcoker.sh

.PHONY: build-deploy-all-local-docker
## build-deploy-all-local-docker: build and deploy local components using Ansible and Docker
build-deploy-all-local-docker: build-docker-images push-docker-images
	@echo "Running the bash script for building and deploying components ..."
	@PROJECT_ABSOLUTE_PATH=${PROJECT_ABSOLUTE_PATH} BUILD_MODE="local" ./scripts/deploy_components_with_dcoker.sh

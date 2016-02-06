SHORT_NAME := kubedeploy
REPO_PATH := github.com/arschles/${SHORT_NAME}
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_CMD := docker run --rm -e CGO_ENABLED=0 -e GO15VENDOREXPERIMENT=1 -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR} quay.io/deis/go-dev:0.5.0

TEST_PACKAGES := $(shell ${DEV_ENV_CMD} glide nv)

DOCKER_REPO ?= quay.io/
DOCKER_ORG ?= arschles
DOCKER_IMAGE ?= kubedeploy
DOCKER_VERSION ?= $(shell git rev-parse HEAD)
DOCKER_IMAGE := ${DOCKER_REPO}${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_VERSION}

bootstrap:
	${DEV_ENV_CMD} glide install
glideget:
	${DEV_ENV_CMD} glide get ${PACKAGE}
build:
	${DEV_ENV_CMD} go build
test:
	${DEV_ENV_CMD} go test ${TEST_PACKAGES}
docker-build:
	docker build --rm -t ${DOCKER_IMAGE} .
docker-push:
	docker push ${DOCKER_IMAGE}

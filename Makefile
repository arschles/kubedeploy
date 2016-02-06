SHORT_NAME := kubedeploy
REPO_PATH := github.com/arschles/${SHORT_NAME}
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_CMD := docker run --rm -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR} quay.io/deis/go-dev:0.5.0

TEST_PACKAGES := $(shell ${DEV_ENV_CMD} glide nv)

bootstrap:
	${DEV_ENV_CMD} glide install
glideget:
	${DEV_ENV_CMD} glide get ${PACKAGE}
build:
	${DEV_ENV_CMD} go build
test:
	${DEV_ENV_CMD} go test ${TEST_PACKAGES}

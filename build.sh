#!/bin/sh

TAG=${1:-"latest"}

BUILD_NAME=${2:-"todalist-app"}


env DOCKER_BUILDKIT=1 podman build --build-arg APP_NAME=todalist-app -t ${BUILD_NAME}:${TAG} .

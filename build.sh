#!/bin/bash

APP_NAME=circularqueue
TMP_DIR=/tmp/$APP_NAME/
IMAGE_NAME=local/$APP_NAME

init() {
  rm -rf $TMP_DIR/ \
   && mkdir -p $TMP_DIR/
}

build() {
  docker build -t $IMAGE_NAME .
}

panic() {
  local message=$1
  echo $message
  exit 1
}

main() {
  init  || panic "init failed"
  build || panic "build failed"
}
main

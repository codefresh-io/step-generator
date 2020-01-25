#!/bin/bash
set -e

build () {
    VERSION=$(cat ./VERSION)
    packr build -o ./dist/step-generator .
}

echo "Building go binary"
build

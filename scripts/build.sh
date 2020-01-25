#!/bin/bash
set -e

build () {
    packr build -o ./dist/step-generator .
}

echo "Building go binary"
build

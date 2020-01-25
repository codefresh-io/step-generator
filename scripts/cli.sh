#!/bin/bash
set -e
cli-generator generate \
    --project-dir . \
    --language go \
     --spec ./build/cli.yaml \
    --go-package github.com/codefresh/step-generator \
    --verbose 
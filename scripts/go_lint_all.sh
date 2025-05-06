#!/bin/bash

# move to project root
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/..

# linter (requires bash)
golangci-lint run ./playgrounds/*

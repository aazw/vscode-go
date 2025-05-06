#!/bin/bash

# move to project root
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/..

# linter (bash必須. Bashの機能で*をファイル名に展開)
golangci-lint run ./playgrounds/*

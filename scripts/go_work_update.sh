#!/bin/bash

# move to project root
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/..

# sync ./playgrounds
go work use -r ./playgrounds

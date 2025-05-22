#!/bin/bash

# move to project root
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
cd ${SCRIPT_DIR}/..

# build all
set -euo pipefail

for dir in ./playgrounds/*; do
	if [ -f "$dir/go.mod" ]; then
		echo "🧹 Running go mod tidy in $dir"
		(cd "$dir" && go mod tidy)
	fi
done

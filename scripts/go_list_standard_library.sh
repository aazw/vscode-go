#!/bin/bash

# move to project root
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
cd ${SCRIPT_DIR}/..

# print header
echo "# Standard library"
echo ""
echo "https://pkg.go.dev/std"
echo ""
echo "| Package | Since Version | Description |"
echo "|---------|---------------|-------------|"

# print go std
go list -json std | jq -r '
      select(.Standard == true)
      | select(.ImportPath | test("(^|/)(?:internal|vendor)($|/)") | not)
      | .ImportPath as $p
      | (.Doc // "" | split("\n") | .[0]) as $syn
      | "| https://pkg.go.dev/\($p) | | \($syn) |"
    ' |
	column -t -s '|' -o '|'

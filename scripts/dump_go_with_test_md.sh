#!/bin/bash

# move to project root
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/..

set -euo pipefail
IFS=$'\n\t'

# 引数チェック
if [ $# -lt 1 ]; then
  echo "Usage: $0 <search_root_directory>" >&2
  exit 1
fi

TARGET_DIR="${1%/}"

# ディレクトリ存在確認
if [ ! -d "$TARGET_DIR" ]; then
  echo "Error: '$TARGET_DIR' is not a directory." >&2
  exit 1
fi

# go.mod の１行目を出力（存在すれば）
if [ -f "$TARGET_DIR/go.mod" ]; then
  echo "## ${TARGET_DIR}/go.mod"
  echo 
  echo '```go'
  cat "$TARGET_DIR/go.mod"
  echo '```'
  echo
fi

# .go ファイルを探して Markdown 化
find "$TARGET_DIR" \
    -type f \
    -name '*.go' \
    | sort \
    | while read -r file; do
        echo "## $file"
        echo
        echo '```go'
        cat "$file"
        echo '```'
        echo
    done

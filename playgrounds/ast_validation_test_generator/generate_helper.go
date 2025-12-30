package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	crErrors "github.com/cockroachdb/errors"
)

const helperFileName = "unittestgen_helper" + filenameSuffix

func GenerateHelperFile(dir string, packageName string) error {

	var buf bytes.Buffer
	fmt.Fprintf(&buf, `package %s

type EquivalencyType string

const (
	EquivalencyNone    EquivalencyType = "EquivalencyNone"
	EquivalencyValid   EquivalencyType = "EquivalencyValid"
	EquivalencyInvalid EquivalencyType = "EquivalencyInvalid"
)

type BoundaryType string

const (
	BoundaryNone       BoundaryType = "BoundaryNone"
	BoundaryLowerBelow BoundaryType = "BoundaryLowerBelow"
	BoundaryLower      BoundaryType = "BoundaryLower"
	BoundaryLowerAbove BoundaryType = "BoundaryLowerAbove"
	BoundaryUpperBelow BoundaryType = "BoundaryUpperBelow"
	BoundaryUpper      BoundaryType = "BoundaryUpper"
	BoundaryUpperAbove BoundaryType = "BoundaryUpperAbove"
)
`,
		packageName,
	)

	// 書き出すソースコードを整形
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return crErrors.Wrapf(err, "failed to format generated helper code for file: %s", helperFileName)
	}

	// ファイルに書き込み
	outputPath := filepath.Join(dir, helperFileName)
	err = os.WriteFile(outputPath, src, 0666)
	if err != nil {
		return crErrors.Wrapf(err, "failed to write helper code to file: %s", outputPath)
	}

	return nil
}

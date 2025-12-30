package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/urfave/cli/v3"
)

const filenameSuffix = "_autogen_test.go"

const (
	appName       = "unittestgen"
	appUsageShort = ""
	appUsageLong  = ""
)

// https://pkg.go.dev/myitcv.io/gogenerate#pkg-constants
//   GOARCH: arm64
//   GOFILE: target.go
//   GOLINE: 11
//   GOOS: linux
//   GOPACKAGE: main
//   GOPATH: /go

var (
	targetPackageName string
	targetFileName    string
	targetStructName  string
	outputFileName    string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	cmd := &cli.Command{
		Name:      appName,
		Usage:     appUsageShort,
		UsageText: appUsageLong,
		Flags:     []cli.Flag{},
		Action:    run,
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(context.Context, *cli.Command) error {

	// go:generateで実行されたとき自動で設定される環境変数値の取得
	gogenPackageName := os.Getenv("GOPACKAGE")        //どのパッケージのファイルに記述されたgo:genereateが実行されたかの情報
	gogenFileName := os.Getenv("GOFILE")              //どのファイルに記述されたgo:genereateが実行されたかの情報
	gogenLine, _ := strconv.Atoi(os.Getenv("GOLINE")) //どの行に記述されたgo:genereateが実行されたかの情報

	// 取得対象箇所の宣言
	targetPackageName = gogenPackageName
	targetFileName = gogenFileName
	targetLine := gogenLine + 1

	// ファイルの読み込み&パース
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, targetFileName, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// ターゲットとなる構造体の取得
	var structName string
	var structType *ast.StructType
	switch {
	case targetStructName == "":
		for _, decl := range node.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				stType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue // 構造体以外は除外（例: type Alias = string など）
				}

				start := fset.Position(typeSpec.Pos())
				if start.Line == targetLine { // go:generateが書かれた次の行の構造体定義を探す
					structName = typeSpec.Name.Name
					structType = stType
					break
				}
			}
		}
	case targetStructName != "":
		// TODO
	}

	// 目的の関数があったかチェック
	if structName == "" {
		log.Fatal("go:generate unittestgen 構造体が見つかりません")
	}

	// ヘルパーファイル生成
	err = GenerateHelperFile(".", targetPackageName)
	if err != nil {
		log.Fatalf("go:generate unittestgen テスト生成に失敗しました: %+v", err)
	}

	// 出力先設定
	outputFileName = fileNameWithoutExtension(targetFileName) + "_" + strings.ToLower(structName) + filenameSuffix

	// テスト生成
	err = generateTestForStruct(node, structType, structName)
	if err != nil {
		log.Fatalf("go:generate unittestgen テスト生成に失敗しました: %+v", err)
	}

	return nil
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

type StructField struct {
	Name         string
	Type         string
	ValidateTags []string
}

func generateTestForStruct(file *ast.File, structType *ast.StructType, structName string) error {

	structFields := make([]*StructField, 0)
	for _, field := range structType.Fields.List {
		// タグがなければスキップ
		if field.Tag == nil {
			continue
		}

		// validateタグ取得
		tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
		validateTagStr := tag.Get("validate")
		if validateTagStr == "" {
			continue
		}

		fieldName := field.Names[0].Name
		fieldTypeStr := fmt.Sprint(field.Type)

		// 分割
		validateTags := strings.Split(validateTagStr, ",")

		structFields = append(structFields, &StructField{
			Name:         fieldName,
			Type:         fieldTypeStr,
			ValidateTags: validateTags,
		})
	}

	testCases := GenerateTestCases(structFields)

	out := fmt.Sprintf(`package %s

var %sTestCases = []struct {
		name string
		params *%s
		field string
		validateTag string
		wantErr bool
		equivalency EquivalencyType
		boundary BoundaryType
	}{
`,
		targetPackageName,
		structName,
		structName,
	)

	for _, tc := range testCases {
		out += testCaseToString(tc)
	}

	out += `}`

	src, err := format.Source([]byte(out))
	if err != nil {
		log.Fatalf("go/format error: %v\n%s", err, out)
	}
	err = os.WriteFile(outputFileName, src, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func testCaseToString(tc TestCase) string {
	return fmt.Sprintf(`		{
			name: "%s",
			params: %s,
			field: "%s",
			validateTag: "%s",
			wantErr: %v,
			equivalency: %s,
			boundary: %s,
		},
`,
		tc.Name, tc.Params, tc.Field, tc.ValidateTag, tc.WantErr, tc.Equivalency, tc.Boundary)
}

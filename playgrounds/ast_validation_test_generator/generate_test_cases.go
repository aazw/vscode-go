package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

type TestCase struct {
	Name        string
	Params      string
	Field       string
	ValidateTag string
	WantErr     bool
	Equivalency EquivalencyType
	Boundary    BoundaryType
}

// バリデータタグからテストケースを生成（enum情報付き）
func GenerateTestCases(structFields []*StructField) []TestCase {
	cases := []TestCase{}

	for _, structField := range structFields {
		fieldName := structField.Name
		fieldTypeStr := structField.Type

		// 型判定
		isInt := fieldTypeStr == "int"
		isUint := fieldTypeStr == "uint"
		isFloat := fieldTypeStr == "float64" || fieldTypeStr == "float32"
		isString := fieldTypeStr == "string"

		// バリデーション範囲（下限・上限）をfloat64で抽出
		var lower *float64
		var upper *float64
		for _, tag := range structField.ValidateTags {
			var re = regexp.MustCompile(`(min|max|gte|gt|lte|lt)=([0-9.]+)`)
			matches := re.FindAllStringSubmatch(tag, -1)
			for _, m := range matches {
				key, valStr := m[1], m[2]
				val, _ := strconv.ParseFloat(valStr, 64)
				switch key {
				case "min", "gte":
					if lower == nil || val > *lower {
						lower = &val
					}
				case "gt":
					v := val + getStep(fieldTypeStr)
					if lower == nil || v > *lower {
						lower = &v
					}
				case "max", "lte":
					if upper == nil || val < *upper {
						upper = &val
					}
				case "lt":
					v := val - getStep(fieldTypeStr)
					if upper == nil || v < *upper {
						upper = &v
					}
				}
			}
		}

		for _, tag := range structField.ValidateTags {
			var re = regexp.MustCompile(`(min|max|gte|gt|lte|lt)=([0-9.]+)`)
			matches := re.FindAllStringSubmatch(tag, -1)

			if tag == "required" {
				cases = append(cases, TestCase{
					Name:        fmt.Sprintf("%s is empty", fieldName),
					Params:      fmt.Sprintf("&InputParams{%s: %s}", fieldName, zeroValue(fieldTypeStr)),
					Field:       fieldName,
					ValidateTag: tag,
					WantErr:     true,
					Equivalency: EquivalencyInvalid,
					Boundary:    BoundaryNone,
				})
			}

			if len(matches) > 0 {
				for _, m := range matches {
					key, valStr := m[1], m[2]
					val, _ := strconv.ParseFloat(valStr, 64)
					step := getStep(fieldTypeStr)
					if (key == "min" || key == "gte") && isString {
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s length lower below", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val-step),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     true,
							Equivalency: EquivalencyInvalid,
							Boundary:    BoundaryLowerBelow,
						})
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s length lower", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     false,
							Equivalency: EquivalencyValid,
							Boundary:    BoundaryLower,
						})
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s length lower above", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val+step),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     false,
							Equivalency: EquivalencyValid,
							Boundary:    BoundaryLowerAbove,
						})
					}
					if (key == "max" || key == "lte") && isString {
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s length upper below", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val-step),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     false,
							Equivalency: EquivalencyValid,
							Boundary:    BoundaryUpperBelow,
						})
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s length upper", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     false,
							Equivalency: EquivalencyValid,
							Boundary:    BoundaryUpper,
						})
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s length upper above", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val+step),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     true,
							Equivalency: EquivalencyInvalid,
							Boundary:    BoundaryUpperAbove,
						})
					}
					// int/uint/float型の境界値
					if isInt || isUint || isFloat {
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s lower below", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val-step),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     true,
							Equivalency: EquivalencyInvalid,
							Boundary:    BoundaryLowerBelow,
						})
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s lower", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     false,
							Equivalency: EquivalencyValid,
							Boundary:    BoundaryLower,
						})
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s lower above", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val+step),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     false,
							Equivalency: EquivalencyValid,
							Boundary:    BoundaryLowerAbove,
						})
					}
					if (key == "max" || key == "lte" || key == "lt") && (isInt || isUint || isFloat) {
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s upper below", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val-step),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     false,
							Equivalency: EquivalencyValid,
							Boundary:    BoundaryUpperBelow,
						})
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s upper", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     false,
							Equivalency: EquivalencyValid,
							Boundary:    BoundaryUpper,
						})
						cases = append(cases, TestCase{
							Name:        fmt.Sprintf("%s upper above", fieldName),
							Params:      boundaryParam(fieldTypeStr, fieldName, val+step),
							Field:       fieldName,
							ValidateTag: tag,
							WantErr:     true,
							Equivalency: EquivalencyInvalid,
							Boundary:    BoundaryUpperAbove,
						})
					}
				}
				// lower/upperが両方ある場合、代表値を生成
				if lower != nil && upper != nil {
					mid := (*lower + *upper) / 2
					cases = append(cases, TestCase{
						Name:        fmt.Sprintf("%s valid value", fieldName),
						Params:      boundaryParam(fieldTypeStr, fieldName, mid),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					})
					cases = append(cases, TestCase{
						Name:        fmt.Sprintf("%s invalid value", fieldName),
						Params:      boundaryParam(fieldTypeStr, fieldName, *upper+10*getStep(fieldTypeStr)),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					})
				}
			}

			// --- go-playground/validator 代表的なタグ ---
			if tag == "iscolor" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid color",
						Params:      fmt.Sprintf("&InputParams{%s: \"#FFFFFF\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid color",
						Params:      fmt.Sprintf("&InputParams{%s: \"notacolor\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "country_code" || tag == "eu_country_code" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid country code",
						Params:      fmt.Sprintf("&InputParams{%s: \"JP\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid country code",
						Params:      fmt.Sprintf("&InputParams{%s: \"ZZ\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "alpha" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid alpha",
						Params:      fmt.Sprintf("&InputParams{%s: \"abc\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid alpha",
						Params:      fmt.Sprintf("&InputParams{%s: \"abc123\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "boolean" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid boolean",
						Params:      fmt.Sprintf("&InputParams{%s: true}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid boolean",
						Params:      fmt.Sprintf("&InputParams{%s: \"notabool\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "uuid" || tag == "uuid3" || tag == "uuid4" || tag == "uuid5" || tag == "uuid_rfc4122" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid uuid",
						Params:      fmt.Sprintf("&InputParams{%s: \"123e4567-e89b-12d3-a456-426614174000\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid uuid",
						Params:      fmt.Sprintf("&InputParams{%s: \"notauuid\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "email" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " invalid email",
						Params:      fmt.Sprintf("&InputParams{%s: \"notanemail\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " valid email",
						Params:      fmt.Sprintf("&InputParams{%s: \"a@b.com\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "url" || tag == "http_url" || tag == "uri" || tag == "urn_rfc2141" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid url",
						Params:      fmt.Sprintf("&InputParams{%s: \"https://example.com\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid url",
						Params:      fmt.Sprintf("&InputParams{%s: \"not_a_url\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "json" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid json",
						Params:      fmt.Sprintf("&InputParams{%s: \"{\\\"a\\\":1}\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid json",
						Params:      fmt.Sprintf("&InputParams{%s: \"notjson\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "datetime" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid datetime",
						Params:      fmt.Sprintf("&InputParams{%s: \"2023-01-01T00:00:00Z\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid datetime",
						Params:      fmt.Sprintf("&InputParams{%s: \"notadatetime\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "base64" || tag == "base64url" || tag == "base64rawurl" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid base64",
						Params:      fmt.Sprintf("&InputParams{%s: \"aGVsbG8=\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid base64",
						Params:      fmt.Sprintf("&InputParams{%s: \"notbase64!!\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "contains" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid contains",
						Params:      fmt.Sprintf("&InputParams{%s: \"foobar\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid contains",
						Params:      fmt.Sprintf("&InputParams{%s: \"bar\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "startswith" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid startswith",
						Params:      fmt.Sprintf("&InputParams{%s: \"foobar\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid startswith",
						Params:      fmt.Sprintf("&InputParams{%s: \"barfoo\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
			if tag == "endswith" {
				cases = append(cases,
					TestCase{
						Name:        fieldName + " valid endswith",
						Params:      fmt.Sprintf("&InputParams{%s: \"foobar\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     false,
						Equivalency: EquivalencyValid,
						Boundary:    BoundaryNone,
					},
					TestCase{
						Name:        fieldName + " invalid endswith",
						Params:      fmt.Sprintf("&InputParams{%s: \"barfoo\"}", fieldName),
						Field:       fieldName,
						ValidateTag: tag,
						WantErr:     true,
						Equivalency: EquivalencyInvalid,
						Boundary:    BoundaryNone,
					},
				)
			}
		}
	}

	return cases
}

// 型ごとに値を生成
func boundaryParam(typeStr, name string, val float64) string {
	switch typeStr {
	case "string":
		return fmt.Sprintf("&InputParams{%s: %q}", name, strings.Repeat("a", int(val)))
	case "int":
		return fmt.Sprintf("&InputParams{%s: %d}", name, int(val))
	case "uint":
		return fmt.Sprintf("&InputParams{%s: %d}", name, uint(val))
	case "float64", "float32":
		return fmt.Sprintf("&InputParams{%s: %f}", name, val)
	default:
		return "nil"
	}
}

// 型ごとにゼロ値
func zeroValue(typeStr string) string {
	switch typeStr {
	case "string":
		return "\"\""
	case "int", "uint":
		return "0"
	case "float64", "float32":
		return "0.0"
	default:
		return "nil"
	}
}

// 型ごとに刻み幅
func getStep(typeStr string) float64 {
	switch typeStr {
	case "float64", "float32":
		return 0.1
	default:
		return 1
	}
}

package validatorx

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// 1) ルール型
type ruleKind int8

const (
	NoParam ruleKind = iota
	SingleParam
	FieldCompare
)

// 2) ルール定義
type RuleSpec struct {
	Kind       ruleKind // NoParam／SingleParam／FieldCompare
	Desc       string   // e.g. "a valid email", "length"
	Verb       string   // FieldCompare 用 e.g. "equal to"
	NeedsBe    bool     // “must be …” にするか (“must …” vs “must be …”)
	QuoteParam bool     // SingleParam のときパラメータを '…' で囲むか
}

// Core
// Shortcut
func NP(desc string) RuleSpec     { return RuleSpec{NoParam, desc, "", true, false} }      // NoParam
func NPnoBe(desc string) RuleSpec { return RuleSpec{NoParam, desc, "", false, false} }     // NoParam without "be" verb
func SP(desc string) RuleSpec     { return RuleSpec{SingleParam, desc, "", true, false} }  // SingleParam
func SPq(desc string) RuleSpec    { return RuleSpec{SingleParam, desc, "", true, true} }   // SingleParam with quota
func FC(verb string) RuleSpec     { return RuleSpec{FieldCompare, "", verb, true, false} } // FieldCompare

var ruleSet = map[string]RuleSpec{
	// --- Field ↔ Field ---
	"eqcsfield":     FC("equal to"),
	"eqfield":       FC("equal to"),
	"fieldcontains": {FieldCompare, "", "contain", false, false},
	"fieldexcludes": {FieldCompare, "", "not contain", false, false},
	"gtcsfield":     FC("greater than"),
	"gtecsfield":    FC("greater than or equal to"),
	"gtefield":      FC("greater than or equal to"),
	"gtfield":       FC("greater than"),
	"ltcsfield":     FC("less than"),
	"ltecsfield":    FC("less than or equal to"),
	"ltefield":      FC("less than or equal to"),
	"ltfield":       FC("less than"),
	"necsfield":     FC("different from"),
	"nefield":       FC("different from"),

	// --- Network ---
	"cidr":              NP("a valid CIDR"),
	"cidrv4":            NP("a valid CIDRv4"),
	"cidrv6":            NP("a valid CIDRv6"),
	"datauri":           NP("a valid Data URI"),
	"fqdn":              NP("a valid FQDN"),
	"hostname":          NP("a valid hostname"),
	"hostname_port":     NP("a valid host:port"),
	"port":              NP("a valid TCP/UDP port"),
	"hostname_rfc1123":  NP("a valid RFC 1123 hostname"),
	"dns_rfc1035_label": NP("a valid RFC1035 DNS label"),
	"ip":                NP("a valid IP address"),
	"ip4_addr":          NP("a valid IPv4 address"),
	"ip6_addr":          NP("a valid IPv6 address"),
	"ip_addr":           NP("a valid IP address"),
	"ipv4":              NP("a valid IPv4 address"),
	"ipv6":              NP("a valid IPv6 address"),
	"mac":               NP("a valid MAC address"),
	"tcp4_addr":         NP("a valid TCPv4 address"),
	"tcp6_addr":         NP("a valid TCPv6 address"),
	"tcp_addr":          NP("a valid TCP address"),
	"udp4_addr":         NP("a valid UDPv4 address"),
	"udp6_addr":         NP("a valid UDPv6 address"),
	"udp_addr":          NP("a valid UDP address"),
	"unix_addr":         NP("a valid Unix-domain address"),
	"uri":               NP("a valid URI"),
	"url":               NP("a valid URL"),
	"http_url":          NP("a valid HTTP URL"),
	"url_encoded":       NPnoBe("URL-encoded data"), // “must URL-encoded data” ではなく “must be URL-encoded data”
	"urn_rfc2141":       NP("a valid RFC 2141 URN"),

	// --- Strings ---
	"alpha":           NPnoBe("alphabetic characters"), // “must alphabetic” のほうが自然
	"alphanum":        NPnoBe("alphanumeric characters"),
	"alphanumunicode": NPnoBe("alphanumeric Unicode characters"),
	"alphaunicode":    NPnoBe("alphabetic Unicode characters"),
	"ascii":           NPnoBe("ASCII characters"),
	"boolean":         NP("a boolean"),
	"contains":        SPq("contain"),
	"containsany":     SPq("contain any of"),
	"containsrune":    SPq("contain rune"),
	"endsnotwith":     SPq("not end with"),
	"endswith":        SPq("end with"),
	"excludes":        SPq("not contain"),
	"excludesall":     SPq("contain none of"),
	"excludesrune":    SPq("not contain rune"),
	"lowercase":       NPnoBe("lower-case"),
	"multibyte":       NPnoBe("multi-byte characters"),
	"number":          NP("a number"),
	"numeric":         NP("numeric"),
	"printascii":      NPnoBe("printable ASCII"),
	"startsnotwith":   SPq("not start with"),
	"startswith":      SPq("start with"),
	"uppercase":       NPnoBe("upper-case"),

	// --- Format ---
	"base32":                        NP("a valid Base32"),
	"base64":                        NP("a valid Base64"),
	"base64url":                     NP("a valid Base64URL"),
	"base64rawurl":                  NP("a valid Base64RawURL"),
	"bic":                           NP("a valid BIC"),
	"bcp47_language_tag":            NP("a valid BCP-47 language tag"),
	"btc_addr":                      NP("a valid Bitcoin address"),
	"btc_addr_bech32":               NP("a valid Bech32 Bitcoin address"),
	"credit_card":                   NP("a valid credit-card number"),
	"mongodb":                       NP("a valid MongoDB ObjectID"),
	"mongodb_connection_string":     NP("a valid MongoDB connection string"),
	"cron":                          NP("a valid cron expression"),
	"spicedb":                       NP("a valid SpiceDB identifier"),
	"datetime":                      NPnoBe("a valid datetime"),
	"e164":                          NP("a valid E.164 phone number"),
	"ein":                           NP("a valid EIN"),
	"email":                         NP("a valid email"),
	"eth_addr":                      NP("a valid Ethereum address"),
	"eth_addr_checksum":             NP("a checksummed Ethereum address"),
	"hexadecimal":                   NP("hexadecimal"),
	"hexcolor":                      NP("a valid hex color"),
	"hsl":                           NP("a valid HSL color"),
	"hsla":                          NP("a valid HSLA color"),
	"html":                          NPnoBe("HTML tags"),
	"html_encoded":                  NPnoBe("HTML-encoded text"),
	"isbn":                          NP("a valid ISBN"),
	"isbn10":                        NP("a valid ISBN-10"),
	"isbn13":                        NP("a valid ISBN-13"),
	"issn":                          NP("a valid ISSN"),
	"iso3166_1_alpha2":              NP("a valid ISO-3166-1 alpha-2 code"),
	"iso3166_1_alpha2_eu":           NP("a valid ISO-3166-1 alpha-2 EU code"),
	"iso3166_1_alpha3":              NP("a valid ISO-3166-1 alpha-3 code"),
	"iso3166_1_alpha3_eu":           NP("a valid ISO-3166-1 alpha-3 EU code"),
	"iso3166_1_alpha_numeric":       NP("a valid ISO-3166-1 numeric code"),
	"iso3166_1_alpha_numeric_eu":    NP("a valid ISO-3166-1 numeric EU code"),
	"iso3166_2":                     NP("a valid ISO-3166-2 code"),
	"iso4217":                       NP("a valid ISO-4217 currency code"),
	"iso4217_numeric":               NP("a valid ISO-4217 numeric currency code"),
	"json":                          NP("valid JSON"),
	"jwt":                           NP("a valid JWT"),
	"latitude":                      NP("a valid latitude"),
	"longitude":                     NP("a valid longitude"),
	"luhn_checksum":                 NP("a valid Luhn checksum"),
	"postcode_iso3166_alpha2":       NP("a valid postcode"),
	"postcode_iso3166_alpha2_field": NP("a valid postcode"),
	"rgb":                           NP("a valid RGB color"),
	"rgba":                          NP("a valid RGBA color"),
	"ssn":                           NP("a valid SSN"),
	"timezone":                      NP("a valid timezone"),
	"uuid":                          NP("a valid UUID"),
	"uuid3":                         NP("a valid UUID v3"),
	"uuid3_rfc4122":                 NP("a valid UUID v3 (RFC4122)"),
	"uuid4":                         NP("a valid UUID v4"),
	"uuid4_rfc4122":                 NP("a valid UUID v4 (RFC4122)"),
	"uuid5":                         NP("a valid UUID v5"),
	"uuid5_rfc4122":                 NP("a valid UUID v5 (RFC4122)"),
	"uuid_rfc4122":                  NP("a valid UUID (RFC4122)"),
	"md4":                           NP("a valid MD4 hash"),
	"md5":                           NP("a valid MD5 hash"),
	"sha256":                        NP("a valid SHA-256 hash"),
	"sha384":                        NP("a valid SHA-384 hash"),
	"sha512":                        NP("a valid SHA-512 hash"),
	"ripemd128":                     NP("a valid RIPEMD-128 hash"),
	"ripemd160":                     NP("a valid RIPEMD-160 hash"),
	"tiger128":                      NP("a valid TIGER-128 hash"),
	"tiger160":                      NP("a valid TIGER-160 hash"),
	"tiger192":                      NP("a valid TIGER-192 hash"),
	"semver":                        NP("a valid semantic version"),
	"ulid":                          NP("a valid ULID"),
	"cve":                           NP("a valid CVE identifier"),

	// --- Comparisons ---
	"eq":             {SingleParam, "", "equal to", true, true},
	"eq_ignore_case": {SingleParam, "", "equal (case-insensitive) to", true, true},
	"ne":             {SingleParam, "not equal to", "", true, true},
	"ne_ignore_case": {SingleParam, "not equal (case-insensitive) to", "", true, true},
	"gt":             SP("greater than"),
	"gte":            SP("at least"),
	"lt":             SP("less than"),
	"lte":            SP("at most"),

	// --- Other ---
	"dir":                  NP("an existing directory"),
	"dirpath":              NP("a valid directory path"),
	"file":                 NP("an existing file"),
	"filepath":             NP("a valid file path"),
	"image":                NP("a valid image"),
	"isdefault":            NPnoBe("the default value"),
	"len":                  SP("length"),
	"max":                  SP("at most"),
	"min":                  SP("at least"),
	"oneof":                SPq("one of"),
	"oneofci":              SPq("one of (case-insensitive)"),
	"required":             NP("set"),
	"required_if":          SP("set if"),
	"required_unless":      SP("set unless"),
	"skip_unless":          SP("validated unless"),
	"required_with":        SP("set with"),
	"required_with_all":    SP("set with all"),
	"required_without":     SP("set without"),
	"required_without_all": SP("set without all"),
	"excluded_if":          SP("excluded if"),
	"excluded_unless":      SP("excluded unless"),
	"excluded_with":        SP("excluded with"),
	"excluded_with_all":    SP("excluded with all"),
	"excluded_without":     SP("excluded without"),
	"excluded_without_all": SP("excluded without all"),
	"unique":               NPnoBe("unique"),
	"validateFn":           NPnoBe("valid according to Validate()"),

	// --- Aliases ---
	"iscolor":         NP("a valid color (hex, rgb[a], hsl[a])"),
	"country_code":    NP("a valid country code"),
	"eu_country_code": NP("a valid EU country code"),
}

// 3) 汎用フォールバック
func fallback(fe validator.FieldError) string {
	f := fe.Field()
	if fe.Param() == "" {
		return fmt.Sprintf("%s must be %s, but %s is %v", f, fe.Tag(), f, fe.Value())
	}
	return fmt.Sprintf("%s must be %s %s, but %s is %v", f, fe.Tag(), fe.Param(), f, fe.Value())
}

// 3) 共通フォーマッタ
func fmtVal(v any) string {
	switch vv := v.(type) {
	case string:
		return "'" + vv + "'" // 空文字は '' と出す
	default:
		return fmt.Sprintf("%v", vv)
	}
}

// 4) 公開 API
func camelToSnake(s string) string {
	var b strings.Builder
	for i, r := range s {
		if r == '_' { // 既存の '_' はそのまま・連続を抑制
			if b.Len() > 0 && b.String()[b.Len()-1] != '_' {
				b.WriteByte('_')
			}
			continue
		}

		if unicode.IsUpper(r) {
			// 前が小文字  または   次が小文字（HTML → HTM_L にならない）
			if i > 0 && s[i-1] != '_' &&
				(unicode.IsLower(rune(s[i-1])) ||
					(i+1 < len(s) && unicode.IsLower(rune(s[i+1])))) {
				b.WriteByte('_')
			}
			r = unicode.ToLower(r)
		}
		b.WriteRune(r)
	}
	return strings.Trim(b.String(), "_") // 先頭/末尾 '_' を除去
}

func Message(top reflect.Value, fe validator.FieldError) string {
	rule, ok := ruleSet[fe.Tag()]
	if !ok {
		return fallback(fe)
	}

	// エラー対象フィールド名と値
	fieldName := fe.Namespace()[strings.IndexByte(fe.Namespace(), '.')+1:]
	valStr := fmtVal(fe.Value())

	// 元のタグパラメータ名を保持
	rawParam := fe.Param()
	param := rawParam

	// Param を “JSON 名 or human readable” に変換
	if rule.Kind == FieldCompare && param != "" {
		if f, ok := top.Type().FieldByName(rawParam); ok {
			if tag := f.Tag.Get("json"); tag != "" && tag != "-" {
				param = strings.Split(tag, ",")[0] // json:"foo,omitempty" → foo
			} else {
				param = camelToSnake(param) // fallback
			}
		}
	}

	// QuoteParam があれば 'param' に
	if rule.QuoteParam {
		param = fmtVal(param)
	}

	// “must” vs “must be”
	verb := "must"
	if rule.NeedsBe {
		verb = "must be"
	}

	switch rule.Kind {
	case NoParam:
		// "f_field must be a valid email, but f_field is 'xxx'"
		return fmt.Sprintf("%s %s %s, but %s is %s",
			fieldName, verb, rule.Desc, fieldName, valStr)

	case SingleParam:
		// "f_field must contain 'abc', but f_field is 'xxx'"
		return fmt.Sprintf("%s %s %s %s, but %s is %s",
			fieldName, verb, rule.Desc, param, fieldName, valStr)

	case FieldCompare:
		// 複数パラメータはスペース区切り
		param = strings.ReplaceAll(param, ",", " ")
		// 参照先フィールドの値を取得
		var refVal string
		if f := top.FieldByName(rawParam); f.IsValid() {
			refVal = fmtVal(f.Interface())
		}
		// フィールド名はクォートせず、参照先の値を ` (value: ...)` として追加
		return fmt.Sprintf("%s %s %s %s (value: %s), but %s is %s",
			fieldName, verb, rule.Verb, param, refVal, fieldName, valStr)

	default:
		return fallback(fe)
	}
}

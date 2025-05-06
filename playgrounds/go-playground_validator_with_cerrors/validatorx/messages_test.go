package validatorx

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

// 以下の2つの定義をvalidatorの実装から持ってくる
var bakedTags = []string{
	// bakedInAliases
	// https://github.com/go-playground/validator/blob/72c45031a9e850aa51997745e8f8afa2723092c0/baked_in.go#L70
	"iscolor",
	"country_code",
	"eu_country_code",

	// bakedInValidators
	// https://github.com/go-playground/validator/blob/72c45031a9e850aa51997745e8f8afa2723092c0/baked_in.go#L79
	"required",
	"required_if",
	"required_unless",
	"skip_unless",
	"required_with",
	"required_with_all",
	"required_without",
	"required_without_all",
	"excluded_if",
	"excluded_unless",
	"excluded_with",
	"excluded_with_all",
	"excluded_without",
	"excluded_without_all",
	"isdefault",
	"len",
	"min",
	"max",
	"eq",
	"eq_ignore_case",
	"ne",
	"ne_ignore_case",
	"lt",
	"lte",
	"gt",
	"gte",
	"eqfield",
	"eqcsfield",
	"necsfield",
	"gtcsfield",
	"gtecsfield",
	"ltcsfield",
	"ltecsfield",
	"nefield",
	"gtefield",
	"gtfield",
	"ltefield",
	"ltfield",
	"fieldcontains",
	"fieldexcludes",
	"alpha",
	"alphanum",
	"alphaunicode",
	"alphanumunicode",
	"boolean",
	"numeric",
	"number",
	"hexadecimal",
	"hexcolor",
	"rgb",
	"rgba",
	"hsl",
	"hsla",
	"e164",
	"email",
	"url",
	"http_url",
	"uri",
	"urn_rfc2141",
	"file",
	"filepath",
	"base32",
	"base64",
	"base64url",
	"base64rawurl",
	"contains",
	"containsany",
	"containsrune",
	"excludes",
	"excludesall",
	"excludesrune",
	"startswith",
	"endswith",
	"startsnotwith",
	"endsnotwith",
	"image",
	"isbn",
	"isbn10",
	"isbn13",
	"issn",
	"eth_addr",
	"eth_addr_checksum",
	"btc_addr",
	"btc_addr_bech32",
	"uuid",
	"uuid3",
	"uuid4",
	"uuid5",
	"uuid_rfc4122",
	"uuid3_rfc4122",
	"uuid4_rfc4122",
	"uuid5_rfc4122",
	"ulid",
	"md4",
	"md5",
	"sha256",
	"sha384",
	"sha512",
	"ripemd128",
	"ripemd160",
	"tiger128",
	"tiger160",
	"tiger192",
	"ascii",
	"printascii",
	"multibyte",
	"datauri",
	"latitude",
	"longitude",
	"ssn",
	"ipv4",
	"ipv6",
	"ip",
	"cidrv4",
	"cidrv6",
	"cidr",
	"tcp4_addr",
	"tcp6_addr",
	"tcp_addr",
	"udp4_addr",
	"udp6_addr",
	"udp_addr",
	"ip4_addr",
	"ip6_addr",
	"ip_addr",
	"unix_addr",
	"mac",
	"hostname",
	"hostname_rfc1123",
	"fqdn",
	"unique",
	"oneof",
	"oneofci",
	"html",
	"html_encoded",
	"url_encoded",
	"dir",
	"dirpath",
	"json",
	"jwt",
	"hostname_port",
	"port",
	"lowercase",
	"uppercase",
	"datetime",
	"timezone",
	"iso3166_1_alpha2",
	"iso3166_1_alpha2_eu",
	"iso3166_1_alpha3",
	"iso3166_1_alpha3_eu",
	"iso3166_1_alpha_numeric",
	"iso3166_1_alpha_numeric_eu",
	"iso3166_2",
	"iso4217",
	"iso4217_numeric",
	"bcp47_language_tag",
	"postcode_iso3166_alpha2",
	"postcode_iso3166_alpha2_field",
	"bic",
	"semver",
	"dns_rfc1035_label",
	"credit_card",
	"cve",
	"luhn_checksum",
	"mongodb",
	"mongodb_connection_string",
	"cron",
	"spicedb",
	"ein",
	"validateFn", // Verify if the method Validate() error does not return an error (or any specified method)
}

// バリデーションで常に PASS してしまう既知のタグ
var skipCoverage = []string{
	// 実装上、バリデーションが機能していない. Issueにはなっている.
	// https://github.com/go-playground/validator/issues/1348
	"unix_addr",
}

// 空値 (ゼロ値) だと常に PASS になる既知タグ
// - max / min / ascii などは「空なら OK」という仕様
// - パラメータ比較や条件タグで参照フィールドが空だとスキップされる
var skipEmpty = []string{
	"ascii",
	"printascii",
	"multibyte",
	"max",
	"min",
	"isdefault",
	"unique",

	// 条件付き required / excluded 系
	"required_if",
	"required_unless",
	"required_with",
	"required_with_all",
	"required_without",
	"required_without_all",
	"excluded_if",
	"excluded_unless",
	"excluded_with",
	"excluded_with_all",
	"excluded_without",
	"excluded_without_all",
	"skip_unless",

	// フィールド比較系 (参照フィールドが空でスキップ)
	"eqfield",
	"eqcsfield",
	"nefield",
	"necsfield",
	"gtefield",
	"gtecsfield",
	"gtfield",
	"gtcsfield",
	"ltefield",
	"ltecsfield",
	"ltfield",
	"ltcsfield",
	"fieldcontains",

	// 文字列 param 系で空が許可される
	"startsnotwith",
	"endsnotwith",
	"excludes",
	"excludesall",
	"excludesrune",
	"url_encoded",
	"datetime",

	// AllTagsの定義上、空→ゼロ値だとPASSする
	"ne",
	"ne_ignore_case",
	"lt",
	"lte",
}

type AllTags struct {
	// ---------- Field ↔ Field ----------
	F_EqCSFieldA       int    `json:"f_eq_cs_field_a"  validate:"eqcsfield=F_EqCSFieldB"`
	F_EqCSFieldB       int    // 参照用
	F_EqFieldA         string `json:"f_eq_field_a"     validate:"eqfield=F_EqFieldB"`
	F_EqFieldB         string // 参照用
	F_FieldContains    string `json:"f_field_contains" validate:"fieldcontains=F_FieldContainsSrc"`
	F_FieldContainsSrc string // 参照用
	F_FieldExcludes    string `json:"f_field_excludes" validate:"fieldexcludes=F_FieldExcludesSrc"`
	F_FieldExcludesSrc string // 参照用
	F_GtCSFieldA       int    `json:"f_gt_cs_field_a"  validate:"gtcsfield=F_GtCSFieldB"`
	F_GtCSFieldB       int    // 参照用
	F_GteCSFieldA      int    `json:"f_gte_cs_field_a" validate:"gtecsfield=F_GteCSFieldB"`
	F_GteCSFieldB      int    // 参照用
	F_GteFieldA        int    `json:"f_gte_field_a"    validate:"gtefield=F_GteFieldB"`
	F_GteFieldB        int    // 参照用
	F_GtFieldA         int    `json:"f_gt_field_a"     validate:"gtfield=F_GtFieldB"`
	F_GtFieldB         int    // 参照用
	F_LtCSFieldA       int    `json:"f_lt_cs_field_a"  validate:"ltcsfield=F_LtCSFieldB"`
	F_LtCSFieldB       int    // 参照用
	F_LteCSFieldA      int    `json:"f_lte_cs_field_a" validate:"ltecsfield=F_LteCSFieldB"`
	F_LteCSFieldB      int    // 参照用
	F_LteFieldA        int    `json:"f_lte_field_a"    validate:"ltefield=F_LteFieldB"`
	F_LteFieldB        int    // 参照用
	F_LtFieldA         int    `json:"f_lt_field_a"     validate:"ltfield=F_LtFieldB"`
	F_LtFieldB         int    // 参照用
	F_NeCSFieldA       int    `json:"f_ne_cs_field_a"  validate:"necsfield=F_NeCSFieldB"`
	F_NeCSFieldB       int    // 参照用
	F_NeFieldA         int    `json:"f_ne_field_a"     validate:"nefield=F_NeFieldB"`
	F_NeFieldB         int    // 参照用

	// ---------- Network ----------
	N_CIDR            string `json:"n_cidr"             validate:"cidr"`
	N_CIDRv4          string `json:"n_cidr_v4"          validate:"cidrv4"`
	N_CIDRv6          string `json:"n_cidr_v6"          validate:"cidrv6"`
	N_DataURI         string `json:"n_data_uri"         validate:"datauri"`
	N_FQDN            string `json:"n_fqdn"             validate:"fqdn"`
	N_Hostname        string `json:"n_hostname"         validate:"hostname"`
	N_HostnamePort    string `json:"n_hostname_port"    validate:"hostname_port"`
	N_Port            uint16 `json:"n_port"             validate:"port"`
	N_HostnameRFC1123 string `json:"n_hostname_rfc1123" validate:"hostname_rfc1123"`
	N_DNSLabel        string `json:"n_dns_label"        validate:"dns_rfc1035_label"`
	N_IP              string `json:"n_ip"               validate:"ip"`
	N_IP4Addr         string `json:"n_ip4_addr"         validate:"ip4_addr"`
	N_IP6Addr         string `json:"n_ip6_addr"         validate:"ip6_addr"`
	N_IPAddr          string `json:"n_ip_addr"          validate:"ip_addr"`
	N_IPv4            string `json:"n_ip_v4"            validate:"ipv4"`
	N_IPv6            string `json:"n_ip_v6"            validate:"ipv6"`
	N_MAC             string `json:"n_mac"              validate:"mac"`
	N_TCP4Addr        string `json:"n_tcp4_addr"        validate:"tcp4_addr"`
	N_TCP6Addr        string `json:"n_tcp6_addr"        validate:"tcp6_addr"`
	N_TCPAddr         string `json:"n_tcp_addr"         validate:"tcp_addr"`
	N_UDP4Addr        string `json:"n_udp4_addr"        validate:"udp4_addr"`
	N_UDP6Addr        string `json:"n_udp6_addr"        validate:"udp6_addr"`
	N_UDPAddr         string `json:"n_udp_addr"         validate:"udp_addr"`
	N_UnixAddr        string `json:"n_unix_addr"        validate:"unix_addr"`
	N_URI             string `json:"n_uri"              validate:"uri"`
	N_URL             string `json:"n_url"              validate:"url"`
	N_HTTPURL         string `json:"n_http_url"         validate:"http_url"`
	N_URLEncoded      string `json:"n_url_encoded"      validate:"url_encoded"`
	N_URN             string `json:"n_urn"              validate:"urn_rfc2141"`

	// ---------- Strings ----------
	S_Alpha         string `json:"s_alpha"           validate:"alpha"`
	S_Alnum         string `json:"s_alnum"           validate:"alphanum"`
	S_AlnumUnicode  string `json:"s_alnum_unicode"   validate:"alphanumunicode"`
	S_AlphaUnicode  string `json:"s_alpha_unicode"   validate:"alphaunicode"`
	S_ASCII         string `json:"s_ascii"           validate:"ascii"`
	S_Bool          string `json:"s_bool"            validate:"boolean"`
	S_Contains      string `json:"s_contains"        validate:"contains=xyz"`
	S_ContainsAny   string `json:"s_contains_any"    validate:"containsany=xyz"`
	S_ContainsRune  string `json:"s_contains_rune"   validate:"containsrune=😊"`
	S_EndsNotWith   string `json:"s_ends_not_with"   validate:"endsnotwith=abc"`
	S_EndsWith      string `json:"s_ends_with"       validate:"endswith=abc"`
	S_Excludes      string `json:"s_excludes"        validate:"excludes=xyz"`
	S_ExcludesAll   string `json:"s_excludes_all"    validate:"excludesall=abc"`
	S_ExcludesRune  string `json:"s_excludes_rune"   validate:"excludesrune=😊"`
	S_Lowercase     string `json:"s_lowercase"       validate:"lowercase"`
	S_Multibyte     string `json:"s_multibyte"       validate:"multibyte"`
	S_Number        string `json:"s_number"          validate:"number"`
	S_Numeric       string `json:"s_numeric"         validate:"numeric"`
	S_PrintASCII    string `json:"s_print_ascii"     validate:"printascii"`
	S_StartsNotWith string `json:"s_starts_not_with" validate:"startsnotwith=abc"`
	S_StartsWith    string `json:"s_starts_with"     validate:"startswith=abc"`
	S_Uppercase     string `json:"s_uppercase"       validate:"uppercase"`

	// ---------- Format ----------
	F_Base32        string `json:"f_base32"         validate:"base32"`
	F_Base64        string `json:"f_base64"         validate:"base64"`
	F_Base64URL     string `json:"f_base64_url"     validate:"base64url"`
	F_Base64RawURL  string `json:"f_base64_raw_url" validate:"base64rawurl"`
	F_BIC           string `json:"f_bic"            validate:"bic"`
	F_BCP47         string `json:"f_bcp47"          validate:"bcp47_language_tag"`
	F_BTCAddr       string `json:"f_btc_addr"       validate:"btc_addr"`
	F_BTCBech32     string `json:"f_btc_bech32"     validate:"btc_addr_bech32"`
	F_CreditCard    string `json:"f_credit_card"    validate:"credit_card"`
	F_MongoID       string `json:"f_mongo_id"       validate:"mongodb"`
	F_MongoConn     string `json:"f_mongo_conn"     validate:"mongodb_connection_string"`
	F_Cron          string `json:"f_cron"           validate:"cron"`
	F_SpiceDB       string `json:"f_spice_db"       validate:"spicedb"`
	F_DateTime      string `json:"f_date_time"      validate:"datetime"`
	F_E164          string `json:"f_e164"           validate:"e164"`
	F_EIN           string `json:"f_ein"            validate:"ein"`
	F_Email         string `json:"f_email"          validate:"email"`
	F_EthAddr       string `json:"f_eth_addr"       validate:"eth_addr"`
	F_EthChecksum   string `json:"f_eth_checksum"   validate:"eth_addr_checksum"`
	F_Hexadecimal   string `json:"f_hexadecimal"    validate:"hexadecimal"`
	F_HexColor      string `json:"f_hex_color"      validate:"hexcolor"`
	F_HSL           string `json:"f_hsl"            validate:"hsl"`
	F_HSLA          string `json:"f_hsla"           validate:"hsla"`
	F_HTML          string `json:"f_html"           validate:"html"`
	F_HTMLEnc       string `json:"f_html_enc"       validate:"html_encoded"`
	F_ISBN          string `json:"f_isbn"           validate:"isbn"`
	F_ISBN10        string `json:"f_isbn10"         validate:"isbn10"`
	F_ISBN13        string `json:"f_isbn13"         validate:"isbn13"`
	F_ISSN          string `json:"f_issn"           validate:"issn"`
	F_ISO3166A2     string `json:"f_iso3166_a2"     validate:"iso3166_1_alpha2"`
	F_ISO3166A2_EU  string `json:"f_iso3166_a2_eu"  validate:"iso3166_1_alpha2_eu"`
	F_ISO3166A3     string `json:"f_iso3166_a3"     validate:"iso3166_1_alpha3"`
	F_ISO3166A3_EU  string `json:"f_iso3166_a3_eu"  validate:"iso3166_1_alpha3_eu"`
	F_ISO3166Num    string `json:"f_iso3166_num"    validate:"iso3166_1_alpha_numeric"`
	F_ISO3166Num_EU string `json:"f_iso3166_num_eu" validate:"iso3166_1_alpha_numeric_eu"`
	F_ISO3166_2     string `json:"f_iso3166_2"      validate:"iso3166_2"`
	F_ISO4217       string `json:"f_iso4217"        validate:"iso4217"`
	F_ISO4217Num    uint16 `json:"f_iso4217_num"    validate:"iso4217_numeric"`
	F_JSON          string `json:"f_json"           validate:"json"`
	F_JWT           string `json:"f_jwt"            validate:"jwt"`
	F_Lat           string `json:"f_lat"            validate:"latitude"`
	F_Lon           string `json:"f_lon"            validate:"longitude"`
	F_Luhn          string `json:"f_luhn"           validate:"luhn_checksum"`
	F_Postcode      string `json:"f_postcode"       validate:"postcode_iso3166_alpha2"`
	F_PostcodeField string `json:"f_postcode_field" validate:"postcode_iso3166_alpha2_field"`
	F_RGB           string `json:"f_rgb"            validate:"rgb"`
	F_RGBA          string `json:"f_rgba"           validate:"rgba"`
	F_SSN           string `json:"f_ssn"            validate:"ssn"`
	F_Timezone      string `json:"f_timezone"       validate:"timezone"`
	F_UUID          string `json:"f_uuid"           validate:"uuid"`
	F_UUID3         string `json:"f_uuid3"          validate:"uuid3"`
	F_UUID3RFC      string `json:"f_uuid3_rfc"      validate:"uuid3_rfc4122"`
	F_UUID4         string `json:"f_uuid4"          validate:"uuid4"`
	F_UUID4RFC      string `json:"f_uuid4_rfc"      validate:"uuid4_rfc4122"`
	F_UUID5         string `json:"f_uuid5"          validate:"uuid5"`
	F_UUID5RFC      string `json:"f_uuid5_rfc"      validate:"uuid5_rfc4122"`
	F_UUIDRFC       string `json:"f_uuid_rfc"       validate:"uuid_rfc4122"`
	F_MD4           string `json:"f_md4"            validate:"md4"`
	F_MD5           string `json:"f_md5"            validate:"md5"`
	F_SHA256        string `json:"f_sha256"         validate:"sha256"`
	F_SHA384        string `json:"f_sha384"         validate:"sha384"`
	F_SHA512        string `json:"f_sha512"         validate:"sha512"`
	F_RIPEMD128     string `json:"f_ripemd128"      validate:"ripemd128"`
	F_RIPEMD160     string `json:"f_ripemd160"      validate:"ripemd160"`
	F_Tiger128      string `json:"f_tiger128"       validate:"tiger128"`
	F_Tiger160      string `json:"f_tiger160"       validate:"tiger160"`
	F_Tiger192      string `json:"f_tiger192"       validate:"tiger192"`
	F_Semver        string `json:"f_semver"         validate:"semver"`
	F_ULID          string `json:"f_ulid"           validate:"ulid"`
	F_CVE           string `json:"f_cve"            validate:"cve"`

	// ---------- Comparisons ----------
	C_Eq   string `json:"c_eq"    validate:"eq=abc"`
	C_EqIC string `json:"c_eq_ic" validate:"eq_ignore_case=abc"`
	C_Gt   int    `json:"c_gt"    validate:"gt=10"`
	C_Gte  int    `json:"c_gte"   validate:"gte=10"`
	C_Lt   int    `json:"c_lt"    validate:"lt=10"`
	C_Lte  int    `json:"c_lte"   validate:"lte=10"`
	C_Ne   string `json:"c_ne"    validate:"ne=abc"`
	C_NeIC string `json:"c_ne_ic" validate:"ne_ignore_case=abc"`

	// ---------- Other ----------
	O_Dir       string `json:"o_dir"        validate:"dir"`
	O_DirPath   string `json:"o_dir_path"   validate:"dirpath"`
	O_File      string `json:"o_file"       validate:"file"`
	O_FilePath  string `json:"o_file_path"  validate:"filepath"`
	O_Image     string `json:"o_image"      validate:"image"`
	O_IsDefault int    `json:"o_is_default" validate:"isdefault"`
	O_Len       string `json:"o_len"        validate:"len=5"`
	O_Max       int    `json:"o_max"        validate:"max=5"`
	O_Min       int    `json:"o_min"        validate:"min=5"`
	O_OneOf     string `json:"o_one_of"     validate:"oneof=a b c"`
	C_OneOfCI   string `json:"c_one_of_ci"  validate:"oneofci=a b c"`

	O_SkipUnless string `json:"o_skip_unless"          validate:"skip_unless=O_Other xyz"`

	// --- Required 系 ---
	O_Required           string `json:"o_required"             validate:"required"`
	O_RequiredIf         string `json:"o_required_if"          validate:"required_if=O_Other xyz"`
	O_RequiredUnless     string `json:"o_required_unless"      validate:"required_unless=O_Other xyz"`
	O_RequiredWith       string `json:"o_required_with"        validate:"required_with=O_Other"`
	O_RequiredWithAll    string `json:"o_required_with_all"    validate:"required_with_all=O_Other O_Another"`
	O_RequiredWithout    string `json:"o_required_without"     validate:"required_without=O_Other"`
	O_RequiredWithoutAll string `json:"o_required_without_all" validate:"required_without_all=O_Other O_Another"`

	// --- Excluded 系 ---
	O_ExcludedIf         string `json:"o_excluded_if"          validate:"excluded_if=O_Other xyz"`
	O_ExcludedUnless     string `json:"o_excluded_unless"      validate:"excluded_unless=O_Other xyz"`
	O_ExcludedWith       string `json:"o_excluded_with"        validate:"excluded_with=O_Other"`
	O_ExcludedWithAll    string `json:"o_excluded_with_all"    validate:"excluded_with_all=O_Other O_Another"`
	O_ExcludedWithout    string `json:"o_excluded_without"     validate:"excluded_without=O_Other"`
	O_ExcludedWithoutAll string `json:"o_excluded_without_all" validate:"excluded_without_all=O_Other O_Another"`

	// unique
	O_UniqueSlice []int `json:"o_unique_slice" validate:"unique"`

	// validateFn
	O_ValidateFn Custom `json:"o_validate_fn"  validate:"validateFn"`

	// 参照用フィールド
	O_Other   string `json:"o_other"`
	O_Another string `json:"o_another"`

	// ---------- Aliases ----------
	A_IsColor     string `json:"a_is_color"     validate:"iscolor"`
	A_CountryCode string `json:"a_country_code" validate:"country_code"`
	A_EUCountry   string `json:"a_eu_country"   validate:"eu_country_code"`
}

// Custom は validateFn テスト用のダミー型
type Custom struct{}

func (Custom) Validate() error { return fmt.Errorf("invalid") }

// 1. 共通: validator 初期化 & validateFn
func newValidator() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	// JSON タグ名を優先的に返す関数を登録
	// → namespace、fieldがjsonタグ型の名前になる
	//    これを登録しないと、namespace、fieldにはGoのフィールド名が使われる = namespaceとstruct_namespaceが一緒、fieldとstruct_fieldが一緒になる
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// `json:"-"` のときは空文字を返し、Field() では飛ばす
		tag := fld.Tag.Get("json")
		if tag == "-" || tag == "" {
			return fld.Name // 代替として Go フィールド名
		}
		// オプション（`,omitempty` など）を除去して最初のトークンだけ
		name := strings.Split(tag, ",")[0]
		return name
	})

	// validateFn エイリアス用
	v.RegisterValidation("validateFn", func(fl validator.FieldLevel) bool {
		if v, ok := fl.Field().Interface().(interface{ Validate() error }); ok {
			return v.Validate() == nil
		}
		return false
	})
	return v
}

// 2. “全部失敗” 用の 3 ケース
func badCaseA() AllTags {
	// - O_Other = "xyz" で required_if / excluded_if を発火
	// - O_Another = "abc" で *_with_all 系も発火
	return AllTags{
		// -------- Field ↔ Field --------
		F_EqCSFieldA: 1, //
		F_EqCSFieldB: 2, //

		F_EqFieldA: "foo", //
		F_EqFieldB: "bar", //

		F_FieldContains:    "abcd", //
		F_FieldContainsSrc: "xyz",  //

		F_FieldExcludes:    "xyz", //
		F_FieldExcludesSrc: "y",   //

		F_GtCSFieldA: 1, //
		F_GtCSFieldB: 2, //

		F_GteCSFieldA: 1, //
		F_GteCSFieldB: 2, //

		F_GteFieldA: 9,  //
		F_GteFieldB: 10, //

		F_GtFieldA: 10, //
		F_GtFieldB: 10, //

		F_LtCSFieldA: 2, //
		F_LtCSFieldB: 1, //

		F_LteCSFieldA: 2, //
		F_LteCSFieldB: 1, //

		F_LteFieldA: 11, //
		F_LteFieldB: 10, //

		F_LtFieldA: 10, //
		F_LtFieldB: 10, //

		F_NeCSFieldA: 1, //
		F_NeCSFieldB: 1, //

		F_NeFieldA: 1, //
		F_NeFieldB: 1, //

		// ---------- Network ----------
		N_CIDR:            "invalid",                        //
		N_CIDRv4:          "invalid",                        //
		N_CIDRv6:          "invalid",                        //
		N_DataURI:         "invalid",                        //
		N_FQDN:            "invalid",                        //
		N_Hostname:        "host name with spaces",          //
		N_HostnamePort:    "invalid",                        //
		N_Port:            0,                                // 範囲外 or 不正ポート
		N_HostnameRFC1123: "-starts.with.dash",              //
		N_DNSLabel:        "Invalid_Label!",                 // uppercase & punctuation
		N_IP:              "999.999.999.999",                //
		N_IP4Addr:         "999.999.999.999",                //
		N_IP6Addr:         "zzzz::zzzz",                     //
		N_IPAddr:          "not-an-ip",                      //
		N_IPv4:            "300.300.300.300",                //
		N_IPv6:            "gggg::gggg",                     //
		N_MAC:             "00:00:00:00:00:zz",              //
		N_TCP4Addr:        "999.999.999.999:99999",          //
		N_TCP6Addr:        "[gggg::gggg]:99999",             //
		N_TCPAddr:         "not-a-tcp-addr",                 //
		N_UDP4Addr:        "999.999.999.999:99999",          //
		N_UDP6Addr:        "[gggg::gggg]:99999",             //
		N_UDPAddr:         "not-a-udp-addr",                 //
		N_UnixAddr:        "/this/path/does/not/exist.sock", //
		N_URI:             "ht!tp://%%%",                    //
		N_URL:             "http:// invalid .com",           //
		N_HTTPURL:         "httpp://example..com",           //
		N_URLEncoded:      "abc%zz",                         // 不正な % エスケープ → url_encoded 失敗
		N_URN:             "urn:invalid:%%%",                //

		// -------- Strings (失敗値が必要なものだけ埋める) --------
		S_Alpha:         "123",    //
		S_Alnum:         "___",    //
		S_AlnumUnicode:  "abc!",   // contains punctuation
		S_AlphaUnicode:  "abc123", // 英数字混在 → alphabetic Unicode でエラー
		S_ASCII:         "©",      // 非 ASCII
		S_Bool:          "maybe",  //
		S_Contains:      "foo",    //
		S_ContainsAny:   "",       //
		S_ContainsRune:  "abc",    //
		S_EndsNotWith:   "abc",    // 禁止 suffix
		S_EndsWith:      "zzz",    //
		S_Excludes:      "xyz",    // 含んでしまう
		S_ExcludesAll:   "b",      //
		S_ExcludesRune:  "😊",      //
		S_Lowercase:     "ABC",    //
		S_Multibyte:     "abc",    // ASCII だけ → multibyte 失敗
		S_Number:        "abc",    //
		S_Numeric:       "abc",    //
		S_PrintASCII:    "é",      //
		S_StartsNotWith: "abcXYZ", // 禁止 prefix
		S_StartsWith:    "XYZ",    //
		S_Uppercase:     "abc",    //

		// ---------- Format ----------
		F_Base32:        "not-base32",                                //
		F_Base64:        "not_base64!!",                              //
		F_Base64URL:     "not-base64url",                             //
		F_Base64RawURL:  "not_base64!!",                              // invalid chars for Base64RawURL
		F_BIC:           "INVALIDBIC",                                //
		F_BCP47:         "invalid_tag",                               //
		F_BTCAddr:       "InvalidBTCAddr",                            //
		F_BTCBech32:     "bc1invalidaddress",                         //
		F_CreditCard:    "1234-5678-9012-3456",                       //
		F_MongoID:       "invalidobjectid",                           //
		F_MongoConn:     "not-a-mongodb://connection",                //
		F_Cron:          "invalid cron",                              // wrong format
		F_SpiceDB:       "bad type",                                  //
		F_DateTime:      "not-a-date",                                // 無効 → datetime 失敗
		F_E164:          "+12345",                                    //
		F_EIN:           "12-3456-789",                               //
		F_Email:         "invalid@",                                  //
		F_EthAddr:       "0xINVALID",                                 //
		F_EthChecksum:   "0x52908400098527886E0F7030069857D2E4169EF", //
		F_Hexadecimal:   "Z123",                                      //
		F_HexColor:      "#GGGGGG",                                   //
		F_HSL:           "hsl(361,100%,50%)",                         //
		F_HSLA:          "hsla(361,100%,50%,1)",                      //
		F_HTML:          "<unclosed tag",                             //
		F_HTMLEnc:       "&invalident;",                              //
		F_ISBN:          "invalidisbn",                               //
		F_ISBN10:        "123456789",                                 //
		F_ISBN13:        "123456789012",                              //
		F_ISSN:          "1234-INVALID",                              //
		F_ISO3166A2:     "XX",                                        //
		F_ISO3166A2_EU:  "XX",                                        //
		F_ISO3166A3:     "XXX",                                       //
		F_ISO3166A3_EU:  "XXX",                                       //
		F_ISO3166Num:    "9999",                                      //
		F_ISO3166Num_EU: "9999",                                      //
		F_ISO3166_2:     "invalid",                                   //
		F_ISO4217:       "123",                                       //
		F_ISO4217Num:    1000,                                        //
		F_JSON:          "{invalid:json}",                            //
		F_JWT:           "header.payload",                            // only one '.' and not Base64
		F_Lat:           "100.0000",                                  //
		F_Lon:           "200.0000",                                  //
		F_Luhn:          "123456789",                                 //
		F_Postcode:      "ABCDE",                                     //
		F_PostcodeField: "12345",                                     //
		F_RGB:           "rgb(256,0,0)",                              //
		F_RGBA:          "rgba(0,256,0,1)",                           //
		F_SSN:           "000-00-0000",                               //
		F_Timezone:      "Invalid/Zone",                              //
		F_UUID:          "invalid-uuid",                              //
		F_UUID3:         "1234",                                      //
		F_UUID3RFC:      "invalid-uuid3",                             //
		F_UUID4:         "invalid-uuid4",                             //
		F_UUID4RFC:      "invalid-uuid4",                             //
		F_UUID5:         "invalid-uuid5",                             //
		F_UUID5RFC:      "invalid-uuid5",                             //
		F_UUIDRFC:       "invalid-uuid",                              //
		F_MD4:           "invalidmd4",                                //
		F_MD5:           "invalidmd5",                                //
		F_SHA256:        "invalidsha256",                             //
		F_SHA384:        "invalidsha384",                             //
		F_SHA512:        "invalidsha512",                             //
		F_RIPEMD128:     "invalidripemd128",                          //
		F_RIPEMD160:     "invalidripemd160",                          //
		F_Tiger128:      "invalidtiger128",                           //
		F_Tiger160:      "invalidtiger160",                           //
		F_Tiger192:      "invalidtiger192",                           //
		F_Semver:        "v1.x",                                      //
		F_ULID:          "01ARZ3NDEKTSV4RRFFQ69G5",                   // incorrect length
		F_CVE:           "CVE-0000-0000",                             //

		// -------- Comparisons --------
		C_Eq:   "xyz", //
		C_EqIC: "DEF", //
		C_Gt:   0,     //
		C_Gte:  9,     //
		C_Lt:   10,    //
		C_Lte:  11,    //
		C_Ne:   "abc", //
		C_NeIC: "ABC", //

		// -------- Other pivots --------
		O_Dir:       "/path/does/not/exist",      //
		O_DirPath:   "/invalid//dir//path",       //
		O_File:      "/path/to/nonexistent.file", //
		O_FilePath:  "invalid\x00path",           // embed a NUL (0x00) to force os.Stat → EINVAL → isFilePath returns false
		O_Image:     "not-an-image-content",      //
		O_IsDefault: 1,                           // default ではない値を強制
		O_Len:       "too_long",                  // len=5 のはずが文字数超過
		O_Max:       6,                           // max=5 を超える → 失敗
		O_Min:       4,                           // min=5 より小さい → 失敗
		O_OneOf:     "invalid_choice",            //
		C_OneOfCI:   "InvalidChoice",             // case-insensitive 一致しない

		O_SkipUnless: "", // skip_unless で検証実行 → skip_unless でエラー

		O_Required:        "", // 空のまま → required でエラー
		O_RequiredIf:      "", // 条件付きなのに値がない → required_if でエラー
		O_RequiredUnless:  "", // unless 条件なのに空 → required_unless でエラー
		O_RequiredWith:    "", // with 条件を満たすのに空 → required_with でエラー
		O_RequiredWithAll: "", // with_all 条件を満たすのに空 → required_with_all でエラー
		// O_RequiredWithout:     "", // without 条件を満たさないのに値あり → required_without でエラー
		// O_RequiredWithoutAll:  "", // without_all 条件を満たさないのに値あり → required_without_all でエラー
		// required_without / excluded_without 系は別ケースで発火させる

		O_ExcludedIf:      "foo", //
		O_ExcludedUnless:  "foo", // unless 条件を満たさないのに値あり → excluded_unless でエラー
		O_ExcludedWith:    "foo", // with 条件を満たすのに値あり → excluded_with でエラー
		O_ExcludedWithAll: "foo", // with_all 条件を満たすのに値あり → excluded_with_all でエラー
		// O_ExcludedWithout:    "", // without 条件を満たさないのに空 → excluded_without でエラー
		// O_ExcludedWithoutAll: "", // without_all 条件を満たさないのに空 → excluded_without_all でエラー
		// required_without / excluded_without 系は別ケースで発火させる

		// unique
		O_UniqueSlice: []int{1, 1}, //

		// validateFn
		O_ValidateFn: Custom{}, //

		// 参照用フィールド
		O_Other:   "xyz", //
		O_Another: "abc", //

		// --- Alias (A_*) 系にも“不正値”を追加 ---
		A_IsColor:     "not-a-color", // hex/rgb/hsl いずれにもマッチしない
		A_CountryCode: "ZZZ",         // ISO 3166-1 alpha-2 ではない
		A_EUCountry:   "XYZ",         // EU country code ではない
	}
}

func badCaseB() AllTags {
	// O_Other/O_Another を空にして *_without 系を発火
	b := badCaseA()
	b.O_Other = ""
	b.O_Another = ""

	b.O_RequiredWithout = ""
	b.O_RequiredWithoutAll = ""
	b.O_ExcludedWithout = "foo"
	b.O_ExcludedWithoutAll = "foo"
	return b
}

func badCaseC() AllTags {
	// O_Other を "zzz" にして required_unless / excluded_unless を発火
	b := badCaseA()
	b.O_Other = "zzz"

	b.O_RequiredUnless = ""
	b.O_ExcludedUnless = "foo"
	return b
}

func Test_AllValidatorTags_AllTagsDefined(t *testing.T) {

	if len(ruleSet) != len(bakedTags) {
		t.Errorf("a tag or some tags are not defined. baked=%d defined=%d", len(bakedTags), len(ruleSet))
	}

BakedTagLoop:
	for _, bakedTag := range bakedTags {
		for tag := range ruleSet {
			if bakedTag == tag {
				continue BakedTagLoop
			}
		}
		t.Errorf("tag '%s' is not defined.", bakedTag)
	}
}

func Test_AllValidatorTags_AllTagsTestCovered(t *testing.T) {

	rt := reflect.TypeOf(AllTags{})
	numFields := rt.NumField()

BakedTagLoop:
	for _, bakedTag := range bakedTags {
		for i := range numFields {
			validateTagValuesStr := rt.Field(i).Tag.Get("validate")
			validateTagValues := strings.Split(validateTagValuesStr, ",")
			for _, validateTagValue := range validateTagValues {
				if strings.Contains(validateTagValue, "=") {
					validateTagValue = strings.Split(validateTagValue, "=")[0]
				}

				if bakedTag == validateTagValue {
					continue BakedTagLoop
				}
			}
		}
		t.Errorf("tag '%s' is not defined.", bakedTag)
	}

	for key := range ruleSet {
		if !slices.Contains(bakedTags, key) {
			t.Errorf("tag '%s' is not defined.", key)
		}
	}
}

func Test_AllValidatorTags_Empty(t *testing.T) {

	validate := newValidator()
	v := AllTags{}

	err := validate.Struct(&v) // ゼロ値で全部 NG
	if err == nil {
		t.Fatal("error: expected errors")
	}

	var verr validator.ValidationErrors
	if !errors.As(err, &verr) {
		t.Fatalf("error: unexpected error: %v", err)
	}

	seen := make(map[string]struct{})
	for _, fe := range verr {
		msg := Message(reflect.ValueOf(v), fe) // ruleSet を呼び出してみる
		t.Logf("msg: %s", msg)

		seen[fe.Tag()] = struct{}{}
	}

	// 取りこぼしチェック
	for tag := range ruleSet {
		if _, ok := seen[tag]; !ok {
			// エラーになっていない場合

			// バグ等でバリデーションが機能していない
			if slices.Contains(skipCoverage, tag) {
				t.Logf("log: tag %q did NOT fail, but tag %q is in skipCoverage (known-always-pass validator)", tag, tag)
				continue
			}

			// ゼロ値は PASS が仕様
			if slices.Contains(skipEmpty, tag) {
				t.Logf("log: tag %q did NOT fail, but tag %q is in empty-value test (allowed-empty validator)", tag, tag)
				continue
			}

			t.Errorf("error: tag %q did NOT fail; adjust test value for full coverage", tag)
		}
	}

	if t.Failed() {
		t.Logf("error: tags NOT covered: %d", len(ruleSet)-len(seen))
	}
}

func Test_AllValidatorTags_UnhappyPath(t *testing.T) {

	validate := newValidator()

	cases := []AllTags{badCaseA(), badCaseB(), badCaseC()}

	seen := make(map[string]struct{})

	for index, tc := range cases {
		if err := validate.Struct(tc); err != nil {
			var verr validator.ValidationErrors
			if !errors.As(err, &verr) {
				t.Fatalf("error: [%02d] unexpected error: %v", index, err)
			}
			for _, fe := range verr {
				msg := Message(reflect.ValueOf(tc), fe) // ruleSet を呼び出してみる
				t.Logf("msg: %s", msg)

				seen[fe.Tag()] = struct{}{}
			}
		}

		t.Logf("\n") // 空白行
	}

	// 取りこぼしチェック
	for tag := range ruleSet {
		if _, ok := seen[tag]; !ok {
			if slices.Contains(skipCoverage, tag) {
				t.Logf("log: tag %q did NOT fail, but tag %q is in skipCoverage (known-always-pass validator)", tag, tag)
				continue
			}

			t.Errorf("error: tag %q did NOT fail; adjust test value for full coverage", tag)
		}
	}

	if t.Failed() {
		t.Logf("error: tags NOT covered: %d", len(ruleSet)-len(seen))
	}
}

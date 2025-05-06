# playgrounds/go-playground_validator_with_cerrors

* https://github.com/go-playground/validator

## main.go

```bash
$ go run -v main.go 
command-line-arguments
namespace: User.age
field: age
struct_namespace: User.Age
struct_field: Age
tag: lte
actual_tag: lte
kind: uint8
type: uint8
value: 135
param: 130

namespace: User.favorite_color
field: favorite_color
struct_namespace: User.FavouriteColor
struct_field: FavouriteColor
tag: iscolor
actual_tag: hexcolor|rgb|rgba|hsl|hsla
kind: string
type: string
value: #000-
param: 

namespace: User.addresses[0].city
field: city
struct_namespace: User.Addresses[0].City
struct_field: City
tag: required
actual_tag: required
kind: string
type: string
value: 
param: 

msg: age must be at most 130, but age is 135
msg: favorite_color must be a valid color (hex, rgb[a], hsl[a]), but favorite_color is '#000-'
msg: addresses[0].city must set, but addresses[0].city is ''

{"time":"2025-05-06T09:44:05.443014635Z","level":"ERROR","source":{"function":"main.validateStruct","file":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","line":130},"msg":"error message","err":{"code":"UNKNOWN_ERROR","detail":"an unknown error occurred","cause":"Key: 'User.age' Error:Field validation for 'age' failed on the 'lte' tag\nKey: 'User.favorite_color' Error:Field validation for 'favorite_color' failed on the 'iscolor' tag\nKey: 'User.addresses[0].city' Error:Field validation for 'city' failed on the 'required' tag","stacktrace":[{"function":"validateStruct","module":"main","filename":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","abs_path":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","lineno":122,"in_app":true},{"function":"main","module":"main","filename":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","abs_path":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","lineno":57,"in_app":true},{"function":"main","module":"runtime","filename":"runtime/proc.go","abs_path":"/usr/local/go/src/runtime/proc.go","lineno":283,"in_app":true},{"function":"goexit","module":"runtime","filename":"runtime/asm_arm64.s","abs_path":"/usr/local/go/src/runtime/asm_arm64.s","lineno":1223,"in_app":true}],"stacktrace_order":"newest_first","messages":[{"function":"validateStruct","module":"main","filename":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","abs_path":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","lineno":122,"in_app":false,"message":"age must be at most 130, but age is 135"},{"function":"validateStruct","module":"main","filename":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","abs_path":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","lineno":122,"in_app":false,"message":"favorite_color must be a valid color (hex, rgb[a], hsl[a]), but favorite_color is '#000-'"},{"function":"validateStruct","module":"main","filename":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","abs_path":"/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go","lineno":122,"in_app":false,"message":"addresses[0].city must set, but addresses[0].city is ''"}]}}
Key: '' Error:Field validation for '' failed on the 'email' tag
```

```json
{
  "time": "2025-05-06T09:44:05.443014635Z",
  "level": "ERROR",
  "source": {
    "function": "main.validateStruct",
    "file": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
    "line": 130
  },
  "msg": "error message",
  "err": {
    "code": "UNKNOWN_ERROR",
    "detail": "an unknown error occurred",
    "cause": "Key: 'User.age' Error:Field validation for 'age' failed on the 'lte' tag\nKey: 'User.favorite_color' Error:Field validation for 'favorite_color' failed on the 'iscolor' tag\nKey: 'User.addresses[0].city' Error:Field validation for 'city' failed on the 'required' tag",
    "stacktrace": [
      {
        "function": "validateStruct",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "lineno": 122,
        "in_app": true
      },
      {
        "function": "main",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "lineno": 57,
        "in_app": true
      },
      {
        "function": "main",
        "module": "runtime",
        "filename": "runtime/proc.go",
        "abs_path": "/usr/local/go/src/runtime/proc.go",
        "lineno": 283,
        "in_app": true
      },
      {
        "function": "goexit",
        "module": "runtime",
        "filename": "runtime/asm_arm64.s",
        "abs_path": "/usr/local/go/src/runtime/asm_arm64.s",
        "lineno": 1223,
        "in_app": true
      }
    ],
    "stacktrace_order": "newest_first",
    "messages": [
      {
        "function": "validateStruct",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "lineno": 122,
        "in_app": false,
        "message": "age must be at most 130, but age is 135"
      },
      {
        "function": "validateStruct",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "lineno": 122,
        "in_app": false,
        "message": "favorite_color must be a valid color (hex, rgb[a], hsl[a]), but favorite_color is '#000-'"
      },
      {
        "function": "validateStruct",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/go-playground_validator_with_cerrors/main.go",
        "lineno": 122,
        "in_app": false,
        "message": "addresses[0].city must set, but addresses[0].city is ''"
      }
    ]
  }
}
```

## validatorx: unit test

```bash
$ go test -v validatorx/*.go
=== RUN   Test_AllValidatorTags_AllTagsDefined
--- PASS: Test_AllValidatorTags_AllTagsDefined (0.00s)
=== RUN   Test_AllValidatorTags_AllTagsTestCovered
--- PASS: Test_AllValidatorTags_AllTagsTestCovered (0.00s)
=== RUN   Test_AllValidatorTags_Empty
    messages_test.go:832: msg: f_field_excludes must not contain f_field_excludes_src (value: ''), but f_field_excludes is ''
    messages_test.go:832: msg: f_gt_cs_field_a must be greater than f_gt_cs_field_b (value: 0), but f_gt_cs_field_a is 0
    messages_test.go:832: msg: f_gt_field_a must be greater than f_gt_field_b (value: 0), but f_gt_field_a is 0
    messages_test.go:832: msg: f_lt_cs_field_a must be less than f_lt_cs_field_b (value: 0), but f_lt_cs_field_a is 0
    messages_test.go:832: msg: f_lt_field_a must be less than f_lt_field_b (value: 0), but f_lt_field_a is 0
    messages_test.go:832: msg: f_ne_cs_field_a must be different from f_ne_cs_field_b (value: 0), but f_ne_cs_field_a is 0
    messages_test.go:832: msg: f_ne_field_a must be different from f_ne_field_b (value: 0), but f_ne_field_a is 0
    messages_test.go:832: msg: n_cidr must be a valid CIDR, but n_cidr is ''
    messages_test.go:832: msg: n_cidr_v4 must be a valid CIDRv4, but n_cidr_v4 is ''
    messages_test.go:832: msg: n_cidr_v6 must be a valid CIDRv6, but n_cidr_v6 is ''
    messages_test.go:832: msg: n_data_uri must be a valid Data URI, but n_data_uri is ''
    messages_test.go:832: msg: n_fqdn must be a valid FQDN, but n_fqdn is ''
    messages_test.go:832: msg: n_hostname must be a valid hostname, but n_hostname is ''
    messages_test.go:832: msg: n_hostname_port must be a valid host:port, but n_hostname_port is ''
    messages_test.go:832: msg: n_port must be a valid TCP/UDP port, but n_port is 0
    messages_test.go:832: msg: n_hostname_rfc1123 must be a valid RFC 1123 hostname, but n_hostname_rfc1123 is ''
    messages_test.go:832: msg: n_dns_label must be a valid RFC1035 DNS label, but n_dns_label is ''
    messages_test.go:832: msg: n_ip must be a valid IP address, but n_ip is ''
    messages_test.go:832: msg: n_ip4_addr must be a valid IPv4 address, but n_ip4_addr is ''
    messages_test.go:832: msg: n_ip6_addr must be a valid IPv6 address, but n_ip6_addr is ''
    messages_test.go:832: msg: n_ip_addr must be a valid IP address, but n_ip_addr is ''
    messages_test.go:832: msg: n_ip_v4 must be a valid IPv4 address, but n_ip_v4 is ''
    messages_test.go:832: msg: n_ip_v6 must be a valid IPv6 address, but n_ip_v6 is ''
    messages_test.go:832: msg: n_mac must be a valid MAC address, but n_mac is ''
    messages_test.go:832: msg: n_tcp4_addr must be a valid TCPv4 address, but n_tcp4_addr is ''
    messages_test.go:832: msg: n_tcp6_addr must be a valid TCPv6 address, but n_tcp6_addr is ''
    messages_test.go:832: msg: n_tcp_addr must be a valid TCP address, but n_tcp_addr is ''
    messages_test.go:832: msg: n_udp4_addr must be a valid UDPv4 address, but n_udp4_addr is ''
    messages_test.go:832: msg: n_udp6_addr must be a valid UDPv6 address, but n_udp6_addr is ''
    messages_test.go:832: msg: n_udp_addr must be a valid UDP address, but n_udp_addr is ''
    messages_test.go:832: msg: n_uri must be a valid URI, but n_uri is ''
    messages_test.go:832: msg: n_url must be a valid URL, but n_url is ''
    messages_test.go:832: msg: n_http_url must be a valid HTTP URL, but n_http_url is ''
    messages_test.go:832: msg: n_urn must be a valid RFC 2141 URN, but n_urn is ''
    messages_test.go:832: msg: s_alpha must alphabetic characters, but s_alpha is ''
    messages_test.go:832: msg: s_alnum must alphanumeric characters, but s_alnum is ''
    messages_test.go:832: msg: s_alnum_unicode must alphanumeric Unicode characters, but s_alnum_unicode is ''
    messages_test.go:832: msg: s_alpha_unicode must alphabetic Unicode characters, but s_alpha_unicode is ''
    messages_test.go:832: msg: s_bool must be a boolean, but s_bool is ''
    messages_test.go:832: msg: s_contains must be contain 'xyz', but s_contains is ''
    messages_test.go:832: msg: s_contains_any must be contain any of 'xyz', but s_contains_any is ''
    messages_test.go:832: msg: s_contains_rune must be contain rune '😊', but s_contains_rune is ''
    messages_test.go:832: msg: s_ends_with must be end with 'abc', but s_ends_with is ''
    messages_test.go:832: msg: s_lowercase must lower-case, but s_lowercase is ''
    messages_test.go:832: msg: s_number must be a number, but s_number is ''
    messages_test.go:832: msg: s_numeric must be numeric, but s_numeric is ''
    messages_test.go:832: msg: s_starts_with must be start with 'abc', but s_starts_with is ''
    messages_test.go:832: msg: s_uppercase must upper-case, but s_uppercase is ''
    messages_test.go:832: msg: f_base32 must be a valid Base32, but f_base32 is ''
    messages_test.go:832: msg: f_base64 must be a valid Base64, but f_base64 is ''
    messages_test.go:832: msg: f_base64_url must be a valid Base64URL, but f_base64_url is ''
    messages_test.go:832: msg: f_base64_raw_url must be a valid Base64RawURL, but f_base64_raw_url is ''
    messages_test.go:832: msg: f_bic must be a valid BIC, but f_bic is ''
    messages_test.go:832: msg: f_bcp47 must be a valid BCP-47 language tag, but f_bcp47 is ''
    messages_test.go:832: msg: f_btc_addr must be a valid Bitcoin address, but f_btc_addr is ''
    messages_test.go:832: msg: f_btc_bech32 must be a valid Bech32 Bitcoin address, but f_btc_bech32 is ''
    messages_test.go:832: msg: f_credit_card must be a valid credit-card number, but f_credit_card is ''
    messages_test.go:832: msg: f_mongo_id must be a valid MongoDB ObjectID, but f_mongo_id is ''
    messages_test.go:832: msg: f_mongo_conn must be a valid MongoDB connection string, but f_mongo_conn is ''
    messages_test.go:832: msg: f_cron must be a valid cron expression, but f_cron is ''
    messages_test.go:832: msg: f_spice_db must be a valid SpiceDB identifier, but f_spice_db is ''
    messages_test.go:832: msg: f_e164 must be a valid E.164 phone number, but f_e164 is ''
    messages_test.go:832: msg: f_ein must be a valid EIN, but f_ein is ''
    messages_test.go:832: msg: f_email must be a valid email, but f_email is ''
    messages_test.go:832: msg: f_eth_addr must be a valid Ethereum address, but f_eth_addr is ''
    messages_test.go:832: msg: f_eth_checksum must be a checksummed Ethereum address, but f_eth_checksum is ''
    messages_test.go:832: msg: f_hexadecimal must be hexadecimal, but f_hexadecimal is ''
    messages_test.go:832: msg: f_hex_color must be a valid hex color, but f_hex_color is ''
    messages_test.go:832: msg: f_hsl must be a valid HSL color, but f_hsl is ''
    messages_test.go:832: msg: f_hsla must be a valid HSLA color, but f_hsla is ''
    messages_test.go:832: msg: f_html must HTML tags, but f_html is ''
    messages_test.go:832: msg: f_html_enc must HTML-encoded text, but f_html_enc is ''
    messages_test.go:832: msg: f_isbn must be a valid ISBN, but f_isbn is ''
    messages_test.go:832: msg: f_isbn10 must be a valid ISBN-10, but f_isbn10 is ''
    messages_test.go:832: msg: f_isbn13 must be a valid ISBN-13, but f_isbn13 is ''
    messages_test.go:832: msg: f_issn must be a valid ISSN, but f_issn is ''
    messages_test.go:832: msg: f_iso3166_a2 must be a valid ISO-3166-1 alpha-2 code, but f_iso3166_a2 is ''
    messages_test.go:832: msg: f_iso3166_a2_eu must be a valid ISO-3166-1 alpha-2 EU code, but f_iso3166_a2_eu is ''
    messages_test.go:832: msg: f_iso3166_a3 must be a valid ISO-3166-1 alpha-3 code, but f_iso3166_a3 is ''
    messages_test.go:832: msg: f_iso3166_a3_eu must be a valid ISO-3166-1 alpha-3 EU code, but f_iso3166_a3_eu is ''
    messages_test.go:832: msg: f_iso3166_num must be a valid ISO-3166-1 numeric code, but f_iso3166_num is ''
    messages_test.go:832: msg: f_iso3166_num_eu must be a valid ISO-3166-1 numeric EU code, but f_iso3166_num_eu is ''
    messages_test.go:832: msg: f_iso3166_2 must be a valid ISO-3166-2 code, but f_iso3166_2 is ''
    messages_test.go:832: msg: f_iso4217 must be a valid ISO-4217 currency code, but f_iso4217 is ''
    messages_test.go:832: msg: f_iso4217_num must be a valid ISO-4217 numeric currency code, but f_iso4217_num is 0
    messages_test.go:832: msg: f_json must be valid JSON, but f_json is ''
    messages_test.go:832: msg: f_jwt must be a valid JWT, but f_jwt is ''
    messages_test.go:832: msg: f_lat must be a valid latitude, but f_lat is ''
    messages_test.go:832: msg: f_lon must be a valid longitude, but f_lon is ''
    messages_test.go:832: msg: f_luhn must be a valid Luhn checksum, but f_luhn is ''
    messages_test.go:832: msg: f_postcode must be a valid postcode, but f_postcode is ''
    messages_test.go:832: msg: f_postcode_field must be a valid postcode, but f_postcode_field is ''
    messages_test.go:832: msg: f_rgb must be a valid RGB color, but f_rgb is ''
    messages_test.go:832: msg: f_rgba must be a valid RGBA color, but f_rgba is ''
    messages_test.go:832: msg: f_ssn must be a valid SSN, but f_ssn is ''
    messages_test.go:832: msg: f_timezone must be a valid timezone, but f_timezone is ''
    messages_test.go:832: msg: f_uuid must be a valid UUID, but f_uuid is ''
    messages_test.go:832: msg: f_uuid3 must be a valid UUID v3, but f_uuid3 is ''
    messages_test.go:832: msg: f_uuid3_rfc must be a valid UUID v3 (RFC4122), but f_uuid3_rfc is ''
    messages_test.go:832: msg: f_uuid4 must be a valid UUID v4, but f_uuid4 is ''
    messages_test.go:832: msg: f_uuid4_rfc must be a valid UUID v4 (RFC4122), but f_uuid4_rfc is ''
    messages_test.go:832: msg: f_uuid5 must be a valid UUID v5, but f_uuid5 is ''
    messages_test.go:832: msg: f_uuid5_rfc must be a valid UUID v5 (RFC4122), but f_uuid5_rfc is ''
    messages_test.go:832: msg: f_uuid_rfc must be a valid UUID (RFC4122), but f_uuid_rfc is ''
    messages_test.go:832: msg: f_md4 must be a valid MD4 hash, but f_md4 is ''
    messages_test.go:832: msg: f_md5 must be a valid MD5 hash, but f_md5 is ''
    messages_test.go:832: msg: f_sha256 must be a valid SHA-256 hash, but f_sha256 is ''
    messages_test.go:832: msg: f_sha384 must be a valid SHA-384 hash, but f_sha384 is ''
    messages_test.go:832: msg: f_sha512 must be a valid SHA-512 hash, but f_sha512 is ''
    messages_test.go:832: msg: f_ripemd128 must be a valid RIPEMD-128 hash, but f_ripemd128 is ''
    messages_test.go:832: msg: f_ripemd160 must be a valid RIPEMD-160 hash, but f_ripemd160 is ''
    messages_test.go:832: msg: f_tiger128 must be a valid TIGER-128 hash, but f_tiger128 is ''
    messages_test.go:832: msg: f_tiger160 must be a valid TIGER-160 hash, but f_tiger160 is ''
    messages_test.go:832: msg: f_tiger192 must be a valid TIGER-192 hash, but f_tiger192 is ''
    messages_test.go:832: msg: f_semver must be a valid semantic version, but f_semver is ''
    messages_test.go:832: msg: f_ulid must be a valid ULID, but f_ulid is ''
    messages_test.go:832: msg: f_cve must be a valid CVE identifier, but f_cve is ''
    messages_test.go:832: msg: c_eq must be  'abc', but c_eq is ''
    messages_test.go:832: msg: c_eq_ic must be  'abc', but c_eq_ic is ''
    messages_test.go:832: msg: c_gt must be greater than 10, but c_gt is 0
    messages_test.go:832: msg: c_gte must be at least 10, but c_gte is 0
    messages_test.go:832: msg: o_dir must be an existing directory, but o_dir is ''
    messages_test.go:832: msg: o_dir_path must be a valid directory path, but o_dir_path is ''
    messages_test.go:832: msg: o_file must be an existing file, but o_file is ''
    messages_test.go:832: msg: o_file_path must be a valid file path, but o_file_path is ''
    messages_test.go:832: msg: o_image must be a valid image, but o_image is ''
    messages_test.go:832: msg: o_len must be length 5, but o_len is ''
    messages_test.go:832: msg: o_min must be at least 5, but o_min is 0
    messages_test.go:832: msg: o_one_of must be one of 'a b c', but o_one_of is ''
    messages_test.go:832: msg: c_one_of_ci must be one of (case-insensitive) 'a b c', but c_one_of_ci is ''
    messages_test.go:832: msg: o_required must set, but o_required is ''
    messages_test.go:832: msg: o_required_unless must be set unless O_Other xyz, but o_required_unless is ''
    messages_test.go:832: msg: o_required_without must be set without O_Other, but o_required_without is ''
    messages_test.go:832: msg: o_required_without_all must be set without all O_Other O_Another, but o_required_without_all is ''
    messages_test.go:832: msg: o_validate_fn must valid according to Validate(), but o_validate_fn is {}
    messages_test.go:832: msg: a_is_color must be a valid color (hex, rgb[a], hsl[a]), but a_is_color is ''
    messages_test.go:832: msg: a_country_code must be a valid country code, but a_country_code is ''
    messages_test.go:832: msg: a_eu_country must be a valid EU country code, but a_eu_country is ''
    messages_test.go:850: log: tag "required_if" did NOT fail, but tag "required_if" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "max" did NOT fail, but tag "max" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "ltecsfield" did NOT fail, but tag "ltecsfield" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excludesall" did NOT fail, but tag "excludesall" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "printascii" did NOT fail, but tag "printascii" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "skip_unless" did NOT fail, but tag "skip_unless" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "gtefield" did NOT fail, but tag "gtefield" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "multibyte" did NOT fail, but tag "multibyte" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "lte" did NOT fail, but tag "lte" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excluded_with" did NOT fail, but tag "excluded_with" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excluded_with_all" did NOT fail, but tag "excluded_with_all" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excluded_without_all" did NOT fail, but tag "excluded_without_all" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "eqfield" did NOT fail, but tag "eqfield" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "gtecsfield" did NOT fail, but tag "gtecsfield" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "required_with" did NOT fail, but tag "required_with" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "endsnotwith" did NOT fail, but tag "endsnotwith" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "isdefault" did NOT fail, but tag "isdefault" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "lt" did NOT fail, but tag "lt" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "required_with_all" did NOT fail, but tag "required_with_all" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excludes" did NOT fail, but tag "excludes" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excluded_if" did NOT fail, but tag "excluded_if" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excluded_unless" did NOT fail, but tag "excluded_unless" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "ascii" did NOT fail, but tag "ascii" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "ltefield" did NOT fail, but tag "ltefield" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excluded_without" did NOT fail, but tag "excluded_without" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "datetime" did NOT fail, but tag "datetime" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "excludesrune" did NOT fail, but tag "excludesrune" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "fieldcontains" did NOT fail, but tag "fieldcontains" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "startsnotwith" did NOT fail, but tag "startsnotwith" is in empty-value test (allowed-empty validator)
    messages_test.go:844: log: tag "unix_addr" did NOT fail, but tag "unix_addr" is in skipCoverage (known-always-pass validator)
    messages_test.go:850: log: tag "url_encoded" did NOT fail, but tag "url_encoded" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "unique" did NOT fail, but tag "unique" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "ne" did NOT fail, but tag "ne" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "ne_ignore_case" did NOT fail, but tag "ne_ignore_case" is in empty-value test (allowed-empty validator)
    messages_test.go:850: log: tag "eqcsfield" did NOT fail, but tag "eqcsfield" is in empty-value test (allowed-empty validator)
--- PASS: Test_AllValidatorTags_Empty (0.00s)
=== RUN   Test_AllValidatorTags_UnhappyPath
    messages_test.go:879: msg: f_eq_cs_field_a must be equal to f_eq_cs_field_b (value: 2), but f_eq_cs_field_a is 1
    messages_test.go:879: msg: f_eq_field_a must be equal to f_eq_field_b (value: 'bar'), but f_eq_field_a is 'foo'
    messages_test.go:879: msg: f_field_contains must contain f_field_contains_src (value: 'xyz'), but f_field_contains is 'abcd'
    messages_test.go:879: msg: f_field_excludes must not contain f_field_excludes_src (value: 'y'), but f_field_excludes is 'xyz'
    messages_test.go:879: msg: f_gt_cs_field_a must be greater than f_gt_cs_field_b (value: 2), but f_gt_cs_field_a is 1
    messages_test.go:879: msg: f_gte_cs_field_a must be greater than or equal to f_gte_cs_field_b (value: 2), but f_gte_cs_field_a is 1
    messages_test.go:879: msg: f_gte_field_a must be greater than or equal to f_gte_field_b (value: 10), but f_gte_field_a is 9
    messages_test.go:879: msg: f_gt_field_a must be greater than f_gt_field_b (value: 10), but f_gt_field_a is 10
    messages_test.go:879: msg: f_lt_cs_field_a must be less than f_lt_cs_field_b (value: 1), but f_lt_cs_field_a is 2
    messages_test.go:879: msg: f_lte_cs_field_a must be less than or equal to f_lte_cs_field_b (value: 1), but f_lte_cs_field_a is 2
    messages_test.go:879: msg: f_lte_field_a must be less than or equal to f_lte_field_b (value: 10), but f_lte_field_a is 11
    messages_test.go:879: msg: f_lt_field_a must be less than f_lt_field_b (value: 10), but f_lt_field_a is 10
    messages_test.go:879: msg: f_ne_cs_field_a must be different from f_ne_cs_field_b (value: 1), but f_ne_cs_field_a is 1
    messages_test.go:879: msg: f_ne_field_a must be different from f_ne_field_b (value: 1), but f_ne_field_a is 1
    messages_test.go:879: msg: n_cidr must be a valid CIDR, but n_cidr is 'invalid'
    messages_test.go:879: msg: n_cidr_v4 must be a valid CIDRv4, but n_cidr_v4 is 'invalid'
    messages_test.go:879: msg: n_cidr_v6 must be a valid CIDRv6, but n_cidr_v6 is 'invalid'
    messages_test.go:879: msg: n_data_uri must be a valid Data URI, but n_data_uri is 'invalid'
    messages_test.go:879: msg: n_fqdn must be a valid FQDN, but n_fqdn is 'invalid'
    messages_test.go:879: msg: n_hostname must be a valid hostname, but n_hostname is 'host name with spaces'
    messages_test.go:879: msg: n_hostname_port must be a valid host:port, but n_hostname_port is 'invalid'
    messages_test.go:879: msg: n_port must be a valid TCP/UDP port, but n_port is 0
    messages_test.go:879: msg: n_hostname_rfc1123 must be a valid RFC 1123 hostname, but n_hostname_rfc1123 is '-starts.with.dash'
    messages_test.go:879: msg: n_dns_label must be a valid RFC1035 DNS label, but n_dns_label is 'Invalid_Label!'
    messages_test.go:879: msg: n_ip must be a valid IP address, but n_ip is '999.999.999.999'
    messages_test.go:879: msg: n_ip4_addr must be a valid IPv4 address, but n_ip4_addr is '999.999.999.999'
    messages_test.go:879: msg: n_ip6_addr must be a valid IPv6 address, but n_ip6_addr is 'zzzz::zzzz'
    messages_test.go:879: msg: n_ip_addr must be a valid IP address, but n_ip_addr is 'not-an-ip'
    messages_test.go:879: msg: n_ip_v4 must be a valid IPv4 address, but n_ip_v4 is '300.300.300.300'
    messages_test.go:879: msg: n_ip_v6 must be a valid IPv6 address, but n_ip_v6 is 'gggg::gggg'
    messages_test.go:879: msg: n_mac must be a valid MAC address, but n_mac is '00:00:00:00:00:zz'
    messages_test.go:879: msg: n_tcp4_addr must be a valid TCPv4 address, but n_tcp4_addr is '999.999.999.999:99999'
    messages_test.go:879: msg: n_tcp6_addr must be a valid TCPv6 address, but n_tcp6_addr is '[gggg::gggg]:99999'
    messages_test.go:879: msg: n_tcp_addr must be a valid TCP address, but n_tcp_addr is 'not-a-tcp-addr'
    messages_test.go:879: msg: n_udp4_addr must be a valid UDPv4 address, but n_udp4_addr is '999.999.999.999:99999'
    messages_test.go:879: msg: n_udp6_addr must be a valid UDPv6 address, but n_udp6_addr is '[gggg::gggg]:99999'
    messages_test.go:879: msg: n_udp_addr must be a valid UDP address, but n_udp_addr is 'not-a-udp-addr'
    messages_test.go:879: msg: n_uri must be a valid URI, but n_uri is 'ht!tp://%%%'
    messages_test.go:879: msg: n_url must be a valid URL, but n_url is 'http:// invalid .com'
    messages_test.go:879: msg: n_http_url must be a valid HTTP URL, but n_http_url is 'httpp://example..com'
    messages_test.go:879: msg: n_url_encoded must URL-encoded data, but n_url_encoded is 'abc%zz'
    messages_test.go:879: msg: n_urn must be a valid RFC 2141 URN, but n_urn is 'urn:invalid:%%%'
    messages_test.go:879: msg: s_alpha must alphabetic characters, but s_alpha is '123'
    messages_test.go:879: msg: s_alnum must alphanumeric characters, but s_alnum is '___'
    messages_test.go:879: msg: s_alnum_unicode must alphanumeric Unicode characters, but s_alnum_unicode is 'abc!'
    messages_test.go:879: msg: s_alpha_unicode must alphabetic Unicode characters, but s_alpha_unicode is 'abc123'
    messages_test.go:879: msg: s_ascii must ASCII characters, but s_ascii is '©'
    messages_test.go:879: msg: s_bool must be a boolean, but s_bool is 'maybe'
    messages_test.go:879: msg: s_contains must be contain 'xyz', but s_contains is 'foo'
    messages_test.go:879: msg: s_contains_any must be contain any of 'xyz', but s_contains_any is ''
    messages_test.go:879: msg: s_contains_rune must be contain rune '😊', but s_contains_rune is 'abc'
    messages_test.go:879: msg: s_ends_not_with must be not end with 'abc', but s_ends_not_with is 'abc'
    messages_test.go:879: msg: s_ends_with must be end with 'abc', but s_ends_with is 'zzz'
    messages_test.go:879: msg: s_excludes must be not contain 'xyz', but s_excludes is 'xyz'
    messages_test.go:879: msg: s_excludes_all must be contain none of 'abc', but s_excludes_all is 'b'
    messages_test.go:879: msg: s_excludes_rune must be not contain rune '😊', but s_excludes_rune is '😊'
    messages_test.go:879: msg: s_lowercase must lower-case, but s_lowercase is 'ABC'
    messages_test.go:879: msg: s_multibyte must multi-byte characters, but s_multibyte is 'abc'
    messages_test.go:879: msg: s_number must be a number, but s_number is 'abc'
    messages_test.go:879: msg: s_numeric must be numeric, but s_numeric is 'abc'
    messages_test.go:879: msg: s_print_ascii must printable ASCII, but s_print_ascii is 'é'
    messages_test.go:879: msg: s_starts_not_with must be not start with 'abc', but s_starts_not_with is 'abcXYZ'
    messages_test.go:879: msg: s_starts_with must be start with 'abc', but s_starts_with is 'XYZ'
    messages_test.go:879: msg: s_uppercase must upper-case, but s_uppercase is 'abc'
    messages_test.go:879: msg: f_base32 must be a valid Base32, but f_base32 is 'not-base32'
    messages_test.go:879: msg: f_base64 must be a valid Base64, but f_base64 is 'not_base64!!'
    messages_test.go:879: msg: f_base64_url must be a valid Base64URL, but f_base64_url is 'not-base64url'
    messages_test.go:879: msg: f_base64_raw_url must be a valid Base64RawURL, but f_base64_raw_url is 'not_base64!!'
    messages_test.go:879: msg: f_bic must be a valid BIC, but f_bic is 'INVALIDBIC'
    messages_test.go:879: msg: f_bcp47 must be a valid BCP-47 language tag, but f_bcp47 is 'invalid_tag'
    messages_test.go:879: msg: f_btc_addr must be a valid Bitcoin address, but f_btc_addr is 'InvalidBTCAddr'
    messages_test.go:879: msg: f_btc_bech32 must be a valid Bech32 Bitcoin address, but f_btc_bech32 is 'bc1invalidaddress'
    messages_test.go:879: msg: f_credit_card must be a valid credit-card number, but f_credit_card is '1234-5678-9012-3456'
    messages_test.go:879: msg: f_mongo_id must be a valid MongoDB ObjectID, but f_mongo_id is 'invalidobjectid'
    messages_test.go:879: msg: f_mongo_conn must be a valid MongoDB connection string, but f_mongo_conn is 'not-a-mongodb://connection'
    messages_test.go:879: msg: f_cron must be a valid cron expression, but f_cron is 'invalid cron'
    messages_test.go:879: msg: f_spice_db must be a valid SpiceDB identifier, but f_spice_db is 'bad type'
    messages_test.go:879: msg: f_date_time must a valid datetime, but f_date_time is 'not-a-date'
    messages_test.go:879: msg: f_e164 must be a valid E.164 phone number, but f_e164 is '+12345'
    messages_test.go:879: msg: f_ein must be a valid EIN, but f_ein is '12-3456-789'
    messages_test.go:879: msg: f_email must be a valid email, but f_email is 'invalid@'
    messages_test.go:879: msg: f_eth_addr must be a valid Ethereum address, but f_eth_addr is '0xINVALID'
    messages_test.go:879: msg: f_eth_checksum must be a checksummed Ethereum address, but f_eth_checksum is '0x52908400098527886E0F7030069857D2E4169EF'
    messages_test.go:879: msg: f_hexadecimal must be hexadecimal, but f_hexadecimal is 'Z123'
    messages_test.go:879: msg: f_hex_color must be a valid hex color, but f_hex_color is '#GGGGGG'
    messages_test.go:879: msg: f_hsl must be a valid HSL color, but f_hsl is 'hsl(361,100%,50%)'
    messages_test.go:879: msg: f_hsla must be a valid HSLA color, but f_hsla is 'hsla(361,100%,50%,1)'
    messages_test.go:879: msg: f_html must HTML tags, but f_html is '<unclosed tag'
    messages_test.go:879: msg: f_html_enc must HTML-encoded text, but f_html_enc is '&invalident;'
    messages_test.go:879: msg: f_isbn must be a valid ISBN, but f_isbn is 'invalidisbn'
    messages_test.go:879: msg: f_isbn10 must be a valid ISBN-10, but f_isbn10 is '123456789'
    messages_test.go:879: msg: f_isbn13 must be a valid ISBN-13, but f_isbn13 is '123456789012'
    messages_test.go:879: msg: f_issn must be a valid ISSN, but f_issn is '1234-INVALID'
    messages_test.go:879: msg: f_iso3166_a2 must be a valid ISO-3166-1 alpha-2 code, but f_iso3166_a2 is 'XX'
    messages_test.go:879: msg: f_iso3166_a2_eu must be a valid ISO-3166-1 alpha-2 EU code, but f_iso3166_a2_eu is 'XX'
    messages_test.go:879: msg: f_iso3166_a3 must be a valid ISO-3166-1 alpha-3 code, but f_iso3166_a3 is 'XXX'
    messages_test.go:879: msg: f_iso3166_a3_eu must be a valid ISO-3166-1 alpha-3 EU code, but f_iso3166_a3_eu is 'XXX'
    messages_test.go:879: msg: f_iso3166_num must be a valid ISO-3166-1 numeric code, but f_iso3166_num is '9999'
    messages_test.go:879: msg: f_iso3166_num_eu must be a valid ISO-3166-1 numeric EU code, but f_iso3166_num_eu is '9999'
    messages_test.go:879: msg: f_iso3166_2 must be a valid ISO-3166-2 code, but f_iso3166_2 is 'invalid'
    messages_test.go:879: msg: f_iso4217 must be a valid ISO-4217 currency code, but f_iso4217 is '123'
    messages_test.go:879: msg: f_iso4217_num must be a valid ISO-4217 numeric currency code, but f_iso4217_num is 1000
    messages_test.go:879: msg: f_json must be valid JSON, but f_json is '{invalid:json}'
    messages_test.go:879: msg: f_jwt must be a valid JWT, but f_jwt is 'header.payload'
    messages_test.go:879: msg: f_lat must be a valid latitude, but f_lat is '100.0000'
    messages_test.go:879: msg: f_lon must be a valid longitude, but f_lon is '200.0000'
    messages_test.go:879: msg: f_luhn must be a valid Luhn checksum, but f_luhn is '123456789'
    messages_test.go:879: msg: f_postcode must be a valid postcode, but f_postcode is 'ABCDE'
    messages_test.go:879: msg: f_postcode_field must be a valid postcode, but f_postcode_field is '12345'
    messages_test.go:879: msg: f_rgb must be a valid RGB color, but f_rgb is 'rgb(256,0,0)'
    messages_test.go:879: msg: f_rgba must be a valid RGBA color, but f_rgba is 'rgba(0,256,0,1)'
    messages_test.go:879: msg: f_ssn must be a valid SSN, but f_ssn is '000-00-0000'
    messages_test.go:879: msg: f_timezone must be a valid timezone, but f_timezone is 'Invalid/Zone'
    messages_test.go:879: msg: f_uuid must be a valid UUID, but f_uuid is 'invalid-uuid'
    messages_test.go:879: msg: f_uuid3 must be a valid UUID v3, but f_uuid3 is '1234'
    messages_test.go:879: msg: f_uuid3_rfc must be a valid UUID v3 (RFC4122), but f_uuid3_rfc is 'invalid-uuid3'
    messages_test.go:879: msg: f_uuid4 must be a valid UUID v4, but f_uuid4 is 'invalid-uuid4'
    messages_test.go:879: msg: f_uuid4_rfc must be a valid UUID v4 (RFC4122), but f_uuid4_rfc is 'invalid-uuid4'
    messages_test.go:879: msg: f_uuid5 must be a valid UUID v5, but f_uuid5 is 'invalid-uuid5'
    messages_test.go:879: msg: f_uuid5_rfc must be a valid UUID v5 (RFC4122), but f_uuid5_rfc is 'invalid-uuid5'
    messages_test.go:879: msg: f_uuid_rfc must be a valid UUID (RFC4122), but f_uuid_rfc is 'invalid-uuid'
    messages_test.go:879: msg: f_md4 must be a valid MD4 hash, but f_md4 is 'invalidmd4'
    messages_test.go:879: msg: f_md5 must be a valid MD5 hash, but f_md5 is 'invalidmd5'
    messages_test.go:879: msg: f_sha256 must be a valid SHA-256 hash, but f_sha256 is 'invalidsha256'
    messages_test.go:879: msg: f_sha384 must be a valid SHA-384 hash, but f_sha384 is 'invalidsha384'
    messages_test.go:879: msg: f_sha512 must be a valid SHA-512 hash, but f_sha512 is 'invalidsha512'
    messages_test.go:879: msg: f_ripemd128 must be a valid RIPEMD-128 hash, but f_ripemd128 is 'invalidripemd128'
    messages_test.go:879: msg: f_ripemd160 must be a valid RIPEMD-160 hash, but f_ripemd160 is 'invalidripemd160'
    messages_test.go:879: msg: f_tiger128 must be a valid TIGER-128 hash, but f_tiger128 is 'invalidtiger128'
    messages_test.go:879: msg: f_tiger160 must be a valid TIGER-160 hash, but f_tiger160 is 'invalidtiger160'
    messages_test.go:879: msg: f_tiger192 must be a valid TIGER-192 hash, but f_tiger192 is 'invalidtiger192'
    messages_test.go:879: msg: f_semver must be a valid semantic version, but f_semver is 'v1.x'
    messages_test.go:879: msg: f_ulid must be a valid ULID, but f_ulid is '01ARZ3NDEKTSV4RRFFQ69G5'
    messages_test.go:879: msg: f_cve must be a valid CVE identifier, but f_cve is 'CVE-0000-0000'
    messages_test.go:879: msg: c_eq must be  'abc', but c_eq is 'xyz'
    messages_test.go:879: msg: c_eq_ic must be  'abc', but c_eq_ic is 'DEF'
    messages_test.go:879: msg: c_gt must be greater than 10, but c_gt is 0
    messages_test.go:879: msg: c_gte must be at least 10, but c_gte is 9
    messages_test.go:879: msg: c_lt must be less than 10, but c_lt is 10
    messages_test.go:879: msg: c_lte must be at most 10, but c_lte is 11
    messages_test.go:879: msg: c_ne must be not equal to 'abc', but c_ne is 'abc'
    messages_test.go:879: msg: c_ne_ic must be not equal (case-insensitive) to 'abc', but c_ne_ic is 'ABC'
    messages_test.go:879: msg: o_dir must be an existing directory, but o_dir is '/path/does/not/exist'
    messages_test.go:879: msg: o_dir_path must be a valid directory path, but o_dir_path is '/invalid//dir//path'
    messages_test.go:879: msg: o_file must be an existing file, but o_file is '/path/to/nonexistent.file'
    messages_test.go:879: msg: o_file_path must be a valid file path, but o_file_path is 'invalidpath'
    messages_test.go:879: msg: o_image must be a valid image, but o_image is 'not-an-image-content'
    messages_test.go:879: msg: o_is_default must the default value, but o_is_default is 1
    messages_test.go:879: msg: o_len must be length 5, but o_len is 'too_long'
    messages_test.go:879: msg: o_max must be at most 5, but o_max is 6
    messages_test.go:879: msg: o_min must be at least 5, but o_min is 4
    messages_test.go:879: msg: o_one_of must be one of 'a b c', but o_one_of is 'invalid_choice'
    messages_test.go:879: msg: c_one_of_ci must be one of (case-insensitive) 'a b c', but c_one_of_ci is 'InvalidChoice'
    messages_test.go:879: msg: o_skip_unless must be validated unless O_Other xyz, but o_skip_unless is ''
    messages_test.go:879: msg: o_required must set, but o_required is ''
    messages_test.go:879: msg: o_required_if must be set if O_Other xyz, but o_required_if is ''
    messages_test.go:879: msg: o_required_with must be set with O_Other, but o_required_with is ''
    messages_test.go:879: msg: o_required_with_all must be set with all O_Other O_Another, but o_required_with_all is ''
    messages_test.go:879: msg: o_excluded_if must be excluded if O_Other xyz, but o_excluded_if is 'foo'
    messages_test.go:879: msg: o_excluded_with must be excluded with O_Other, but o_excluded_with is 'foo'
    messages_test.go:879: msg: o_excluded_with_all must be excluded with all O_Other O_Another, but o_excluded_with_all is 'foo'
    messages_test.go:879: msg: o_unique_slice must unique, but o_unique_slice is [1 1]
    messages_test.go:879: msg: o_validate_fn must valid according to Validate(), but o_validate_fn is {}
    messages_test.go:879: msg: a_is_color must be a valid color (hex, rgb[a], hsl[a]), but a_is_color is 'not-a-color'
    messages_test.go:879: msg: a_country_code must be a valid country code, but a_country_code is 'ZZZ'
    messages_test.go:879: msg: a_eu_country must be a valid EU country code, but a_eu_country is 'XYZ'
    messages_test.go:885: 
    messages_test.go:879: msg: f_eq_cs_field_a must be equal to f_eq_cs_field_b (value: 2), but f_eq_cs_field_a is 1
    messages_test.go:879: msg: f_eq_field_a must be equal to f_eq_field_b (value: 'bar'), but f_eq_field_a is 'foo'
    messages_test.go:879: msg: f_field_contains must contain f_field_contains_src (value: 'xyz'), but f_field_contains is 'abcd'
    messages_test.go:879: msg: f_field_excludes must not contain f_field_excludes_src (value: 'y'), but f_field_excludes is 'xyz'
    messages_test.go:879: msg: f_gt_cs_field_a must be greater than f_gt_cs_field_b (value: 2), but f_gt_cs_field_a is 1
    messages_test.go:879: msg: f_gte_cs_field_a must be greater than or equal to f_gte_cs_field_b (value: 2), but f_gte_cs_field_a is 1
    messages_test.go:879: msg: f_gte_field_a must be greater than or equal to f_gte_field_b (value: 10), but f_gte_field_a is 9
    messages_test.go:879: msg: f_gt_field_a must be greater than f_gt_field_b (value: 10), but f_gt_field_a is 10
    messages_test.go:879: msg: f_lt_cs_field_a must be less than f_lt_cs_field_b (value: 1), but f_lt_cs_field_a is 2
    messages_test.go:879: msg: f_lte_cs_field_a must be less than or equal to f_lte_cs_field_b (value: 1), but f_lte_cs_field_a is 2
    messages_test.go:879: msg: f_lte_field_a must be less than or equal to f_lte_field_b (value: 10), but f_lte_field_a is 11
    messages_test.go:879: msg: f_lt_field_a must be less than f_lt_field_b (value: 10), but f_lt_field_a is 10
    messages_test.go:879: msg: f_ne_cs_field_a must be different from f_ne_cs_field_b (value: 1), but f_ne_cs_field_a is 1
    messages_test.go:879: msg: f_ne_field_a must be different from f_ne_field_b (value: 1), but f_ne_field_a is 1
    messages_test.go:879: msg: n_cidr must be a valid CIDR, but n_cidr is 'invalid'
    messages_test.go:879: msg: n_cidr_v4 must be a valid CIDRv4, but n_cidr_v4 is 'invalid'
    messages_test.go:879: msg: n_cidr_v6 must be a valid CIDRv6, but n_cidr_v6 is 'invalid'
    messages_test.go:879: msg: n_data_uri must be a valid Data URI, but n_data_uri is 'invalid'
    messages_test.go:879: msg: n_fqdn must be a valid FQDN, but n_fqdn is 'invalid'
    messages_test.go:879: msg: n_hostname must be a valid hostname, but n_hostname is 'host name with spaces'
    messages_test.go:879: msg: n_hostname_port must be a valid host:port, but n_hostname_port is 'invalid'
    messages_test.go:879: msg: n_port must be a valid TCP/UDP port, but n_port is 0
    messages_test.go:879: msg: n_hostname_rfc1123 must be a valid RFC 1123 hostname, but n_hostname_rfc1123 is '-starts.with.dash'
    messages_test.go:879: msg: n_dns_label must be a valid RFC1035 DNS label, but n_dns_label is 'Invalid_Label!'
    messages_test.go:879: msg: n_ip must be a valid IP address, but n_ip is '999.999.999.999'
    messages_test.go:879: msg: n_ip4_addr must be a valid IPv4 address, but n_ip4_addr is '999.999.999.999'
    messages_test.go:879: msg: n_ip6_addr must be a valid IPv6 address, but n_ip6_addr is 'zzzz::zzzz'
    messages_test.go:879: msg: n_ip_addr must be a valid IP address, but n_ip_addr is 'not-an-ip'
    messages_test.go:879: msg: n_ip_v4 must be a valid IPv4 address, but n_ip_v4 is '300.300.300.300'
    messages_test.go:879: msg: n_ip_v6 must be a valid IPv6 address, but n_ip_v6 is 'gggg::gggg'
    messages_test.go:879: msg: n_mac must be a valid MAC address, but n_mac is '00:00:00:00:00:zz'
    messages_test.go:879: msg: n_tcp4_addr must be a valid TCPv4 address, but n_tcp4_addr is '999.999.999.999:99999'
    messages_test.go:879: msg: n_tcp6_addr must be a valid TCPv6 address, but n_tcp6_addr is '[gggg::gggg]:99999'
    messages_test.go:879: msg: n_tcp_addr must be a valid TCP address, but n_tcp_addr is 'not-a-tcp-addr'
    messages_test.go:879: msg: n_udp4_addr must be a valid UDPv4 address, but n_udp4_addr is '999.999.999.999:99999'
    messages_test.go:879: msg: n_udp6_addr must be a valid UDPv6 address, but n_udp6_addr is '[gggg::gggg]:99999'
    messages_test.go:879: msg: n_udp_addr must be a valid UDP address, but n_udp_addr is 'not-a-udp-addr'
    messages_test.go:879: msg: n_uri must be a valid URI, but n_uri is 'ht!tp://%%%'
    messages_test.go:879: msg: n_url must be a valid URL, but n_url is 'http:// invalid .com'
    messages_test.go:879: msg: n_http_url must be a valid HTTP URL, but n_http_url is 'httpp://example..com'
    messages_test.go:879: msg: n_url_encoded must URL-encoded data, but n_url_encoded is 'abc%zz'
    messages_test.go:879: msg: n_urn must be a valid RFC 2141 URN, but n_urn is 'urn:invalid:%%%'
    messages_test.go:879: msg: s_alpha must alphabetic characters, but s_alpha is '123'
    messages_test.go:879: msg: s_alnum must alphanumeric characters, but s_alnum is '___'
    messages_test.go:879: msg: s_alnum_unicode must alphanumeric Unicode characters, but s_alnum_unicode is 'abc!'
    messages_test.go:879: msg: s_alpha_unicode must alphabetic Unicode characters, but s_alpha_unicode is 'abc123'
    messages_test.go:879: msg: s_ascii must ASCII characters, but s_ascii is '©'
    messages_test.go:879: msg: s_bool must be a boolean, but s_bool is 'maybe'
    messages_test.go:879: msg: s_contains must be contain 'xyz', but s_contains is 'foo'
    messages_test.go:879: msg: s_contains_any must be contain any of 'xyz', but s_contains_any is ''
    messages_test.go:879: msg: s_contains_rune must be contain rune '😊', but s_contains_rune is 'abc'
    messages_test.go:879: msg: s_ends_not_with must be not end with 'abc', but s_ends_not_with is 'abc'
    messages_test.go:879: msg: s_ends_with must be end with 'abc', but s_ends_with is 'zzz'
    messages_test.go:879: msg: s_excludes must be not contain 'xyz', but s_excludes is 'xyz'
    messages_test.go:879: msg: s_excludes_all must be contain none of 'abc', but s_excludes_all is 'b'
    messages_test.go:879: msg: s_excludes_rune must be not contain rune '😊', but s_excludes_rune is '😊'
    messages_test.go:879: msg: s_lowercase must lower-case, but s_lowercase is 'ABC'
    messages_test.go:879: msg: s_multibyte must multi-byte characters, but s_multibyte is 'abc'
    messages_test.go:879: msg: s_number must be a number, but s_number is 'abc'
    messages_test.go:879: msg: s_numeric must be numeric, but s_numeric is 'abc'
    messages_test.go:879: msg: s_print_ascii must printable ASCII, but s_print_ascii is 'é'
    messages_test.go:879: msg: s_starts_not_with must be not start with 'abc', but s_starts_not_with is 'abcXYZ'
    messages_test.go:879: msg: s_starts_with must be start with 'abc', but s_starts_with is 'XYZ'
    messages_test.go:879: msg: s_uppercase must upper-case, but s_uppercase is 'abc'
    messages_test.go:879: msg: f_base32 must be a valid Base32, but f_base32 is 'not-base32'
    messages_test.go:879: msg: f_base64 must be a valid Base64, but f_base64 is 'not_base64!!'
    messages_test.go:879: msg: f_base64_url must be a valid Base64URL, but f_base64_url is 'not-base64url'
    messages_test.go:879: msg: f_base64_raw_url must be a valid Base64RawURL, but f_base64_raw_url is 'not_base64!!'
    messages_test.go:879: msg: f_bic must be a valid BIC, but f_bic is 'INVALIDBIC'
    messages_test.go:879: msg: f_bcp47 must be a valid BCP-47 language tag, but f_bcp47 is 'invalid_tag'
    messages_test.go:879: msg: f_btc_addr must be a valid Bitcoin address, but f_btc_addr is 'InvalidBTCAddr'
    messages_test.go:879: msg: f_btc_bech32 must be a valid Bech32 Bitcoin address, but f_btc_bech32 is 'bc1invalidaddress'
    messages_test.go:879: msg: f_credit_card must be a valid credit-card number, but f_credit_card is '1234-5678-9012-3456'
    messages_test.go:879: msg: f_mongo_id must be a valid MongoDB ObjectID, but f_mongo_id is 'invalidobjectid'
    messages_test.go:879: msg: f_mongo_conn must be a valid MongoDB connection string, but f_mongo_conn is 'not-a-mongodb://connection'
    messages_test.go:879: msg: f_cron must be a valid cron expression, but f_cron is 'invalid cron'
    messages_test.go:879: msg: f_spice_db must be a valid SpiceDB identifier, but f_spice_db is 'bad type'
    messages_test.go:879: msg: f_date_time must a valid datetime, but f_date_time is 'not-a-date'
    messages_test.go:879: msg: f_e164 must be a valid E.164 phone number, but f_e164 is '+12345'
    messages_test.go:879: msg: f_ein must be a valid EIN, but f_ein is '12-3456-789'
    messages_test.go:879: msg: f_email must be a valid email, but f_email is 'invalid@'
    messages_test.go:879: msg: f_eth_addr must be a valid Ethereum address, but f_eth_addr is '0xINVALID'
    messages_test.go:879: msg: f_eth_checksum must be a checksummed Ethereum address, but f_eth_checksum is '0x52908400098527886E0F7030069857D2E4169EF'
    messages_test.go:879: msg: f_hexadecimal must be hexadecimal, but f_hexadecimal is 'Z123'
    messages_test.go:879: msg: f_hex_color must be a valid hex color, but f_hex_color is '#GGGGGG'
    messages_test.go:879: msg: f_hsl must be a valid HSL color, but f_hsl is 'hsl(361,100%,50%)'
    messages_test.go:879: msg: f_hsla must be a valid HSLA color, but f_hsla is 'hsla(361,100%,50%,1)'
    messages_test.go:879: msg: f_html must HTML tags, but f_html is '<unclosed tag'
    messages_test.go:879: msg: f_html_enc must HTML-encoded text, but f_html_enc is '&invalident;'
    messages_test.go:879: msg: f_isbn must be a valid ISBN, but f_isbn is 'invalidisbn'
    messages_test.go:879: msg: f_isbn10 must be a valid ISBN-10, but f_isbn10 is '123456789'
    messages_test.go:879: msg: f_isbn13 must be a valid ISBN-13, but f_isbn13 is '123456789012'
    messages_test.go:879: msg: f_issn must be a valid ISSN, but f_issn is '1234-INVALID'
    messages_test.go:879: msg: f_iso3166_a2 must be a valid ISO-3166-1 alpha-2 code, but f_iso3166_a2 is 'XX'
    messages_test.go:879: msg: f_iso3166_a2_eu must be a valid ISO-3166-1 alpha-2 EU code, but f_iso3166_a2_eu is 'XX'
    messages_test.go:879: msg: f_iso3166_a3 must be a valid ISO-3166-1 alpha-3 code, but f_iso3166_a3 is 'XXX'
    messages_test.go:879: msg: f_iso3166_a3_eu must be a valid ISO-3166-1 alpha-3 EU code, but f_iso3166_a3_eu is 'XXX'
    messages_test.go:879: msg: f_iso3166_num must be a valid ISO-3166-1 numeric code, but f_iso3166_num is '9999'
    messages_test.go:879: msg: f_iso3166_num_eu must be a valid ISO-3166-1 numeric EU code, but f_iso3166_num_eu is '9999'
    messages_test.go:879: msg: f_iso3166_2 must be a valid ISO-3166-2 code, but f_iso3166_2 is 'invalid'
    messages_test.go:879: msg: f_iso4217 must be a valid ISO-4217 currency code, but f_iso4217 is '123'
    messages_test.go:879: msg: f_iso4217_num must be a valid ISO-4217 numeric currency code, but f_iso4217_num is 1000
    messages_test.go:879: msg: f_json must be valid JSON, but f_json is '{invalid:json}'
    messages_test.go:879: msg: f_jwt must be a valid JWT, but f_jwt is 'header.payload'
    messages_test.go:879: msg: f_lat must be a valid latitude, but f_lat is '100.0000'
    messages_test.go:879: msg: f_lon must be a valid longitude, but f_lon is '200.0000'
    messages_test.go:879: msg: f_luhn must be a valid Luhn checksum, but f_luhn is '123456789'
    messages_test.go:879: msg: f_postcode must be a valid postcode, but f_postcode is 'ABCDE'
    messages_test.go:879: msg: f_postcode_field must be a valid postcode, but f_postcode_field is '12345'
    messages_test.go:879: msg: f_rgb must be a valid RGB color, but f_rgb is 'rgb(256,0,0)'
    messages_test.go:879: msg: f_rgba must be a valid RGBA color, but f_rgba is 'rgba(0,256,0,1)'
    messages_test.go:879: msg: f_ssn must be a valid SSN, but f_ssn is '000-00-0000'
    messages_test.go:879: msg: f_timezone must be a valid timezone, but f_timezone is 'Invalid/Zone'
    messages_test.go:879: msg: f_uuid must be a valid UUID, but f_uuid is 'invalid-uuid'
    messages_test.go:879: msg: f_uuid3 must be a valid UUID v3, but f_uuid3 is '1234'
    messages_test.go:879: msg: f_uuid3_rfc must be a valid UUID v3 (RFC4122), but f_uuid3_rfc is 'invalid-uuid3'
    messages_test.go:879: msg: f_uuid4 must be a valid UUID v4, but f_uuid4 is 'invalid-uuid4'
    messages_test.go:879: msg: f_uuid4_rfc must be a valid UUID v4 (RFC4122), but f_uuid4_rfc is 'invalid-uuid4'
    messages_test.go:879: msg: f_uuid5 must be a valid UUID v5, but f_uuid5 is 'invalid-uuid5'
    messages_test.go:879: msg: f_uuid5_rfc must be a valid UUID v5 (RFC4122), but f_uuid5_rfc is 'invalid-uuid5'
    messages_test.go:879: msg: f_uuid_rfc must be a valid UUID (RFC4122), but f_uuid_rfc is 'invalid-uuid'
    messages_test.go:879: msg: f_md4 must be a valid MD4 hash, but f_md4 is 'invalidmd4'
    messages_test.go:879: msg: f_md5 must be a valid MD5 hash, but f_md5 is 'invalidmd5'
    messages_test.go:879: msg: f_sha256 must be a valid SHA-256 hash, but f_sha256 is 'invalidsha256'
    messages_test.go:879: msg: f_sha384 must be a valid SHA-384 hash, but f_sha384 is 'invalidsha384'
    messages_test.go:879: msg: f_sha512 must be a valid SHA-512 hash, but f_sha512 is 'invalidsha512'
    messages_test.go:879: msg: f_ripemd128 must be a valid RIPEMD-128 hash, but f_ripemd128 is 'invalidripemd128'
    messages_test.go:879: msg: f_ripemd160 must be a valid RIPEMD-160 hash, but f_ripemd160 is 'invalidripemd160'
    messages_test.go:879: msg: f_tiger128 must be a valid TIGER-128 hash, but f_tiger128 is 'invalidtiger128'
    messages_test.go:879: msg: f_tiger160 must be a valid TIGER-160 hash, but f_tiger160 is 'invalidtiger160'
    messages_test.go:879: msg: f_tiger192 must be a valid TIGER-192 hash, but f_tiger192 is 'invalidtiger192'
    messages_test.go:879: msg: f_semver must be a valid semantic version, but f_semver is 'v1.x'
    messages_test.go:879: msg: f_ulid must be a valid ULID, but f_ulid is '01ARZ3NDEKTSV4RRFFQ69G5'
    messages_test.go:879: msg: f_cve must be a valid CVE identifier, but f_cve is 'CVE-0000-0000'
    messages_test.go:879: msg: c_eq must be  'abc', but c_eq is 'xyz'
    messages_test.go:879: msg: c_eq_ic must be  'abc', but c_eq_ic is 'DEF'
    messages_test.go:879: msg: c_gt must be greater than 10, but c_gt is 0
    messages_test.go:879: msg: c_gte must be at least 10, but c_gte is 9
    messages_test.go:879: msg: c_lt must be less than 10, but c_lt is 10
    messages_test.go:879: msg: c_lte must be at most 10, but c_lte is 11
    messages_test.go:879: msg: c_ne must be not equal to 'abc', but c_ne is 'abc'
    messages_test.go:879: msg: c_ne_ic must be not equal (case-insensitive) to 'abc', but c_ne_ic is 'ABC'
    messages_test.go:879: msg: o_dir must be an existing directory, but o_dir is '/path/does/not/exist'
    messages_test.go:879: msg: o_dir_path must be a valid directory path, but o_dir_path is '/invalid//dir//path'
    messages_test.go:879: msg: o_file must be an existing file, but o_file is '/path/to/nonexistent.file'
    messages_test.go:879: msg: o_file_path must be a valid file path, but o_file_path is 'invalidpath'
    messages_test.go:879: msg: o_image must be a valid image, but o_image is 'not-an-image-content'
    messages_test.go:879: msg: o_is_default must the default value, but o_is_default is 1
    messages_test.go:879: msg: o_len must be length 5, but o_len is 'too_long'
    messages_test.go:879: msg: o_max must be at most 5, but o_max is 6
    messages_test.go:879: msg: o_min must be at least 5, but o_min is 4
    messages_test.go:879: msg: o_one_of must be one of 'a b c', but o_one_of is 'invalid_choice'
    messages_test.go:879: msg: c_one_of_ci must be one of (case-insensitive) 'a b c', but c_one_of_ci is 'InvalidChoice'
    messages_test.go:879: msg: o_required must set, but o_required is ''
    messages_test.go:879: msg: o_required_unless must be set unless O_Other xyz, but o_required_unless is ''
    messages_test.go:879: msg: o_required_without must be set without O_Other, but o_required_without is ''
    messages_test.go:879: msg: o_required_without_all must be set without all O_Other O_Another, but o_required_without_all is ''
    messages_test.go:879: msg: o_excluded_unless must be excluded unless O_Other xyz, but o_excluded_unless is 'foo'
    messages_test.go:879: msg: o_excluded_without must be excluded without O_Other, but o_excluded_without is 'foo'
    messages_test.go:879: msg: o_excluded_without_all must be excluded without all O_Other O_Another, but o_excluded_without_all is 'foo'
    messages_test.go:879: msg: o_unique_slice must unique, but o_unique_slice is [1 1]
    messages_test.go:879: msg: o_validate_fn must valid according to Validate(), but o_validate_fn is {}
    messages_test.go:879: msg: a_is_color must be a valid color (hex, rgb[a], hsl[a]), but a_is_color is 'not-a-color'
    messages_test.go:879: msg: a_country_code must be a valid country code, but a_country_code is 'ZZZ'
    messages_test.go:879: msg: a_eu_country must be a valid EU country code, but a_eu_country is 'XYZ'
    messages_test.go:885: 
    messages_test.go:879: msg: f_eq_cs_field_a must be equal to f_eq_cs_field_b (value: 2), but f_eq_cs_field_a is 1
    messages_test.go:879: msg: f_eq_field_a must be equal to f_eq_field_b (value: 'bar'), but f_eq_field_a is 'foo'
    messages_test.go:879: msg: f_field_contains must contain f_field_contains_src (value: 'xyz'), but f_field_contains is 'abcd'
    messages_test.go:879: msg: f_field_excludes must not contain f_field_excludes_src (value: 'y'), but f_field_excludes is 'xyz'
    messages_test.go:879: msg: f_gt_cs_field_a must be greater than f_gt_cs_field_b (value: 2), but f_gt_cs_field_a is 1
    messages_test.go:879: msg: f_gte_cs_field_a must be greater than or equal to f_gte_cs_field_b (value: 2), but f_gte_cs_field_a is 1
    messages_test.go:879: msg: f_gte_field_a must be greater than or equal to f_gte_field_b (value: 10), but f_gte_field_a is 9
    messages_test.go:879: msg: f_gt_field_a must be greater than f_gt_field_b (value: 10), but f_gt_field_a is 10
    messages_test.go:879: msg: f_lt_cs_field_a must be less than f_lt_cs_field_b (value: 1), but f_lt_cs_field_a is 2
    messages_test.go:879: msg: f_lte_cs_field_a must be less than or equal to f_lte_cs_field_b (value: 1), but f_lte_cs_field_a is 2
    messages_test.go:879: msg: f_lte_field_a must be less than or equal to f_lte_field_b (value: 10), but f_lte_field_a is 11
    messages_test.go:879: msg: f_lt_field_a must be less than f_lt_field_b (value: 10), but f_lt_field_a is 10
    messages_test.go:879: msg: f_ne_cs_field_a must be different from f_ne_cs_field_b (value: 1), but f_ne_cs_field_a is 1
    messages_test.go:879: msg: f_ne_field_a must be different from f_ne_field_b (value: 1), but f_ne_field_a is 1
    messages_test.go:879: msg: n_cidr must be a valid CIDR, but n_cidr is 'invalid'
    messages_test.go:879: msg: n_cidr_v4 must be a valid CIDRv4, but n_cidr_v4 is 'invalid'
    messages_test.go:879: msg: n_cidr_v6 must be a valid CIDRv6, but n_cidr_v6 is 'invalid'
    messages_test.go:879: msg: n_data_uri must be a valid Data URI, but n_data_uri is 'invalid'
    messages_test.go:879: msg: n_fqdn must be a valid FQDN, but n_fqdn is 'invalid'
    messages_test.go:879: msg: n_hostname must be a valid hostname, but n_hostname is 'host name with spaces'
    messages_test.go:879: msg: n_hostname_port must be a valid host:port, but n_hostname_port is 'invalid'
    messages_test.go:879: msg: n_port must be a valid TCP/UDP port, but n_port is 0
    messages_test.go:879: msg: n_hostname_rfc1123 must be a valid RFC 1123 hostname, but n_hostname_rfc1123 is '-starts.with.dash'
    messages_test.go:879: msg: n_dns_label must be a valid RFC1035 DNS label, but n_dns_label is 'Invalid_Label!'
    messages_test.go:879: msg: n_ip must be a valid IP address, but n_ip is '999.999.999.999'
    messages_test.go:879: msg: n_ip4_addr must be a valid IPv4 address, but n_ip4_addr is '999.999.999.999'
    messages_test.go:879: msg: n_ip6_addr must be a valid IPv6 address, but n_ip6_addr is 'zzzz::zzzz'
    messages_test.go:879: msg: n_ip_addr must be a valid IP address, but n_ip_addr is 'not-an-ip'
    messages_test.go:879: msg: n_ip_v4 must be a valid IPv4 address, but n_ip_v4 is '300.300.300.300'
    messages_test.go:879: msg: n_ip_v6 must be a valid IPv6 address, but n_ip_v6 is 'gggg::gggg'
    messages_test.go:879: msg: n_mac must be a valid MAC address, but n_mac is '00:00:00:00:00:zz'
    messages_test.go:879: msg: n_tcp4_addr must be a valid TCPv4 address, but n_tcp4_addr is '999.999.999.999:99999'
    messages_test.go:879: msg: n_tcp6_addr must be a valid TCPv6 address, but n_tcp6_addr is '[gggg::gggg]:99999'
    messages_test.go:879: msg: n_tcp_addr must be a valid TCP address, but n_tcp_addr is 'not-a-tcp-addr'
    messages_test.go:879: msg: n_udp4_addr must be a valid UDPv4 address, but n_udp4_addr is '999.999.999.999:99999'
    messages_test.go:879: msg: n_udp6_addr must be a valid UDPv6 address, but n_udp6_addr is '[gggg::gggg]:99999'
    messages_test.go:879: msg: n_udp_addr must be a valid UDP address, but n_udp_addr is 'not-a-udp-addr'
    messages_test.go:879: msg: n_uri must be a valid URI, but n_uri is 'ht!tp://%%%'
    messages_test.go:879: msg: n_url must be a valid URL, but n_url is 'http:// invalid .com'
    messages_test.go:879: msg: n_http_url must be a valid HTTP URL, but n_http_url is 'httpp://example..com'
    messages_test.go:879: msg: n_url_encoded must URL-encoded data, but n_url_encoded is 'abc%zz'
    messages_test.go:879: msg: n_urn must be a valid RFC 2141 URN, but n_urn is 'urn:invalid:%%%'
    messages_test.go:879: msg: s_alpha must alphabetic characters, but s_alpha is '123'
    messages_test.go:879: msg: s_alnum must alphanumeric characters, but s_alnum is '___'
    messages_test.go:879: msg: s_alnum_unicode must alphanumeric Unicode characters, but s_alnum_unicode is 'abc!'
    messages_test.go:879: msg: s_alpha_unicode must alphabetic Unicode characters, but s_alpha_unicode is 'abc123'
    messages_test.go:879: msg: s_ascii must ASCII characters, but s_ascii is '©'
    messages_test.go:879: msg: s_bool must be a boolean, but s_bool is 'maybe'
    messages_test.go:879: msg: s_contains must be contain 'xyz', but s_contains is 'foo'
    messages_test.go:879: msg: s_contains_any must be contain any of 'xyz', but s_contains_any is ''
    messages_test.go:879: msg: s_contains_rune must be contain rune '😊', but s_contains_rune is 'abc'
    messages_test.go:879: msg: s_ends_not_with must be not end with 'abc', but s_ends_not_with is 'abc'
    messages_test.go:879: msg: s_ends_with must be end with 'abc', but s_ends_with is 'zzz'
    messages_test.go:879: msg: s_excludes must be not contain 'xyz', but s_excludes is 'xyz'
    messages_test.go:879: msg: s_excludes_all must be contain none of 'abc', but s_excludes_all is 'b'
    messages_test.go:879: msg: s_excludes_rune must be not contain rune '😊', but s_excludes_rune is '😊'
    messages_test.go:879: msg: s_lowercase must lower-case, but s_lowercase is 'ABC'
    messages_test.go:879: msg: s_multibyte must multi-byte characters, but s_multibyte is 'abc'
    messages_test.go:879: msg: s_number must be a number, but s_number is 'abc'
    messages_test.go:879: msg: s_numeric must be numeric, but s_numeric is 'abc'
    messages_test.go:879: msg: s_print_ascii must printable ASCII, but s_print_ascii is 'é'
    messages_test.go:879: msg: s_starts_not_with must be not start with 'abc', but s_starts_not_with is 'abcXYZ'
    messages_test.go:879: msg: s_starts_with must be start with 'abc', but s_starts_with is 'XYZ'
    messages_test.go:879: msg: s_uppercase must upper-case, but s_uppercase is 'abc'
    messages_test.go:879: msg: f_base32 must be a valid Base32, but f_base32 is 'not-base32'
    messages_test.go:879: msg: f_base64 must be a valid Base64, but f_base64 is 'not_base64!!'
    messages_test.go:879: msg: f_base64_url must be a valid Base64URL, but f_base64_url is 'not-base64url'
    messages_test.go:879: msg: f_base64_raw_url must be a valid Base64RawURL, but f_base64_raw_url is 'not_base64!!'
    messages_test.go:879: msg: f_bic must be a valid BIC, but f_bic is 'INVALIDBIC'
    messages_test.go:879: msg: f_bcp47 must be a valid BCP-47 language tag, but f_bcp47 is 'invalid_tag'
    messages_test.go:879: msg: f_btc_addr must be a valid Bitcoin address, but f_btc_addr is 'InvalidBTCAddr'
    messages_test.go:879: msg: f_btc_bech32 must be a valid Bech32 Bitcoin address, but f_btc_bech32 is 'bc1invalidaddress'
    messages_test.go:879: msg: f_credit_card must be a valid credit-card number, but f_credit_card is '1234-5678-9012-3456'
    messages_test.go:879: msg: f_mongo_id must be a valid MongoDB ObjectID, but f_mongo_id is 'invalidobjectid'
    messages_test.go:879: msg: f_mongo_conn must be a valid MongoDB connection string, but f_mongo_conn is 'not-a-mongodb://connection'
    messages_test.go:879: msg: f_cron must be a valid cron expression, but f_cron is 'invalid cron'
    messages_test.go:879: msg: f_spice_db must be a valid SpiceDB identifier, but f_spice_db is 'bad type'
    messages_test.go:879: msg: f_date_time must a valid datetime, but f_date_time is 'not-a-date'
    messages_test.go:879: msg: f_e164 must be a valid E.164 phone number, but f_e164 is '+12345'
    messages_test.go:879: msg: f_ein must be a valid EIN, but f_ein is '12-3456-789'
    messages_test.go:879: msg: f_email must be a valid email, but f_email is 'invalid@'
    messages_test.go:879: msg: f_eth_addr must be a valid Ethereum address, but f_eth_addr is '0xINVALID'
    messages_test.go:879: msg: f_eth_checksum must be a checksummed Ethereum address, but f_eth_checksum is '0x52908400098527886E0F7030069857D2E4169EF'
    messages_test.go:879: msg: f_hexadecimal must be hexadecimal, but f_hexadecimal is 'Z123'
    messages_test.go:879: msg: f_hex_color must be a valid hex color, but f_hex_color is '#GGGGGG'
    messages_test.go:879: msg: f_hsl must be a valid HSL color, but f_hsl is 'hsl(361,100%,50%)'
    messages_test.go:879: msg: f_hsla must be a valid HSLA color, but f_hsla is 'hsla(361,100%,50%,1)'
    messages_test.go:879: msg: f_html must HTML tags, but f_html is '<unclosed tag'
    messages_test.go:879: msg: f_html_enc must HTML-encoded text, but f_html_enc is '&invalident;'
    messages_test.go:879: msg: f_isbn must be a valid ISBN, but f_isbn is 'invalidisbn'
    messages_test.go:879: msg: f_isbn10 must be a valid ISBN-10, but f_isbn10 is '123456789'
    messages_test.go:879: msg: f_isbn13 must be a valid ISBN-13, but f_isbn13 is '123456789012'
    messages_test.go:879: msg: f_issn must be a valid ISSN, but f_issn is '1234-INVALID'
    messages_test.go:879: msg: f_iso3166_a2 must be a valid ISO-3166-1 alpha-2 code, but f_iso3166_a2 is 'XX'
    messages_test.go:879: msg: f_iso3166_a2_eu must be a valid ISO-3166-1 alpha-2 EU code, but f_iso3166_a2_eu is 'XX'
    messages_test.go:879: msg: f_iso3166_a3 must be a valid ISO-3166-1 alpha-3 code, but f_iso3166_a3 is 'XXX'
    messages_test.go:879: msg: f_iso3166_a3_eu must be a valid ISO-3166-1 alpha-3 EU code, but f_iso3166_a3_eu is 'XXX'
    messages_test.go:879: msg: f_iso3166_num must be a valid ISO-3166-1 numeric code, but f_iso3166_num is '9999'
    messages_test.go:879: msg: f_iso3166_num_eu must be a valid ISO-3166-1 numeric EU code, but f_iso3166_num_eu is '9999'
    messages_test.go:879: msg: f_iso3166_2 must be a valid ISO-3166-2 code, but f_iso3166_2 is 'invalid'
    messages_test.go:879: msg: f_iso4217 must be a valid ISO-4217 currency code, but f_iso4217 is '123'
    messages_test.go:879: msg: f_iso4217_num must be a valid ISO-4217 numeric currency code, but f_iso4217_num is 1000
    messages_test.go:879: msg: f_json must be valid JSON, but f_json is '{invalid:json}'
    messages_test.go:879: msg: f_jwt must be a valid JWT, but f_jwt is 'header.payload'
    messages_test.go:879: msg: f_lat must be a valid latitude, but f_lat is '100.0000'
    messages_test.go:879: msg: f_lon must be a valid longitude, but f_lon is '200.0000'
    messages_test.go:879: msg: f_luhn must be a valid Luhn checksum, but f_luhn is '123456789'
    messages_test.go:879: msg: f_postcode must be a valid postcode, but f_postcode is 'ABCDE'
    messages_test.go:879: msg: f_postcode_field must be a valid postcode, but f_postcode_field is '12345'
    messages_test.go:879: msg: f_rgb must be a valid RGB color, but f_rgb is 'rgb(256,0,0)'
    messages_test.go:879: msg: f_rgba must be a valid RGBA color, but f_rgba is 'rgba(0,256,0,1)'
    messages_test.go:879: msg: f_ssn must be a valid SSN, but f_ssn is '000-00-0000'
    messages_test.go:879: msg: f_timezone must be a valid timezone, but f_timezone is 'Invalid/Zone'
    messages_test.go:879: msg: f_uuid must be a valid UUID, but f_uuid is 'invalid-uuid'
    messages_test.go:879: msg: f_uuid3 must be a valid UUID v3, but f_uuid3 is '1234'
    messages_test.go:879: msg: f_uuid3_rfc must be a valid UUID v3 (RFC4122), but f_uuid3_rfc is 'invalid-uuid3'
    messages_test.go:879: msg: f_uuid4 must be a valid UUID v4, but f_uuid4 is 'invalid-uuid4'
    messages_test.go:879: msg: f_uuid4_rfc must be a valid UUID v4 (RFC4122), but f_uuid4_rfc is 'invalid-uuid4'
    messages_test.go:879: msg: f_uuid5 must be a valid UUID v5, but f_uuid5 is 'invalid-uuid5'
    messages_test.go:879: msg: f_uuid5_rfc must be a valid UUID v5 (RFC4122), but f_uuid5_rfc is 'invalid-uuid5'
    messages_test.go:879: msg: f_uuid_rfc must be a valid UUID (RFC4122), but f_uuid_rfc is 'invalid-uuid'
    messages_test.go:879: msg: f_md4 must be a valid MD4 hash, but f_md4 is 'invalidmd4'
    messages_test.go:879: msg: f_md5 must be a valid MD5 hash, but f_md5 is 'invalidmd5'
    messages_test.go:879: msg: f_sha256 must be a valid SHA-256 hash, but f_sha256 is 'invalidsha256'
    messages_test.go:879: msg: f_sha384 must be a valid SHA-384 hash, but f_sha384 is 'invalidsha384'
    messages_test.go:879: msg: f_sha512 must be a valid SHA-512 hash, but f_sha512 is 'invalidsha512'
    messages_test.go:879: msg: f_ripemd128 must be a valid RIPEMD-128 hash, but f_ripemd128 is 'invalidripemd128'
    messages_test.go:879: msg: f_ripemd160 must be a valid RIPEMD-160 hash, but f_ripemd160 is 'invalidripemd160'
    messages_test.go:879: msg: f_tiger128 must be a valid TIGER-128 hash, but f_tiger128 is 'invalidtiger128'
    messages_test.go:879: msg: f_tiger160 must be a valid TIGER-160 hash, but f_tiger160 is 'invalidtiger160'
    messages_test.go:879: msg: f_tiger192 must be a valid TIGER-192 hash, but f_tiger192 is 'invalidtiger192'
    messages_test.go:879: msg: f_semver must be a valid semantic version, but f_semver is 'v1.x'
    messages_test.go:879: msg: f_ulid must be a valid ULID, but f_ulid is '01ARZ3NDEKTSV4RRFFQ69G5'
    messages_test.go:879: msg: f_cve must be a valid CVE identifier, but f_cve is 'CVE-0000-0000'
    messages_test.go:879: msg: c_eq must be  'abc', but c_eq is 'xyz'
    messages_test.go:879: msg: c_eq_ic must be  'abc', but c_eq_ic is 'DEF'
    messages_test.go:879: msg: c_gt must be greater than 10, but c_gt is 0
    messages_test.go:879: msg: c_gte must be at least 10, but c_gte is 9
    messages_test.go:879: msg: c_lt must be less than 10, but c_lt is 10
    messages_test.go:879: msg: c_lte must be at most 10, but c_lte is 11
    messages_test.go:879: msg: c_ne must be not equal to 'abc', but c_ne is 'abc'
    messages_test.go:879: msg: c_ne_ic must be not equal (case-insensitive) to 'abc', but c_ne_ic is 'ABC'
    messages_test.go:879: msg: o_dir must be an existing directory, but o_dir is '/path/does/not/exist'
    messages_test.go:879: msg: o_dir_path must be a valid directory path, but o_dir_path is '/invalid//dir//path'
    messages_test.go:879: msg: o_file must be an existing file, but o_file is '/path/to/nonexistent.file'
    messages_test.go:879: msg: o_file_path must be a valid file path, but o_file_path is 'invalidpath'
    messages_test.go:879: msg: o_image must be a valid image, but o_image is 'not-an-image-content'
    messages_test.go:879: msg: o_is_default must the default value, but o_is_default is 1
    messages_test.go:879: msg: o_len must be length 5, but o_len is 'too_long'
    messages_test.go:879: msg: o_max must be at most 5, but o_max is 6
    messages_test.go:879: msg: o_min must be at least 5, but o_min is 4
    messages_test.go:879: msg: o_one_of must be one of 'a b c', but o_one_of is 'invalid_choice'
    messages_test.go:879: msg: c_one_of_ci must be one of (case-insensitive) 'a b c', but c_one_of_ci is 'InvalidChoice'
    messages_test.go:879: msg: o_required must set, but o_required is ''
    messages_test.go:879: msg: o_required_unless must be set unless O_Other xyz, but o_required_unless is ''
    messages_test.go:879: msg: o_required_with must be set with O_Other, but o_required_with is ''
    messages_test.go:879: msg: o_required_with_all must be set with all O_Other O_Another, but o_required_with_all is ''
    messages_test.go:879: msg: o_excluded_unless must be excluded unless O_Other xyz, but o_excluded_unless is 'foo'
    messages_test.go:879: msg: o_excluded_with must be excluded with O_Other, but o_excluded_with is 'foo'
    messages_test.go:879: msg: o_excluded_with_all must be excluded with all O_Other O_Another, but o_excluded_with_all is 'foo'
    messages_test.go:879: msg: o_unique_slice must unique, but o_unique_slice is [1 1]
    messages_test.go:879: msg: o_validate_fn must valid according to Validate(), but o_validate_fn is {}
    messages_test.go:879: msg: a_is_color must be a valid color (hex, rgb[a], hsl[a]), but a_is_color is 'not-a-color'
    messages_test.go:879: msg: a_country_code must be a valid country code, but a_country_code is 'ZZZ'
    messages_test.go:879: msg: a_eu_country must be a valid EU country code, but a_eu_country is 'XYZ'
    messages_test.go:885: 
    messages_test.go:892: log: tag "unix_addr" did NOT fail, but tag "unix_addr" is in skipCoverage (known-always-pass validator)
--- PASS: Test_AllValidatorTags_UnhappyPath (0.01s)
PASS
ok      command-line-arguments  0.021s
```

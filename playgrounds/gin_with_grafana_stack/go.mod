module github.com/aazw/vscode-go/playgrounds/gin_with_grafana_stack

go 1.24.3

require (
	github.com/gin-gonic/gin v1.11.0
	github.com/grafana/pyroscope-go v1.2.7
	github.com/prometheus/client_golang v1.23.2
	github.com/urfave/cli/v3 v3.6.2
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.64.0
	go.opentelemetry.io/otel v1.39.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.39.0
	go.opentelemetry.io/otel/sdk v1.39.0
	go.opentelemetry.io/otel/trace v1.39.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic v1.14.2 // indirect
	github.com/bytedance/sonic/loader v0.4.0 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/gabriel-vasile/mimetype v1.4.11 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.28.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/goccy/go-yaml v1.19.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/pyroscope-go/godeltaprof v0.1.9 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.66.1 // indirect
	github.com/prometheus/procfs v0.16.1 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.57.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.1 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.39.0 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/proto/otlp v1.9.0 // indirect
	go.yaml.in/yaml/v2 v2.4.2 // indirect
	golang.org/x/arch v0.23.0 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20251222181119-0a764e51fe1b // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251222181119-0a764e51fe1b // indirect
	google.golang.org/grpc v1.78.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

// -----------------------------------------------------------------------------
// TEMPORARY WORKAROUND (genproto ambiguous import)
// -----------------------------------------------------------------------------
//
// 背景:
// - google.golang.org/genproto は過去の変更で「親モジュール」(google.golang.org/genproto) から
//   「分割サブモジュール」(google.golang.org/genproto/googleapis/...) へ段階的に移行している。
// - 依存グラフ内で「古い親モジュール」と「新しい分割サブモジュール」が同時に入ると、同一 import パスの
//   パッケージが “複数モジュールに存在する” 状態となり、Go の解決ができずビルドが失敗する。
//   例:
//     google.golang.org/genproto/googleapis/api/httpbody
//     google.golang.org/genproto/googleapis/rpc/status
//
// 症状:
// - go mod tidy を実行すると、以下のような ambiguous import が発生してビルド/依存解決が壊れる。
//   (手動で go get した genproto の新しい版が tidy により require から除去され、
//    依存側が要求する古い genproto が優先されるため)
//     ambiguous import: found package ... in multiple modules
//       google.golang.org/genproto v0.0.0-202304...
//       google.golang.org/genproto/googleapis/api (または /rpc) v0.0.0-20xx...
//
// 原因(今回の特定結果):
// - go mod graph の結果、github.com/cockroachdb/errors@v1.12.0 が
//   google.golang.org/genproto@v0.0.0-202304... を要求していた。
// - 一方で grpc / grpc-gateway / otel 等の経路から新しい分割サブモジュールが入ってきて混在し、衝突した。
//
// 暫定処置:
// - go.mod に replace を追加し、親モジュール google.golang.org/genproto を新しい版へ強制的に寄せる。
// - これにより依存グラフ上で古い親モジュールに戻らず、ambiguous import を回避できる。
// - 実施内容（固定している版は、分割サブモジュール側と同系統のコミット/日付に揃えること）:
//     replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20251213004720-97cd9d5aeac2
//
// 恒久対応（推奨）:
// - 依存元（今回だと github.com/cockroachdb/errors を引いている上位モジュール含む）を更新し、
//   古い google.golang.org/genproto を要求しない構成に整理する。
// - 恒久対応が完了したら、この replace は削除すること。
// -----------------------------------------------------------------------------
replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20251213004720-97cd9d5aeac2

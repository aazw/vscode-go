# playgrounds/gin_with_grafana_stack

## 概要

* Metrics
* Logging
* Tracing
* Profiling

```text
# Metrics
+---------+     +------------+     +-------+     +---------+
|         |<----|   Alloy    |---->|       |     |         |
|         |     +------------+     |       |     |         |
| Gin App |                        | Mimir |<----| Grafana |
|         |     +------------+     |       |     |         |
|         |<----| Prometheus |---->|       |     |         |
+---------+     +------------+     +-------+     +---------+

# Logging
+---------+     +--------+                  +----------+     +------+
|         |     |        |     +---|\       |          |     |      |
|         |     |        |     |   |_\      |          |     |      |
| Gin App |---->| Docker |---->| File |<----| Promtail |---->| Loki |
|         |     | Engine |     |      |     |          |     |      |
|         |     |        |     +------+     |          |     |      |
+---------+     +--------+                  +----------+     +------+

# Tracing
+---------+ 
|         | 
|         | 
| Gin App |-
|         | 
|         | 
+---------+ 

# Profiling
+---------+ 
|         | 
|         | 
| Gin App |-
|         | 
|         | 
+---------+ 
```



### Grafana (Jsonnet/Grafonnet)

* https://github.com/grafana/grafonnet
* https://github.com/grafana/grafonnet-lib ... archived

### Grafana Mimir

* Prometheusのメトリクスを長期保存するためのオープンソースTSDB(時系列データベース)
    * Prometheusの拡張バックエンド
    * 単一ノードでの短期保存にとどまるPrometheusと異なり、Mimirをバックエンドに組み合わせることで、メトリクスの保持期間を数ヶ月〜数年単位に延長しつつ、大量データを効率的に集約・検索できる
* 水平スケーラビリティ、高可用性、マルチテナント対応が特徴
* Prometheus互換のremote_write/Query API
* Alloyからの受信もできる
    * Apployもremote_writeプロトコルでMimirにデータを送信できるクライアント
    * 小規模環境ではPrometheus単体でMimirへの書き込みを行う
    * 大規模・複雑環境ではAlloyを用いて多様なデータソースを統一的に集約し、Mimirへ送信するパイプラインを構築する
* https://grafana.com/docs/mimir/latest/
    * https://grafana.com/docs/mimir/latest/get-started/
* https://github.com/grafana/mimir
* https://hub.docker.com/r/grafana/mimir
* OpenMetrics
    * OpenMetrics は Prometheus の Exposition Format を一般化した標準仕様
        * OpenMetrics はもともと Prometheus のテキスト／Protocol Buffers 形式を拡張・標準化したもの
    * メトリクスを「どう表現して HTTP で公開するか」を定めている
    * これに対応すれば、Prometheus がスクレイプするエンドポイントも、Prometheus 互換のバックエンド（Mimir を含む）も問題なく取り込める
* OpenMetrics 対応方法
    * https://github.com/prometheus/client_golang
        * github.com/prometheus/client_golang/prometheus
        * github.com/prometheus/client_golang/prometheus/promhttp
    * https://github.com/open-telemetry/opentelemetry-go
        * https://opentelemetry.io/docs/languages/go/getting-started/
* OpenMetrics vs OpenTelemetry
    * 両者は同じ CNCF の傘下にあるにもかかわらず別プロジェクトとして存在する
    * OpenMetrics
        * メトリクス公開フォーマット仕様
        * 元は CNCF Sandbox→Incubating→2024年にPrometheusへ再統合
        * OpenMetrics は OpenTelemetry の一部ではなく、あくまで独立した "メトリクス公開フォーマット" の仕様
        * メトリクスの「公開フォーマット仕様」（Exposition Format）のみを定める軽量な仕様
        * Prometheus や Mimir といったツールが「どういう形式で HTTP レスポンスを返せばいいか」を統一することを目的としている
    * OpenTelemetry
        * 観測データ収集フレームワーク（Metrics, Traces, Logs）
        * OpenTelemetry はOpenMetricsの仕様を「エクスポーター」のひとつとして採用し、相互運用性を実現している
        * OpenTelemetry は Cloud Native Computing Foundation（CNCF）のインキュベーティングプロジェクト（Sandbox→Incubating）
        * トレース／メトリクス／ログといったあらゆる観測データの 収集から送信まで をカバーする包括的フレームワーク
        * SDK や API、OTLP プロトコル、各種エクスポーターを提供し、多様なバックエンドとの連携を実現する
* https://grafana.com/docs/mimir/latest/configure/configuration-parameters/
* https://community.zenduty.com/t/how-to-properly-configure-mimir-data-source-in-grafana/976/3

### Prometheus

* https://prometheus.io/
* https://hub.docker.com/r/prom/prometheus
* https://github.com/prometheus/prometheus
* https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config

### Grafana Alloy

* Alloy は公式に prometheus.scrape コンポーネントを提供し、Prometheus と同じく HTTP エンドポイントをスクレイプ（pull）できる
* prometheus.remote_write コンポーネントを通じて、収集したメトリクスを Prometheus 互換のストレージ（Prometheus 本体はもちろん、Grafana Mimir や Grafana Cloud など）へ転送できる
* 小規模や既存の Prometheus 運用では、Prometheus 本体を直接スクレイプして TSDB に書き込む
* 大規模／複雑環境では、Alloy を前段 Collector として使い、prometheus.scrape で収集 → 演算・フィルター → prometheus.remote_write で Prometheus（または Mimir）に書き込む
* https://grafana.com/docs/alloy/latest/
* https://hub.docker.com/r/grafana/alloy

| ツール        | 役割                                            | 主な用途                                 |
|---------------|-------------------------------------------------|------------------------------------------|
| Prometheus    | TSDB＋スクレイプ＋ルール評価エンジン            | メトリクスの短期保存・アラート           |
| Grafana Alloy | Telemetry Collector＋Prometheus互換パイプライン | 各種データの収集・変換・転送パイプライン |

#### AlloyでPrometheusの設定ファイルをAlloy向けに変換できる

https://grafana.com/docs/alloy/latest/set-up/migrate/from-prometheus/#convert-a-prometheus-configuration

```bash
alloy convert --source-format=prometheus --bypass-errors --output=<OUTPUT_CONFIG_PATH> <INPUT_CONFIG_PATH>
```

`compose.yaml`にて`alloy-convert`として定義済. 以下でprometheus.yamlから変換できる.

```bash
docker compose run alloy-convert
```

### Grafana Loki

* Logging
* Like Prometheus, but for logs
* Loki
    * https://grafana.com/oss/loki/
    * https://grafana.com/docs/loki/latest/
    * https://github.com/grafana/loki
    * https://hub.docker.com/r/grafana/loki
* Promtail
    * https://grafana.com/docs/loki/latest/send-data/promtail/
    * https://hub.docker.com/r/grafana/promtail
    * Loki用エージェント
* https://grafana.com/docs/loki/latest/setup/install/docker/
* Deprecated fields in loki-local-config causing unmarshal error  
  https://github.com/grafana/loki/issues/16990
* https://grafana.com/docs/loki/latest/configure/
* https://grafana.com/docs/loki/latest/reference/loki-http-api/

```bash
docker compose up loki

docker compose up promtail
```

### Grafana Tempo

* Tracing
* https://hub.docker.com/r/grafana/tempo
* https://github.com/grafana/tempo
    * https://github.com/grafana/tempo/tree/main/example/docker-compose
    * https://github.com/grafana/tempo/blob/main/example/docker-compose/local/docker-compose.yaml
* https://grafana.com/oss/tempo/
    * https://grafana.com/docs/tempo/latest/
    * https://grafana.com/docs/tempo/latest/getting-started/
    * https://grafana.com/docs/tempo/latest/configuration/
    * https://grafana.com/docs/tempo/latest/configuration/manifest/
    * https://grafana.com/docs/tempo/latest/setup/linux/#create-a-tempo-configuration-file
    * https://grafana.com/docs/grafana/latest/datasources/tempo/configure-tempo-data-source/#provision-the-data-source
* OpenTelemetry
    * Logs、Metrics、Traces
    * https://opentelemetry.io
        * https://opentelemetry.io/docs/languages/go/
        * https://opentelemetry.io/docs/languages/go/getting-started/
    * https://github.com/open-telemetry/opentelemetry-go
        * https://pkg.go.dev/go.opentelemetry.io/
            * https://pkg.go.dev/go.opentelemetry.io/otel/
            * https://pkg.go.dev/go.opentelemetry.io/otel/exporters/prometheus
            * https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp#example-package
    * https://github.com/open-telemetry/opentelemetry-go-contrib
        * https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/gin-gonic/gin/otelgin
            * https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/example/server.go

### Grafana Pyroscope

* Profiling
* https://hub.docker.com/r/grafana/pyroscope
* https://grafana.com/oss/pyroscope/
* https://grafana.com/docs/pyroscope/latest/
    * https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/go_push/
* https://github.com/grafana/pyroscope
    * https://github.com/grafana/pyroscope/tree/next/examples
        * https://github.com/grafana/pyroscope/tree/next/examples/golang-push

### その他Grafana Stack

* Beyla
    * https://grafana.com/docs/beyla/latest/
* Faro
    * https://grafana.com/docs/grafana-cloud/monitor-applications/frontend-observability/
* k6
    * https://grafana.com/docs/k6/latest/
    * https://hub.docker.com/r/grafana/k6
* xk6
    * https://hub.docker.com/r/grafana/xk6
    * k6の拡張機能開発「xk6」入門  
      https://zenn.dev/moko_poi/articles/72996341dc1665

## 検証



##

* Grizzly
    * https://hub.docker.com/r/grafana/grizzly
    * GrizzlyとGrafonnetで始めるGrafana Dashboards as Code
      https://dasalog.hatenablog.jp/entry/2024/07/16/100252

* Kist
  * https://github.com/grafana/Kost
  * K8s Cost Calculator

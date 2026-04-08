# Changelog

## [v1.2.2](https://github.com/o-ga09/MH-API/compare/v1.2.1...v1.2.2) - 2026-04-08
- Bump golang.org/x/net from 0.17.0 to 0.28.0 by @dependabot[bot] in https://github.com/o-ga09/MH-API/pull/106
- APIのリアーキテクチャ by @o-ga09 in https://github.com/o-ga09/MH-API/pull/108
- Sentryの初期化と設定を追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/107
- ローカル環境のマイグレーションを整備 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/109
- APIレスポンスの再設計 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/110
- BGM検索APIの統合 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/111
- build(deps): bump github.com/go-playground/validator/v10 from 10.20.0 to 10.26.0 by @dependabot[bot] in https://github.com/o-ga09/MH-API/pull/116
- build(deps): bump github.com/swaggo/swag from 1.16.3 to 1.16.4 by @dependabot[bot] in https://github.com/o-ga09/MH-API/pull/115
- ✨ feat: 不要なリソースを削除し、Cloud Runサービスの設定を整理 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/117
- モンスター検索APIのレスポンスに属性フィールドを追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/123
- fix: マイグレーションSQLの修正 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/125
- リクエストコンテキストからDBセッションを取得するように修正 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/124
- build(deps): bump cloud.google.com/go/storage from 1.53.0 to 1.54.0 by @dependabot[bot] in https://github.com/o-ga09/MH-API/pull/121
- パッケージ構成をリファクタリング by @o-ga09 in https://github.com/o-ga09/MH-API/pull/127
- 武器検索APIを実装 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/129
- アイテム一覧取得API (/v1/items) の実装 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/132
- データ整備 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/135
- Add Claude Code GitHub Workflow by @o-ga09 in https://github.com/o-ga09/MH-API/pull/142
- スキル検索APIの実装 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/143
- 防具検索APIの実装 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/144
- MCPサーバのCloudRunの設定を追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/146
- モンハンAPIのMCPサーバーを実装 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/145
- モンスター取得機能を改善し、ページネーションとフィルタリングを追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/147
- MCPサーバーのDB接続を修正 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/148
- モンスター一覧APIに総件数を追加 by @Copilot in https://github.com/o-ga09/MH-API/pull/178
- feat: Dockerfileのgolangバージョンを1.25-bullseyeから1.25.4-alpine3.22に変更🐳 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/181
- Fix panic on nil tribe by @google-labs-jules[bot] in https://github.com/o-ga09/MH-API/pull/184
- feat: Implement MonHun AI Agent using ADK Go with Gemini integration by @Copilot in https://github.com/o-ga09/MH-API/pull/183
- Fix deploy error by @o-ga09 in https://github.com/o-ga09/MH-API/pull/185
- fix: 🐛 エージェントデプロイ用コンテナのポートを8080に修正 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/186
- fix: エージェントのポート設定を修正 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/189
- chore: Go fix by @o-ga09 in https://github.com/o-ga09/MH-API/pull/200
- Instrument OpenTelemetry and remove Sentry by @o-ga09 in https://github.com/o-ga09/MH-API/pull/202
- Integrate Grafana Pyroscope profiling by @o-ga09 in https://github.com/o-ga09/MH-API/pull/204
- feat: OTel Prometheus ExporterによるGrafanaメトリクス観測可能性の追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/205
- feat: CI/CDワークフローを追加し、Terraformのデプロイを統合 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/206
- Init claude code by @o-ga09 in https://github.com/o-ga09/MH-API/pull/207
- mcpサーバーの検索機能強化: モンスターの使用属性と弱点属性での検索を追加 by @Copilot in https://github.com/o-ga09/MH-API/pull/156
- feat: OpenAPI Spec を最新のAPI定義に更新 (#190) by @o-ga09 in https://github.com/o-ga09/MH-API/pull/208
- fix: 属性検索でモンスターデータが返らないバグを修正 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/209
- fix: mcp element search by @o-ga09 in https://github.com/o-ga09/MH-API/pull/210
- fix: リンターエラー時にCIが失敗するよう修正 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/212
- refactor: service層を削除し、controller→repositoryの直接呼び出しに変更 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/214
- feat: UsageElement検索で英語属性名（Fire, Water等）を日本語に正規化 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/215
- feat: tagprによるリリース管理の設定を追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/220
- feat: セキュリティスキャンをCIに追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/221

## [v1.2.1](https://github.com/o-ga09/MH-API/compare/v1.2.0...v1.2.1) - 2024-06-04
- ✨ BGM検索 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/81

## [v1.2.0](https://github.com/o-ga09/MH-API/compare/v1.1.7...v1.2.0) - 2024-04-14
- リファクタリング by @o-ga09 in https://github.com/o-ga09/MH-API/pull/56
- #61 contextを追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/64
- Feature/refactor db structure #60 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/63
- ディレクトリ構造を変更 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/66
- API Spec整理 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/67
- DBヘルスチェック追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/69
- モンスター検索（複数件） #71 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/75
- モンスター検索（1件） by @o-ga09 in https://github.com/o-ga09/MH-API/pull/76
- モンスター人気投票ランキング検索 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/77
- Feature/add test #74 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/78
- モンスターデータ作成バッチ #57 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/79

## [v1.1.7](https://github.com/o-ga09/MH-API/compare/v1.1.6...v1.1.7) - 2023-11-11

## [v1.1.6](https://github.com/o-ga09/MH-API/compare/v1.1.5...v1.1.6) - 2023-11-11
- レスポンスデータ修正  #53 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/54
- 環境変数を削除 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/55

## [v1.1.5](https://github.com/o-ga09/MH-API/compare/v1.1.4...v1.1.5) - 2023-11-06

## [v1.1.4](https://github.com/o-ga09/MH-API/compare/v1.1.3...v1.1.4) - 2023-11-06

## [v1.1.3](https://github.com/o-ga09/MH-API/compare/v1.1.2...v1.1.3) - 2023-11-06

## [v1.1.2](https://github.com/o-ga09/MH-API/compare/v0.0.0...v1.1.2) - 2023-11-05
- シナリオテスト追加（ヘルスチェックのみ） by @o-ga09 in https://github.com/o-ga09/MH-API/pull/46
- add : シナリオテスト by @o-ga09 in https://github.com/o-ga09/MH-API/pull/47
- テスト用GCS削除 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/48
- stg環境追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/49
- コンテナイメージビルドのジョブを追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/50
- コンテナイメージビルドのジョブを追加 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/51

## [v1.1.1](https://github.com/o-ga09/MH-API/compare/v.1.1.0...v1.1.1) - 2023-10-14
- Feature 1 basic api by @o-ga09 in https://github.com/o-ga09/MH-API/pull/10
- to release branch v1.0.0 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/11
- Create LICENSE by @o-ga09 in https://github.com/o-ga09/MH-API/pull/12
- CREATE LICENCE by @o-ga09 in https://github.com/o-ga09/MH-API/pull/13
- v1.1.0 develop merge by @o-ga09 in https://github.com/o-ga09/MH-API/pull/14
- v1.1.0 release merge by @o-ga09 in https://github.com/o-ga09/MH-API/pull/15
- fix (#19) by @o-ga09 in https://github.com/o-ga09/MH-API/pull/20
- Release by @o-ga09 in https://github.com/o-ga09/MH-API/pull/21
- Fix envvariable by @o-ga09 in https://github.com/o-ga09/MH-API/pull/32
- Feature api doc by @o-ga09 in https://github.com/o-ga09/MH-API/pull/34
- Feature addlog by @o-ga09 in https://github.com/o-ga09/MH-API/pull/35
- ビルド用イメージのgoバージョンアップ by @o-ga09 in https://github.com/o-ga09/MH-API/pull/36
- fix : CI/CDデプロイ by @o-ga09 in https://github.com/o-ga09/MH-API/pull/37

## [v1.0.0](https://github.com/o-ga09/MH-API/commits/v1.0.0) - 2023-05-21

## [v0.0.0](https://github.com/o-ga09/MH-API/compare/v1.1.1...v0.0.0) - 2023-10-30
- Fix viewmodel #40 by @o-ga09 in https://github.com/o-ga09/MH-API/pull/41
- CD : apigateway by @o-ga09 in https://github.com/o-ga09/MH-API/pull/42

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## ロール

- あなたは、メガベンチャー企業で働く CTO クラスのエンジニアです
- コードの提案はもちろん、レビューについても CTO の観点からレビューをお願いします
- 日本語で回答すること

## ルール

- GitHubリポジトリ: `o-ga09/MH-API`
- コミット前に `make test` を実行すること

## 使用技術

- Go / Gin / Gorm
- MySQL（Docker Compose で起動）

## アーキテクチャ

[Go Standard layout](https://github.com/golang-standards/project-layout) に則り、クリーンアーキテクチャっぽい構成を採用している。

```
internal/
  DI/          - 依存性注入（各ドメインのDI設定）
  controller/  - HTTPハンドラー（Ginルーター）
  database/mysql/ - リポジトリ実装（GORM）
  domain/      - ドメインモデル・リポジトリインターフェース
  presenter/   - サーバー・ミドルウェア設定
  service/     - ユースケース層
pkg/           - 共通ライブラリ（config, testutil, validator, ptr, uuid など）
cmd/api/       - APIサーバーエントリポイント
cmd/migration/ - マイグレーションエントリポイント
```

データフロー: `controller` → `service` → `domain`（インターフェース） → `database/mysql`（実装）

## コマンド

```bash
make run              # APIサーバー起動
make build            # ビルド
make test             # テスト実行（DB コンテナも起動）
make lint             # go vet + golangci-lint
make generate         # moq によるモック自動生成
make compose-up       # DBコンテナのみ起動
make compose-down     # コンテナ停止・ボリューム削除
make migrate-up       # マイグレーション適用
make migrate-down     # マイグレーション ロールバック
make migrate-new name=<name>  # 新規マイグレーションファイル作成
make seed             # シードデータ投入
```

単一テスト実行:

```bash
go test ./internal/service/monsters/... -run TestGetMonsters
```

## テスト実装規約

- テーブルドリブンテストで実装すること
- テストケースは正常系・異常系を必ず含めること
- `controller` パッケージのテストは**ゴールデンテスト**を用いること（`pkg/testutil/golden.go`）
  - `controller` に限り 400 などの準正常ケースも実装すること
- モックは `moq` ライブラリで自動生成すること（`make generate`）
  - 生成先: 各ドメインの `repository_mock.go`、各サービスの `*_mock.go`
- 外部制約を持つテーブルのテストデータは、関連テーブルのデータも合わせて作成すること
- `pkg` 配下の共通処理を適宜利用すること

### ゴールデンファイルの更新

ゴールデンテストの期待値（`testdata/`以下の`.golden`ファイル）を更新する場合:

```bash
go test ./internal/controller/... -update
```

### モック再生成のタイミング

`domain/` 配下のインターフェース定義を変更した場合は必ず実行すること:

```bash
make generate
```

## マイグレーション運用規約

- DBスキーマ変更時は必ず専用ファイルを作成すること(`make migrate-new name=<name>`)
- 適用後は `make migrate-status` で状態確認
- ロールバックが必要な場合は `make migrate-down`（`docker compose down -v` は**使用しないこと**）
- マイグレーション適用後にシードが必要な場合は `make seed`

## プルリクエスト作成規約

- ベースブランチは `main` に固定
- タイトルとボディは日本語で記述
- タイトルフォーマット: `コミットタイプ: 変更内容の要約`（例: `feat: ドキュメントレビュー承認機能の追加`）

```bash
gh pr create \
  --base main \
  --head "$(git branch --show-current)" \
  --title "feat: 変更内容の要約" \
  --body "## 変更内容
- 変更点1

## 変更の背景・目的
- 背景の説明

## テスト結果
- [ ] ユニットテスト実行済み
- [ ] 動作確認済み"
```

# MonHun AI Agent

ADK Go（Agent Development Kit for Go）を使用したモンスターハンター攻略情報専用のAIエージェントです。

## 概要

このエージェントは、Gemini APIを使用してモンスターハンターのゲーム攻略情報に関する質問に答えるAIアシスタントです。
MH-APIのデータベースに直接アクセスし、以下の情報を提供できます：

- モンスター情報
- 武器情報
- アイテム情報
- スキル情報

## 機能

### 利用可能なツール

1. **get_monsters** - モンスターの一覧を取得
2. **get_monster_by_id** - 特定のモンスターの詳細情報を取得
3. **get_weapons** - 武器の検索
4. **get_items** - アイテムの一覧を取得
5. **get_item_by_id** - 特定のアイテムの詳細情報を取得
6. **get_items_by_monster** - 特定のモンスターから入手できるアイテムを取得
7. **get_skills** - スキルの一覧を取得
8. **get_skill_by_id** - 特定のスキルの詳細情報を取得

## 必要要件

- Go 1.25以上
- Gemini API Key
- MySQLデータベース（MH-APIのデータベース）

## 環境変数

以下の環境変数を設定する必要があります：

```bash
# 必須
GEMINI_API_KEY=your_gemini_api_key_here
DATABASE_URL=user:password@tcp(host:port)/database?charset=utf8&parseTime=True&loc=Local

# オプション
GEMINI_MODEL=gemini-2.0-flash-exp  # デフォルト
AGENT_PORT=8081                     # デフォルト
ENV=dev                             # dev, stage, prod
```

## ローカル開発

### 1. 依存関係のインストール

```bash
go mod download
```

### 2. データベースの起動

```bash
make compose-start
make migrate-up
make seed
```

### 3. エージェントの起動

```bash
# 環境変数を設定
export GEMINI_API_KEY=your_api_key_here
export DATABASE_URL=mh-api:P@ssw0rd@tcp(127.0.0.1:3306)/mh-api?charset=utf8&parseTime=True&loc=Local

# エージェントを起動
go run ./cmd/agent/main.go
```

### 4. エージェントのテスト

エージェントが起動したら、以下のエンドポイントにアクセスできます：

```bash
# ヘルスチェック
curl http://localhost:8081/v1/agent/health

# ADK REST API エンドポイント
# セッション管理、エージェント実行などのエンドポイントが利用可能
# 詳細は ADK Go のドキュメントを参照
```

## ビルド

```bash
# バイナリのビルド
go build -o agent ./cmd/agent

# Dockerイメージのビルド
docker build -t monhun-agent --target deploy-agent .

# Docker Composeで起動（ローカル開発）
docker compose up agent
```

## Cloud Runへのデプロイ

### 1. Dockerイメージのビルドとプッシュ

```bash
# Google Cloud Projectの設定
export PROJECT_ID=your-project-id
export REGION=asia-northeast1

# Artifact Registryにイメージをプッシュ
gcloud builds submit --tag ${REGION}-docker.pkg.dev/${PROJECT_ID}/monhun/agent:latest --target deploy-agent
```

### 2. Cloud Runへのデプロイ

```bash
gcloud run deploy monhun-agent \
  --image ${REGION}-docker.pkg.dev/${PROJECT_ID}/monhun/agent:latest \
  --platform managed \
  --region ${REGION} \
  --allow-unauthenticated \
  --set-env-vars GEMINI_API_KEY=${GEMINI_API_KEY} \
  --set-env-vars DATABASE_URL=${DATABASE_URL} \
  --set-env-vars GEMINI_MODEL=gemini-2.0-flash-exp \
  --set-env-vars AGENT_PORT=8080 \
  --port 8080 \
  --memory 512Mi \
  --cpu 1
```

## API使用例

### ADK REST APIの使用

ADK GoのREST APIハンドラを使用しているため、以下のようなエンドポイントが利用可能です：

1. **セッションの作成・管理**
2. **エージェントとの対話**
3. **セッション履歴の取得**

詳細なAPI仕様については、[ADK Goのドキュメント](https://google.github.io/adk-docs/)を参照してください。

### 基本的な使用例

```bash
# セッションを作成してエージェントと対話
# 具体的なAPIエンドポイントは ADK REST API の仕様に従います
```

## アーキテクチャ

```
cmd/agent/main.go           # エントリーポイント
├── internal/agent/
│   ├── server.go          # ADK REST APIサーバー
│   └── tools.go           # エージェントツールの実装
├── internal/service/      # 既存のMH-APIサービス層を使用
└── pkg/config/            # 設定管理
```

## 技術スタック

- **ADK Go v0.2.0** - Google製のエージェント開発フレームワーク
- **Gemini API** - LLMとして使用
- **Go 1.25** - プログラミング言語
- **MySQL** - データベース
- **GORM** - ORMライブラリ

## トラブルシューティング

### データベース接続エラー

```bash
# データベースが起動しているか確認
docker ps | grep mh-api

# データベース接続文字列を確認
echo $DATABASE_URL
```

### Gemini APIエラー

```bash
# APIキーが正しく設定されているか確認
echo $GEMINI_API_KEY

# APIキーの有効性を確認（Google Cloud Consoleで確認）
```

## ライセンス

MIT License

## 参考リンク

- [ADK Go Documentation](https://google.github.io/adk-docs/)
- [ADK Go Repository](https://github.com/google/adk-go)
- [Gemini API](https://ai.google.dev/)
- [MH-API](https://github.com/o-ga09/MH-API)

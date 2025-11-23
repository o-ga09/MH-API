# ローカル変数の定義
locals {
  # プロジェクト設定
  project_id = var.project_id
  terraform_service_account = "terraform@${local.project_id}.iam.gserviceaccount.com"
  github_repository = "o-ga09/MH-API"  # プロジェクトのカスタムインストラクションから取得

  # リージョン設定
  region = "asia-northeast1"  # 東京リージョン

  # 有効化するAPIサービス
  services = toset([
    "cloudresourcemanager.googleapis.com",
    "run.googleapis.com",
    "iam.googleapis.com",
    "secretmanager.googleapis.com",
    "artifactregistry.googleapis.com",
    "iamcredentials.googleapis.com",
    "serviceusage.googleapis.com",
    "logging.googleapis.com",
  ])

  # Cloud Run設定
  cloud_run_service_name     = "stg-mh-api"  # APIサービスの名前
  cloud_run_mcp_service_name = "stg-mh-mcp"
  cloud_run_agent_service_name = "stg-mh-agent"

  # コンテナイメージ
  container_image     = "asia-northeast1-docker.pkg.dev/${local.project_id}/mh-api/${var.service_name}:${var.image_tag}"
  container_mcp_image = "asia-northeast1-docker.pkg.dev/${local.project_id}/mh-api/mh-agent:${var.image_tag}"
  container_agent_image = "asia-northeast1-docker.pkg.dev/${local.project_id}/mh-api/mh-agent:${var.image_tag}"
}
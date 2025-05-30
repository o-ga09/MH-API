# ローカル変数の定義
locals {
  # プロジェクト設定
  project_id = var.project_id
  terraform_service_account = "terraform@${local.project_id}.iam.gserviceaccount.com"
  cloud_run_invoke_service_account = "${local.project_id}@appspot.gserviceaccount.com"
  github_repository = "o-ga09/MH-API"  # プロジェクトのカスタムインストラクションから取得

  # リージョン設定
  region = "asia-northeast1"  # 東京リージョン

  # 有効化するAPIサービス
  services = toset([
    "cloudrun.googleapis.com",
    "iam.googleapis.com",
    "artifactregistry.googleapis.com",
    "iamcredentials.googleapis.com",
    "serviceusage.googleapis.com"
  ])

  # Cloud Run設定
  cloud_run_service_name = "stg-mh-api"  # APIサービスの名前
  cloud_run_mcp_service_name = "stg-mh-mcp"
  
  # コンテナイメージ
  container_image     = "asia-northeast1-docker.pkg.dev/${local.project_id}/${local.project_id}/${var.service_name}:${var.image_tag}"
  container_mcp_image = "asia-northeast1-docker.pkg.dev/${local.project_id}/${local.project_id}/mh-mcp:${var.image_tag}"
}
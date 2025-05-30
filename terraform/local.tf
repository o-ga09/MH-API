# ローカル変数の定義
locals {
  # プロジェクト設定
  project_id = "mh-api"
  region     = "asia-northeast1"
  
  # GitHub設定
  github_repository = "o-ga09/MH-API"
  
  # サービスアカウント
  terraform_service_account          = "terraform@mh-api.iam.gserviceaccount.com"
  cloud_run_invoke_service_account   = "cloud-run-invoker@mh-api.iam.gserviceaccount.com"
  
  # Cloud Run設定
  cloud_run_service_name     = "stg-mh-api"
  cloud_run_mcp_service_name = "stg-mh-mcp"
  
  # コンテナイメージ
  container_image     = "asia-northeast1-docker.pkg.dev/${local.project_id}/${local.project_id}/${var.service_name}:${var.image_tag}"
  container_mcp_image = "asia-northeast1-docker.pkg.dev/${local.project_id}/${local.project_id}/mh-mcp:${var.image_tag}"

  
  # 有効化するサービスAPI
  services = toset([
    "run.googleapis.com",
    "secretmanager.googleapis.com",
    "cloudbuild.googleapis.com",
    "containerregistry.googleapis.com",
    "iam.googleapis.com",
    "iamcredentials.googleapis.com"
  ])
}
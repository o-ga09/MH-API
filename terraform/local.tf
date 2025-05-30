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
  container_image     = "gcr.io/mh-api/mh-api:latest"
  container_mcp_image = "gcr.io/mh-api/mh-mcp:latest"
  
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
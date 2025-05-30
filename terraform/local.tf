# ローカル変数の定義
locals {
  # Cloud Run設定
  cloud_run_mcp_service_name = "stg-mh-mcp"
  
  # コンテナイメージ
  container_image     = "asia-northeast1-docker.pkg.dev/${local.project_id}/${local.project_id}/${var.service_name}:latest"
  container_mcp_image = "asia-northeast1-docker.pkg.dev/${local.project_id}/${local.project_id}/mh-mcp:latest"
}
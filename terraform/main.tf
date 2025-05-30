# provider 設定
terraform {
  required_providers {
        google  = {
      source  = "hashicorp/google"
      version = ">= 4.0.0"
    }
  }
  required_version = ">= 1.3.0"
  backend "gcs" {
    bucket = "terraform-state-mh-api"
    prefix = "terraform/state"
  }
}

## API の有効化(Workload Identity 用)
resource "google_project_service" "enable_api" {
  for_each                   = local.services
  project                    = local.project_id
  service                    = each.value
  disable_dependent_services = true
}

# Workload Identity Pool 設定
resource "google_iam_workload_identity_pool" "terraform-pool" {
  provider                  = google-beta
  project                   = local.project_id
  workload_identity_pool_id = "terraform-pool"
  display_name              = "terraform-pool"
  description               = "GitHub Actions で使用"
}

# Workload Identity Provider 設定
resource "google_iam_workload_identity_pool_provider" "terraform-provider" {
  provider                           = google-beta
  project                            = local.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.terraform-pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "terraform-provider"
  display_name                       = "terraform-provider"
  description                        = "GitHub Actions で使用"
  
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
  }
  
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}

# GitHub Actions が借用するサービスアカウント
data "google_service_account" "terraform_sa" {
  account_id = local.terraform_service_account
}

# サービスアカウントの IAM Policy 設定と GitHub リポジトリの指定
resource "google_service_account_iam_member" "terraform_sa" {
  service_account_id = data.google_service_account.terraform_sa.id
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.terraform-pool.name}/attribute.repository/${local.github_repository}"
}

# Cloud Run サービスアカウントに必要な権限を付与
resource "google_project_iam_member" "service_account_user" {
  project = local.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${local.terraform_service_account}"
}

resource "google_project_iam_member" "cloud_run_trace_agent" {
  project = local.project_id
  role    = "roles/cloudtrace.agent"
  member  = "serviceAccount:${local.terraform_service_account}"
}

resource "google_project_iam_member" "logging_writer" {
  project = local.project_id
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${local.terraform_service_account}"
}

resource "google_project_iam_member" "secret_manager_secret_accessor" {
  project = local.project_id
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${local.terraform_service_account}"
}

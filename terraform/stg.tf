# GCS作成用
# resource "google_storage_bucket" "test001" {
#     name            = "${local.project_id}-bucket-test001"
#     project         = local.project_id
#     location        = local.region
#     force_destroy   = true
#     uniform_bucket_level_access = true
# }

# Cloud Run作成用
resource "google_cloud_run_service" "stg-mh-api" {
    name            = local.cloud_run_service_name
    project         = local.project_id
    location        = local.region
    autogenerate_revision_name = true
    template {
    spec {
            containers {
                image = local.container_image
                startup_probe {
                    initial_delay_seconds = 0
                    timeout_seconds = 1
                    period_seconds = 3
                    failure_threshold = 1
                    tcp_socket {
                        port = 8080
                    }
                }
            }
            service_account_name = local.cloud_run_invoke_service_account
        }
    }
    traffic {
        percent         = 100
        latest_revision = true
    }
}

data "google_iam_policy" "auth" {
  binding {
    role = "roles/run.invoker"
    members = [
       "serviceAccount:${local.cloud_run_invoke_service_account}",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "auth" {
  location    = google_cloud_run_service.stg-mh-api.location
  project     = google_cloud_run_service.stg-mh-api.project
  service     = google_cloud_run_service.stg-mh-api.name

  policy_data = data.google_iam_policy.auth.policy_data
}

# API Gateway作成用
resource "google_api_gateway_api" "api" {
    provider = google-beta
    project = local.project_id
    api_id = "mh-api"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  project = local.project_id
  api = google_api_gateway_api.api.api_id
  api_config_id = "mh-api"

  openapi_documents {
    document {
      path = "../doc/openapi/apigateway.yml"
      contents = filebase64("../doc/openapi/apigateway.yml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

data "google_iam_policy" "admin" {
  provider = google-beta
  binding {
    role = "roles/apigateway.admin"
    members = [
      "serviceAccount:${local.cloud_run_invoke_service_account}",
    ]
  }
}

resource "google_api_gateway_api_config_iam_policy" "policy" {
  provider = google-beta
  api = google_api_gateway_api_config.api_cfg.api
  api_config = google_api_gateway_api_config.api_cfg.api_config_id
  policy_data = data.google_iam_policy.admin.policy_data
}
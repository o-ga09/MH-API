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
        container_concurrency = 10
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
          env {
            name  = "DATABASE_URL"
            value_from {
              secret_key_ref {
                name = "DATABASE_URL"
                key  = 1
              }
            }
          }
          env {
            name  = "ENV"
            value_from {
              secret_key_ref {
                name = "ENV"
                key  = 1
              }
            }
          }
          env {
            name  = "ALLOW_URL"
            value_from {
              secret_key_ref {
                  name = "ALLOW_URL"
                  key  = 1
              }
            }
          }
          ports {
            container_port = 8080
            name           = "http1"
          }  
        }
        service_account_name = local.cloud_run_invoke_service_account
      }
      metadata {
        annotations = {
          "autoscaling.knative.dev/maxScale"      = "1"
        }
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
# API Gatewayをterraform時管理から除外
# 除外理由
# --backend-auth-service-accountをterraformから設定できない
# 2023.11.5
# resource "google_api_gateway_api" "api" {
#     provider = google-beta
#     project = local.project_id
#     api_id = "mh-api"
# }

# resource "google_api_gateway_api_config" "api_cfg" {
#   provider = google-beta
#   project = local.project_id
#   api = google_api_gateway_api.api.api_id
#   api_config_id = "mh-api"

#   openapi_documents {
#     document {
#       path = "../doc/openapi/apigateway.yml"
#       contents = filebase64("../doc/openapi/apigateway.yml")
#     }
#   }
#   lifecycle {
#     create_before_destroy = true
#   }
# }

# resource "google_api_gateway_gateway" "api_gw" {
#   provider = google-beta
#   project = local.project_id
#   region = local.region
#   api_config = google_api_gateway_api_config.api_cfg.id
#   gateway_id = "mh-api-gateway"
# }

# data "google_iam_policy" "admin" {
#   provider = google-beta
#   binding {
#     role = "roles/apigateway.admin"
#     members = [
#       "serviceAccount:${local.cloud_run_invoke_service_account}",
#     ]
#   }
# }

# resource "google_api_gateway_api_iam_policy" "policy" {
#   provider = google-beta
#   project = google_api_gateway_api.api.project
#   api = google_api_gateway_api.api.api_id
#   policy_data = data.google_iam_policy.admin.policy_data
# }

# resource "google_api_gateway_api_config_iam_policy" "policy" {
#   provider = google-beta
#   api = google_api_gateway_api_config.api_cfg.api
#   api_config = google_api_gateway_api_config.api_cfg.api_config_id
#   policy_data = data.google_iam_policy.admin.policy_data
# }

# resource "google_api_gateway_gateway_iam_policy" "policy" {
#   provider = google-beta
#   project = google_api_gateway_gateway.api_gw.project
#   region = google_api_gateway_gateway.api_gw.region
#   gateway = google_api_gateway_gateway.api_gw.gateway_id
#   policy_data = data.google_iam_policy.admin.policy_data
# }
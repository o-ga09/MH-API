# MCP Cloud Run Service
resource "google_cloud_run_service" "stg-mh-mcp" {
  name                       = local.cloud_run_mcp_service_name
  project                    = local.project_id
  location                   = local.region
  autogenerate_revision_name = true
  template {
    spec {
      container_concurrency = 10
      containers {
        image = local.container_mcp_image
        startup_probe {
          initial_delay_seconds = 0
          timeout_seconds       = 240
          period_seconds        = 240
          failure_threshold     = 1
          tcp_socket {
            port = 8080
          }
        }
        env {
          name = "DATABASE_URL"
          value_from {
            secret_key_ref {
              name = "DATABASE_URL"
              key  = "latest"
            }
          }
        }
        env {
          name = "SENTRY_DSN"
          value_from {
            secret_key_ref {
              name = "SENTRY_DSN"
              key  = "latest"
            }
          }
        }
        env {
          name  = "ENV"
          value = "PROD"
        }
        env {
          name  = "LOG_LEVEL"
          value = "INFO"
        }
        env {
          name  = "GIN_MODE"
          value = "release"
        }
        env {
          name  = "SERVICE_NAME"
          value = "mh-mcp"
        }
        env {
          name  = "PROJECTID"
          value = "mh-api"
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
        "autoscaling.knative.dev/maxScale" = "1"
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

# IAM Policy for MCP Service
data "google_iam_policy" "mcp_auth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "serviceAccount:${local.cloud_run_invoke_service_account}",
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "mcp_auth" {
  location    = google_cloud_run_service.stg-mh-mcp.location
  project     = google_cloud_run_service.stg-mh-mcp.project
  service     = google_cloud_run_service.stg-mh-mcp.name

  policy_data = data.google_iam_policy.mcp_auth.policy_data
}
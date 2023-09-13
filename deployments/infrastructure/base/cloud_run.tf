resource "google_project_service" "run" {
  service = "run.googleapis.com"
}

resource "google_cloud_run_v2_service" "app" {
  name     = "${var.app_name}-${var.env}"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "europe-central2-docker.pkg.dev/just-code-dev/just-code-dev/just-code-dev:${var.image_tag}"

      ports {
        container_port = 8080
      }

      env {
        name  = "BUCKET_NAME"
        value = google_storage_bucket.main.name
      }

      env {
        name  = "FIREBASE_PROJECT_ID"
        value = google_firebase_project.project.display_name
      }

      env {
        name  = "CDN_BASE_URL"
        value = "https://storage.googleapis.com/${var.app_name}-${var.env}-bucket"
      }

      env {
        name  = "APP_PORT"
        value = "8080"
      }

      env {
        name  = "APP_ENV"
        value = var.env
      }

      env {
        name  = "DB_NAME"
        value = google_sql_database.database.name
      }

      env {
        name  = "DB_CONNECTION_NAME"
        value = "/cloudsql/${google_sql_database_instance.instance.connection_name}"
      }

      env {
        name = "DB_USER"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.dbuser.secret_id
            version = "latest"
          }
        }
      }

      env {
        name = "DB_PASSWORD"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.dbpass.secret_id
            version = "latest"
          }
        }
      }

      volume_mounts {
        name       = "cloudsql"
        mount_path = "/cloudsql"
      }
    }

    scaling {
      min_instance_count = 0
      max_instance_count = 1
    }

    volumes {
      name = "cloudsql"
      cloud_sql_instance {
        instances = [google_sql_database_instance.instance.connection_name]
      }
    }
  }

  depends_on = [google_project_service.run]
}

output "app_url" {
  value = google_cloud_run_v2_service.app.uri
}

resource "google_cloud_run_service_iam_binding" "app" {
  location = google_cloud_run_v2_service.app.location
  service  = google_cloud_run_v2_service.app.name
  role     = "roles/run.invoker"

  members = [
    "allUsers"
  ]
}

resource "google_project_service" "run" {
  service = "run.googleapis.com"
}

resource "google_cloud_run_v2_service" "app" {
  name     = "${var.app_name}-${var.env}"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "europe-central2-docker.pkg.dev/just-code-dev/just-code-dev/just-code-dev:latest"

      ports {
        container_port = 8080
      }

      env {
        name  = "PORT"
        value = "8080"
      }
    }

    scaling {
      min_instance_count = 0
      max_instance_count = 1
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

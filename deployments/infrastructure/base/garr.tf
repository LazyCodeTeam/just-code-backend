resource "google_project_service" "garr" {
  service = "artifactregistry.googleapis.com"
}

resource "google_artifact_registry_repository" "app" {
  repository_id = "${var.app_name}-${var.env}"
  format        = "DOCKER"

  depends_on = [
    google_project_service.garr,
  ]
}

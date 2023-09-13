resource "google_storage_bucket" "main" {
  name          = "${var.app_name}-${var.env}-bucket"
  storage_class = "MULTI_REGIONAL"
  location      = var.multiregion
  force_destroy = true
}

resource "google_project_service" "compute" {
  service = "compute.googleapis.com"
}

resource "google_storage_bucket_iam_member" "default" {
  bucket = google_storage_bucket.main.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}

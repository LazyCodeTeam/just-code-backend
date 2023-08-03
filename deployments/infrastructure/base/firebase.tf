data "google_project" "project" {
}

resource "google_firebase_project" "project" {
  provider = google-beta
  project  = data.google_project.project.project_id
}

resource "google_storage_bucket" "main" {
  name          = "${var.app_name}-${var.env}-bucket"
  storage_class = "MULTI_REGIONAL"
  location      = var.multiregion
  force_destroy = true
}

resource "google_project_service" "compute" {
  service = "compute.googleapis.com"
}

resource "google_compute_backend_bucket" "cdn" {
  name        = "${var.app_name}-${var.env}-backend-bucket"
  bucket_name = google_storage_bucket.main.name
  enable_cdn  = true
  cdn_policy {
    cache_mode        = "CACHE_ALL_STATIC"
    client_ttl        = 3600
    default_ttl       = 3600
    max_ttl           = 86400
    negative_caching  = true
    serve_while_stale = 86400
  }
}

resource "google_storage_bucket_iam_member" "default" {
  bucket = google_storage_bucket.main.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}

resource "google_compute_global_address" "cdn" {
  name = "${var.app_name}-${var.env}-ip"
}

resource "google_compute_url_map" "cdn" {
  name            = "${var.app_name}-${var.env}-url-map"
  default_service = google_compute_backend_bucket.cdn.id
}

resource "google_compute_target_http_proxy" "cdn" {
  name    = "${var.app_name}-${var.env}-proxy"
  url_map = google_compute_url_map.cdn.id
}

resource "google_compute_global_forwarding_rule" "cdn" {
  name                  = "${var.app_name}-${var.env}-forwarding-rule"
  ip_protocol           = "TCP"
  load_balancing_scheme = "EXTERNAL"
  port_range            = "80"
  target                = google_compute_target_http_proxy.cdn.id
  ip_address            = google_compute_global_address.cdn.id
}

output "cdn_ip" {
  value = google_compute_global_address.cdn.address
}

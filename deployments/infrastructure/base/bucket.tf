resource "google_storage_bucket" "main" {
  name          = "${var.app_name}-${var.env}-bucket"
  storage_class = "MULTI_REGIONAL"
  location      = var.multiregion
}

resource "google_sql_database_instance" "instance" {
  name             = "${var.app_name}-${var.env}-db"
  region           = var.region
  database_version = "POSTGRES_15"
  settings {
    tier = "db-f1-micro"
  }

  deletion_protection = "false"
}

resource "google_sql_database" "database" {
  name     = "just-code"
  instance = google_sql_database_instance.instance.name
}

resource "google_sql_database_instance" "instance" {
  name             = "${var.app_name}-${var.env}-db"
  region           = var.region
  database_version = "POSTGRES_15"

  settings {
    tier = "db-f1-micro"
    insights_config {
      query_insights_enabled = true
    }
  }

  deletion_protection = "false"
}

resource "google_sql_database" "database" {
  name     = "just-code"
  instance = google_sql_database_instance.instance.name
}

output "db_connection_name" {
  value     = google_sql_database_instance.instance.connection_name
  sensitive = true
}

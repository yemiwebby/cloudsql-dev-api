provider "google" {
  project = var.project_id
  region  = var.region
  credentials = var.gcp_credentials_json
}

resource "google_sql_database_instance" "postgres" {
  name             = var.instance_name
  database_version = "POSTGRES_13"
  region           = var.region

  lifecycle {
    prevent_destroy = true
  }

  settings {
    tier = "db-f1-micro"

    ip_configuration {
      ipv4_enabled    = true
      authorized_networks {
        name  = "allow-all"
        value = "0.0.0.0/0"
      }
    }

    backup_configuration {
      enabled = false
    }

    deletion_protection_enabled = false
  }
}

resource "google_sql_user" "postgres_user" {
  name     = var.db_user
  instance = google_sql_database_instance.postgres.name
  password = var.db_pass
}

resource "google_sql_database" "default" {
  name     = var.db_name
  instance = google_sql_database_instance.postgres.name
}

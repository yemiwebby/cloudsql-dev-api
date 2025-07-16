output "db_instance_connection_name" {
  value = google_sql_database_instance.postgres.connection_name
}

output "db_public_ip" {
  value = google_sql_database_instance.postgres.public_ip_address
}

output "db_name" {
  value = var.db_name
}

output "db_user" {
  value = var.db_user
}

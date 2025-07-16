variable "project_id" {}
variable "region" {
  default = "europe-west1"
}
variable "instance_name" {}
variable "db_user" {
  default = "postgres"
}
variable "db_pass" {}
variable "db_name" {
  default = "devdb"
}

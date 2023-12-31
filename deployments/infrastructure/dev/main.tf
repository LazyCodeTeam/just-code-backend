terraform {
  backend "gcs" {
    bucket = "just-code-dev-tfstate"
    prefix = "terraform/state"
  }
}

locals {
  region        = "europe-central2"
  multiregion   = "EU"
  env           = "dev"
  app_name      = "just-code"
  full_app_name = "${local.app_name}-${local.env}"
}

provider "google" {
  project = local.full_app_name
  region  = local.region
}

provider "google-beta" {
  project = local.full_app_name
  region  = local.region
}

module "app" {
  source = "../base"

  env         = local.env
  app_name    = local.app_name
  region      = local.region
  multiregion = local.multiregion
  image_tag   = var.image_tag
}

output "all_outputs" {
  value     = module.app
  sensitive = true
}

output "db_connection_name" {
  value     = module.app.db_connection_name
  sensitive = true
}

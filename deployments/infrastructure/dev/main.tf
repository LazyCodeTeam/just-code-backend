terraform {
  backend "gcs" {
    bucket = "just-code-dev-tfstate"
    prefix = "terraform/state"
  }
}

provider "google" {
  project = "just-code-dev"
  region  = "europe-central2"
}

module "app" {
  source = "../base"

  env       = "dev"
  app_name  = "just-code"
  region    = "europe-central2"
  image_tag = var.image_tag
}

output "all_outputs" {
  value     = module.app
  sensitive = true
}

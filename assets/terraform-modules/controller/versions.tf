# Terraform version and plugin versions

terraform {
  required_version = ">= 0.13"

  required_providers {
    ct = {
      source  = "poseidon/ct"
      version = "0.6.1"
    }
    template = {
      source  = "hashicorp/template"
      version = "2.1.2"
    }
  }
}

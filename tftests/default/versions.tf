terraform {
  required_providers {
    random = {
      source  = "registry.terraform.io/hashicorp/random"
      version = ">= 3"
    }
    observe = {
      source  = "terraform.observeinc.com/observeinc/observe"
      version = "~> 0.13"
    }
  }
  required_version = ">= 1.0"
}

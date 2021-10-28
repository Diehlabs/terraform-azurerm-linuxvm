terraform {
  required_version = ">= 0.15.0"
  required_providers {
    azurerm = {
      version = ">= 2.77.0"
      source  = "hashicorp/azurerm"
    }
  }
}

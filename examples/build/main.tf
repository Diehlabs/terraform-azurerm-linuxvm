provider "azurerm" {
  features {}
}

module "linuxvm" {
  source    = "../.."
  tags      = local.tags
  rg_name   = azurerm_resource_group.terratest.name
  subnet_id = azurerm_subnet.terratest.id
  vm_name   = local.vm_name
  ssh_key   = tls_private_key.terratest.public_key_openssh
  nsg_rules = local.nsg_rules
}

resource "azurerm_resource_group" "terratest" {
  name     = "terratest-${var.gh_run_id}"
  location = local.tags.location
}

resource "azurerm_virtual_network" "terratest" {
  name                = "virtualNetwork1"
  location            = azurerm_resource_group.terratest.location
  resource_group_name = azurerm_resource_group.terratest.name
  address_space       = ["172.16.13.0/24"]
  tags                = local.tags
}

resource "azurerm_subnet" "terratest" {
  name                 = "subnet-terratest"
  resource_group_name  = azurerm_resource_group.terratest.name
  virtual_network_name = azurerm_virtual_network.terratest.name
  address_prefixes     = ["172.16.13.0/24"]
}

resource "tls_private_key" "terratest" {
  algorithm = "RSA"
  rsa_bits  = "2048"
}

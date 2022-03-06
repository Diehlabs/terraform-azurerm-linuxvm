locals {
  tags = {
    test_for   = var.test_for
    unique_id  = var.unique_id
    created_by = "terratest"
    location   = "centralus"
  }

  vm_name = "${local.tags.created_by}-${var.unique_id}"

  nsg_rules = {
    HTTPS = {
      priority                   = 1100
      direction                  = "Inbound"
      access                     = "Allow"
      protocol                   = "Tcp"
      source_port_range          = "*"
      destination_port_range     = "443"
      source_address_prefix      = "*"
      destination_address_prefix = "*"
    },
    SSH = {
      priority                   = 1001
      direction                  = "Inbound"
      access                     = "Allow"
      protocol                   = "Tcp"
      source_port_range          = "*"
      destination_port_range     = "22"
      source_address_prefix      = "*"
      destination_address_prefix = "*"
    }
  }
}

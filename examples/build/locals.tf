locals {
  tags = {
    test_for   = var.gh_repo
    gh_run_id  = var.gh_run_id
    created_by = "terratest"
    location   = "centralus"
  }
  # vm_name = replace("${local.tags.created_by}-${var.gh_run_id}-${var.gh_repo}", "/", "-")
  vm_name = "${local.tags.created_by}-${var.gh_run_id}"
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
    }
  }
}

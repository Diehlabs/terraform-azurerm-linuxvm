resource "azurerm_network_security_group" "vm" {
  name                = "${azurerm_linux_virtual_machine.vm.name}-nsg"
  location            = var.tags.location
  resource_group_name = var.rg_name
}

resource "azurerm_network_interface_security_group_association" "vm" {
  network_interface_id      = azurerm_network_interface.vm.id
  network_security_group_id = azurerm_network_security_group.vm.id
}

resource "azurerm_network_security_rule" "vm" {
  for_each                    = var.nsg_rules
  name                        = each.key
  priority                    = each.value.priority
  direction                   = each.value.direction
  access                      = each.value.access
  protocol                    = each.value.protocol
  source_port_range           = each.value.source_port_range
  destination_port_range      = each.value.destination_port_range
  source_address_prefix       = each.value.source_address_prefix
  destination_address_prefix  = each.value.destination_address_prefix
  resource_group_name         = azurerm_network_security_group.vm.resource_group_name
  network_security_group_name = azurerm_network_security_group.vm.name
}

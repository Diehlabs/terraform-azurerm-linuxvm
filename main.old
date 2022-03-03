resource "azurerm_network_interface" "vm" {
  name                = "${var.vm_name}-nic"
  location            = var.tags.location
  resource_group_name = var.rg_name
  ip_configuration {
    name                          = "default"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.vm_pub_ip.id
  }
  tags = var.tags
}

resource "azurerm_linux_virtual_machine" "vm" {
  availability_set_id             = var.availability_set_id
  name                            = var.vm_name
  location                        = var.tags.location
  resource_group_name             = var.rg_name
  size                            = var.size
  admin_username                  = "adminuser"
  disable_password_authentication = true

  network_interface_ids = [
    azurerm_network_interface.vm.id,
  ]

  admin_ssh_key {
    username   = "adminuser"
    public_key = var.ssh_key
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  identity {
    type         = "UserAssigned"
    identity_ids = var.identity_ids
  }

  tags = var.tags
}

resource "azurerm_public_ip" "vm_pub_ip" {
  name                = "${var.vm_name}-pubip"
  location            = var.tags.location
  resource_group_name = var.rg_name
  allocation_method   = "Static"
  sku                 = "Standard"
  tags                = var.tags
}

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

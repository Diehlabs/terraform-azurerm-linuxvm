resource "azurerm_public_ip" "vm_pub_ip" {
  count               = var.use_public_ip ? 1 : 0
  name                = "${var.vm_name}-pubip"
  location            = var.tags.location
  resource_group_name = var.rg_name
  allocation_method   = "Static"
  sku                 = "Standard"
  tags                = var.tags
}

resource "azurerm_network_interface" "vm" {
  name                = "${var.vm_name}-nic"
  location            = var.tags.location
  resource_group_name = var.rg_name
  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
    primary                       = true
  }

  dynamic "ip_configuration" {
    for_each = var.use_public_ip ? { public_ip = azurerm_public_ip.vm_pub_ip[0] } : {}
    content {
      name                          = "public"
      subnet_id                     = var.subnet_id
      private_ip_address_allocation = "Dynamic"
      public_ip_address_id          = ip_configuration.value.id
    }
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

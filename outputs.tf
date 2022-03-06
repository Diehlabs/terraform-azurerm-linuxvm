output "ip_addresses" {
  value = azurerm_network_interface.vm.private_ip_address
}

output "public_ip" {
  value = var.use_public_ip ? azurerm_public_ip.vm_pub_ip[0].ip_address : null
}

output "vm_id" {
  value = azurerm_linux_virtual_machine.vm.id
}

output "vm_name" {
  value = azurerm_linux_virtual_machine.vm.name
}

output "computer_name" {
  value = azurerm_linux_virtual_machine.vm.computer_name
}

output "nic_id" {
  value = azurerm_network_interface.vm.id
}

output "msi" {
  value = azurerm_linux_virtual_machine.vm.identity
}

output "vm_size" {
  value = azurerm_linux_virtual_machine.vm.size
}

output "nsg" {
  value = {
    name = azurerm_network_security_group.vm.name
    rules = { for name, rule in var.nsg_rules :
      name => azurerm_network_security_rule.vm[name]
    }
  }
}

output "ip_addresses" {
  value = azurerm_network_interface.vm.private_ip_address
}

output "public_ip" {
  value = azurerm_public_ip.vm_pub_ip.ip_address
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

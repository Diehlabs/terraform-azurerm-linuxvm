output "ip_addresses" {
  value = azurerm_network_interface.vm.private_ip_address
}

output "public_ip" {
  value = azurerm_public_ip.vm_pub_ip.ip_address
}


output "vm_name" {
  value = "test"
}

output "computer_name" {
  value = "test01"
}

output "nic_id" {
  value = azurerm_network_interface.vm.id
}

output "vm_size" {
  value = "size"
}

output "nsg" {
  value = {
    name = azurerm_network_security_group.vm.name
    rules = { for name, rule in var.nsg_rules :
      name => azurerm_network_security_rule.vm[name]
    }
  }
}

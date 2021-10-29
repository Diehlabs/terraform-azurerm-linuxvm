output "ssh_key" {
  value     = tls_private_key.terratest.private_key_pem
  sensitive = true
}

output "resource_group_name" {
  value = azurerm_resource_group.terratest.name
}

output "nsg_rules" {
  value = module.linuxvm.nsg_rules
}

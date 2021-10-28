# terraform-azurerm-linuxvm

Creates an Azure Linux VM with:
* Public IP
* NSG with rule for SSH
* System assigned MSI

TODO:
* Update TF variables
* Update Terratests

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.15.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | >= 2.77.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | 2.82.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_linux_virtual_machine.vm](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/linux_virtual_machine) | resource |
| [azurerm_network_interface.vm](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface) | resource |
| [azurerm_network_interface_security_group_association.vm](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface_security_group_association) | resource |
| [azurerm_network_security_group.vm](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_security_group) | resource |
| [azurerm_public_ip.vm_pub_ip](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/public_ip) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_availability_set_id"></a> [availability\_set\_id](#input\_availability\_set\_id) | Optional availability set to add the VM to | `any` | `null` | no |
| <a name="input_nsg_rules"></a> [nsg\_rules](#input\_nsg\_rules) | Additional NSG rules to add to the VM network interface | `map(any)` | `{}` | no |
| <a name="input_rg_name"></a> [rg\_name](#input\_rg\_name) | n/a | `any` | n/a | yes |
| <a name="input_ssh_key"></a> [ssh\_key](#input\_ssh\_key) | n/a | <pre>object({<br>    public_key_openssh = string<br>    private_key_pem    = string<br>  })</pre> | n/a | yes |
| <a name="input_subnet_id"></a> [subnet\_id](#input\_subnet\_id) | n/a | `any` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `any` | n/a | yes |
| <a name="input_vm_name"></a> [vm\_name](#input\_vm\_name) | n/a | `any` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_computer_name"></a> [computer\_name](#output\_computer\_name) | n/a |
| <a name="output_ip_addresses"></a> [ip\_addresses](#output\_ip\_addresses) | n/a |
| <a name="output_msi"></a> [msi](#output\_msi) | n/a |
| <a name="output_nic_id"></a> [nic\_id](#output\_nic\_id) | n/a |
| <a name="output_public_ip"></a> [public\_ip](#output\_public\_ip) | n/a |
| <a name="output_vm_name"></a> [vm\_name](#output\_vm\_name) | n/a |
<!-- END_TF_DOCS -->

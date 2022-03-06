variable "tags" {}

variable "rg_name" {}

variable "subnet_id" {}

variable "vm_name" {}

variable "ssh_key" {
  description = "PEM formatted private ssh key"
  type        = string
}

variable "availability_set_id" {
  description = "Optional availability set to add the VM to"
  default     = null
}

variable "nsg_rules" {
  description = "Additional NSG rules to add to the VM network interface"
  type        = map(any)
  default     = {}
}

variable "size" {
  default = "Standard_B1ls"
}

variable "identity_ids" {
  description = "List of identities to assign to the VM"
  type        = list(string)
}


variable "use_public_ip" {
  description = "Set to true to assign a public IP to the VM NIC, defaults to false."
  type        = bool
  default     = false
}

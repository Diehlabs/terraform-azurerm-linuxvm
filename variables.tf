variable "tags" {
  description = "Tags to apply to resources; must include a location key, which sets the Azure region"
  type        = map(string)

  validation {
    condition     = contains(keys(var.tags), "location")
    error_message = "The tags map must include a location key."
  }
}

variable "rg_name" {
  description = "Name of the resource group to create resources in"
  type        = string
}

variable "subnet_id" {
  description = "ID of the subnet for the VM network interface"
  type        = string
}

variable "vm_name" {
  description = "Name of the VM; also used as a prefix for related resources"
  type        = string
}

variable "ssh_key" {
  description = "OpenSSH formatted public ssh key for the admin user"
  type        = string
}

variable "availability_set_id" {
  description = "Optional availability set to add the VM to"
  type        = string
  default     = null
}

variable "nsg_rules" {
  description = "Additional NSG rules to add to the VM network interface"
  type        = map(any)
  default     = {}
}

variable "size" {
  description = "Azure VM size"
  type        = string
  default     = "Standard_B1ls"
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

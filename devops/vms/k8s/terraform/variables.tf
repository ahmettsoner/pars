variable "ssh_public_key_filename" {
  type    = string
  default = "infra_id_rsa"
}

variable "vm_pool" {
  type    = string
  default = "k8s_pool"
}

variable "vm_user" {
  type    = string
  default = "ahmetsoner"
}

variable "vm_memory" {
  type    = number
  default = 1024
}

variable "vm_vcpu" {
  type    = number
  default = 1
}

variable "vm_network_name" {
  description = "Libvirt network name"
  type        = string
  default     = "terraform-net2"
}
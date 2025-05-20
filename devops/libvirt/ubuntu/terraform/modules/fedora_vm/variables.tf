variable "vm_count" {}
variable "base_volume_id" {}
variable "ssh_authorized_key" {}
variable "username" {}
variable "memory" { default = 1024 }
variable "vcpu" { default = 1 }
variable "network_name" { default = "default" }

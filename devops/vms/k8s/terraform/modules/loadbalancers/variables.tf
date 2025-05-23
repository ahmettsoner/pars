variable "vm_pool" {}
variable "base_volume_id" {}
variable "ssh_authorized_key" {}
variable "ssh_private_key_path" {}
variable "username" {}
variable "memory" { default = 1024 }
variable "vcpu" { default = 1 }
variable "network_name" { default = "default" }


variable "module_name" { default = "loadbalancer" }

variable "vm_name" { default = "k8s-loadbalancer-vm" }
variable "vm_hostname" { default = "loadbalancer-vm" }
variable "vm_count" { default = 2}
variable "vm_ip_suffix" { default = "192.168.124.xxx" }
variable "vm_ip_block" { default = "24" }
variable "automation_username" { default = "automation" }

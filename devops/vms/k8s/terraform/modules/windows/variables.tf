variable "vm_pool" {}
variable "ssh_authorized_key" {}
variable "ssh_private_key_path" {}
variable "username" {}
variable "memory" { default = 1024 }
variable "vcpu" { default = 1 }
variable "network_name" { default = "default" }


variable "module_name" { default = "windows" }

variable "vm_name" { default = "k8s-windows-vm" }
variable "vm_hostname" { default = "windows-vm" }
variable "vm_count" { default = 1 }
variable "vm_ip_suffix" { default = "192.168.200.xxx" }
variable "vm_ip_block" { default = "24" }
variable "automation_username" { default = "automation" }



variable "windows_iso_path" {
  type        = string
  description = "Path to Windows Server ISO file"
}

variable "virtio_iso_path" {
  type        = string
  description = "Path to VirtIO drivers ISO"
}


variable "network_id" {
  type        = string
  description = "Libvirt network ID"
}
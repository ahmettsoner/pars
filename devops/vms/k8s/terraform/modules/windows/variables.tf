variable "vm_name" {
  type        = string
  description = "Name of the Windows Server VM"
}

variable "memory" {
  type        = number
  default     = 4096
}

variable "vcpus" {
  type        = number
  default     = 2
}

variable "windows_iso_path" {
  type        = string
  description = "Path to Windows Server ISO file"
}

variable "virtio_iso_path" {
  type        = string
  description = "Path to VirtIO drivers ISO"
}

variable "pool_id" {
  type        = string
  description = "Libvirt storage pool to use"
}

variable "network_id" {
  type        = string
  description = "Libvirt network ID"
}
variable "autounattend_iso_path" {
  type        = string
  description = "Path to ISO file containing Autounattend.xml"
}

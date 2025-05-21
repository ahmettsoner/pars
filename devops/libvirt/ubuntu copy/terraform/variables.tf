variable "ssh_public_key_path" {
  type    = string
  default = "~/.ssh/id_rsa"
}

variable "vm_pool" {
  type    = string
  default = "custom_domain2"
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

variable "haproxy_memory" {
  type    = number
  default = 1024
}

variable "haproxy_vcpu" {
  type    = number
  default = 1
}

variable "network_name" {
  type    = string
  default = "default"
}

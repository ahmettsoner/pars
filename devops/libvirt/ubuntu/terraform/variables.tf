variable "ssh_public_key_path" {
  type    = string
  default = "/home/ahmetsoner/.ssh/id_rsa.pub"
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

output "master_name" {
  value = libvirt_domain.master_vms[*].name
}

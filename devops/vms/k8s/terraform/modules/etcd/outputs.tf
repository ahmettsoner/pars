output "etcd_name" {
  value = libvirt_domain.etcd_vms[*].name
}

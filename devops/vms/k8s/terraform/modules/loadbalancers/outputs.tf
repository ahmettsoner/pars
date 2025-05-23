output "loadbalancer_name" {
  value = libvirt_domain.loadbalancer_vms[*].name
}

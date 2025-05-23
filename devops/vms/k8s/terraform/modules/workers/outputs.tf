output "worker_name" {
  value = libvirt_domain.worker_vms[*].name
}

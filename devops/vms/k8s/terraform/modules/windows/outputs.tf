output "windows_vm_name" {
  value = libvirt_domain.windows_vm.name
}
output "vm_ip" {
  value = libvirt_domain.windows_vm.network_interface[0]
}
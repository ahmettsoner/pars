output "windows_vm_name" {
  value = libvirt_domain.windows_vms[*].name
}
# output "vm_ips" {
#   value = [
#     for vm in libvirt_domain.windows_vms : vm.network_interface[0].addresses[0]
#   ]
# }

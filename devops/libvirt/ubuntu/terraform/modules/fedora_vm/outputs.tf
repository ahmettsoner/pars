output "vm_mac_addresses" {
  value = [for vm in libvirt_domain.fedora_vms : vm.network_interface.0.mac]
}
output "vm_ips" {
  value = libvirt_domain.fedora_vms[*].network_interface[0].addresses
}

output "vm_names" {
  value = libvirt_domain.fedora_vms[*].name
}
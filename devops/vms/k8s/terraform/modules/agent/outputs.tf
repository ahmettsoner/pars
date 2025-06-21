output "agent_vm_name" {
  value = libvirt_domain.agent_vms[*].name
}
# output "vm_ips" {
#   value = [
#     for vm in libvirt_domain.agent_vms : vm.network_interface[0].addresses[0]
#   ]
# }

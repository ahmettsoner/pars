# output "loadbalancer_ip" {
#   value = module.loadbalancer_vm.loadbalancer_ip
# }

# output "master_ips" {
#   value = module.master_vms.vm_ips
# }

# output "worker_ips" {
#   value = module.worker_vms.vm_ips
# }


output "vm_network_name" {
  value = libvirt_network.example_network.name
}
# output "windows_vm_ips" {
#   value = module.windows_vm.vm_ips
# }
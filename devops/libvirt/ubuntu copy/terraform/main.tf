provider "libvirt" {
  uri = "qemu:///system"
}

resource "libvirt_pool" "custom_pool" {
  name = var.vm_pool
  type = "dir"
  path = "/var/lib/libvirt/pools/terraform_pool"
}



# Base image volume (paylaşılan)
resource "libvirt_volume" "base_volume" {
  name   = "fedora41-base"
  pool   = libvirt_pool.custom_pool.name
  source = "../../../images/Fedora-Cloud-Base-Generic-41-1.4.x86_64.qcow2"
  format = "qcow2"
}
# Fedora VM modülünü çağırıyoruz
module "fedora_vms" {
  source            = "./modules/fedora_vm"
  vm_pool           = libvirt_pool.custom_pool.name
  vm_count          = 3
  base_volume_id    = libvirt_volume.base_volume.id
  ssh_authorized_key = file(var.ssh_public_key_path)
  username          = var.vm_user
  memory            = var.vm_memory
  vcpu              = var.vm_vcpu
  network_name      = var.network_name
}

# HAProxy VM modülünü çağırıyoruz
module "haproxy_vm" {
  source            = "./modules/haproxy"
  vm_pool           = libvirt_pool.custom_pool.name
  base_volume_id    = libvirt_volume.base_volume.id
  ssh_authorized_key = file(var.ssh_public_key_path)
  username          = var.vm_user
  memory            = var.haproxy_memory
  vcpu              = var.haproxy_vcpu
  network_name      = var.network_name
}

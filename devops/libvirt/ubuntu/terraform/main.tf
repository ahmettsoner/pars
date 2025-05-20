provider "libvirt" {
  uri = "qemu:///system"
}

# Base image volume (paylaşılan)
resource "libvirt_volume" "base_volume" {
  name   = "fedora41-base"
  pool   = "default"
  source = "../../../images/Fedora-Cloud-Base-Generic-41-1.4.x86_64.qcow2"
  format = "qcow2"
}

# Fedora VM modülünü çağırıyoruz
module "fedora_vms" {
  source            = "./modules/fedora_vm"
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
  base_volume_id    = libvirt_volume.base_volume.id
  ssh_authorized_key = file(var.ssh_public_key_path)
  username          = var.vm_user
  memory            = var.haproxy_memory
  vcpu              = var.haproxy_vcpu
  network_name      = var.network_name
}

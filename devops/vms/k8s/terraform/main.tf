provider "libvirt" {
  uri = "qemu:///system"
}

resource "libvirt_pool" "custom_pool" {
  name = var.vm_pool
  type = "dir"
  target {
    path = "/home/ahmettsoner/pools/k8s_pool2"
  }
}



# Base image volume (paylaşılan)
resource "libvirt_volume" "base_volume" {
  name   = "ubuntu-minimal-base"
  pool   = libvirt_pool.custom_pool.name
  source = "../../../images/jammy-server-cloudimg-amd64.qcow2"
  format = "qcow2"
}



module "windows_vm" {
  source                = "./modules/windows"
  vm_name               = "windows-server-2022"
  memory                = 8192
  vcpu                  = 4
  vm_pool               = libvirt_pool.custom_pool.name
  network_id            = libvirt_network.example_network.id
  network_name          = libvirt_network.example_network.name
  ssh_private_key_path  = local_file.private_key_pem.filename
  ssh_authorized_key    = local_file.public_key_openssh.content
  username              = var.vm_user
  windows_iso_path =    "/home/ahmetsoner/AS/prs/pars/devops/images/SERVER_EVAL_x64FRE_en-us.iso"
  virtio_iso_path  =    "/home/ahmetsoner/AS/prs/pars/devops/images/virtio-win.iso"
}

# Masters modülünü çağırıyoruz
# module "etcd_vm" {
#   source                = "./modules/etcd"
#   vm_pool               = libvirt_pool.custom_pool.name
#   base_volume_id        = libvirt_volume.base_volume.id
#   ssh_private_key_path  = local_file.private_key_pem.filename
#   ssh_authorized_key    = local_file.public_key_openssh.content
#   username              = var.vm_user
#   memory                = 1024
#   vcpu                  = 1
#   network_name          = var.network_name
# }

# # HAProxy VM modülünü çağırıyoruz
# module "loadbalancer_vm" {
#   source                = "./modules/loadbalancers"
#   vm_pool               = libvirt_pool.custom_pool.name
#   base_volume_id        = libvirt_volume.base_volume.id
#   ssh_private_key_path  = local_file.private_key_pem.filename
#   ssh_authorized_key    = local_file.public_key_openssh.content
#   username              = var.vm_user
#   memory                = 1024
#   vcpu                  = 1
#   network_name          = var.network_name
# }

# # Masters modülünü çağırıyoruz
# module "master_vm" {
#   source                = "./modules/masters"
#   vm_pool               = libvirt_pool.custom_pool.name
#   base_volume_id        = libvirt_volume.base_volume.id
#   ssh_private_key_path  = local_file.private_key_pem.filename
#   ssh_authorized_key    = local_file.public_key_openssh.content
#   username              = var.vm_user
#   memory                = 1024
#   vcpu                  = 1
#   network_name          = var.network_name
# }

# # # Workers modülünü çağırıyoruz
# module "worker_vms" {
#   source                = "./modules/workers"
#   vm_pool               = libvirt_pool.custom_pool.name
#   base_volume_id        = libvirt_volume.base_volume.id
#   ssh_private_key_path  = local_file.private_key_pem.filename
#   ssh_authorized_key    = local_file.public_key_openssh.content
#   username              = var.vm_user
#   memory                = 1024
#   vcpu                  = 1
#   network_name          = var.network_name
# }


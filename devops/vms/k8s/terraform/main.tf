provider "libvirt" {
  uri = "qemu:///session"
}

resource "libvirt_pool" "custom_pool" {
  name = var.vm_pool
  type = "dir"
  # path = "/var/lib/libvirt/pools/k8s_pool"
  path = "/home/ahmettsoner/pools/k8s_pool2"
}



# Base image volume (paylaşılan)
resource "libvirt_volume" "base_volume" {
  name   = "ubuntu-minimal-base"
  pool   = libvirt_pool.custom_pool.name
  source = "../../../images/jammy-server-cloudimg-amd64.qcow2"
  format = "qcow2"
}


resource "tls_private_key" "private_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "local_file" "private_key_pem" {
  filename = "${path.root}/files/ssh_keys/${var.ssh_public_key_filename}"
  content  = tls_private_key.private_key.private_key_pem
  file_permission = "0600"
}

resource "local_file" "public_key_openssh" {
  filename = "${path.root}/files/ssh_keys/${var.ssh_public_key_filename}.pub"
  content  = tls_private_key.private_key.public_key_openssh
}



# Masters modülünü çağırıyoruz
module "etcd_vm" {
  source                = "./modules/etcd"
  vm_pool               = libvirt_pool.custom_pool.name
  base_volume_id        = libvirt_volume.base_volume.id
  ssh_private_key_path  = local_file.private_key_pem.filename
  ssh_authorized_key    = local_file.public_key_openssh.content
  username              = var.vm_user
  memory                = 1024
  vcpu                  = 1
  network_name          = var.network_name
}

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


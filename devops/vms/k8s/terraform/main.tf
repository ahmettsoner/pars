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

resource "libvirt_network" "example_network" {
  name      = "terraform-network"
  bridge    = "br-terraform"
  mode      = "nat"
  domain    = "terraform.local"
  addresses = ["192.168.200.0/24"]

  dhcp {
    enabled = true
  }

  autostart = true
}

module "windows_vm" {
  source           = "./modules/windows"
  vm_name          = "windows-server-2022"
  memory           = 8192
  vcpus            = 4
  pool_id          = libvirt_pool.custom_pool.name
  network_id       = libvirt_network.example_network.id
  windows_iso_path = "/home/ahmettsoner/AS/prj/pars/devops/images/SERVER_EVAL_x64FRE_en-us.iso"
  virtio_iso_path  = "/home/ahmettsoner/AS/prj/pars/devops/images/virtio-win.iso"
  autounattend_iso_path = "/home/ahmettsoner/AS/prj/pars/devops/images/autounattend.iso"
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



resource "libvirt_volume" "win_disk" {
  name   = "${var.vm_name}.qcow2"
  pool   = var.vm_pool
  format = "qcow2"
  size   = 50 * 1024 * 1024 * 1024 # 50 GB
}

data "template_file" "autounattend" {
  template = file("${path.module}/autounattend/Autounattend.tpl")

  vars = {
    ssh_authorized_key = var.ssh_authorized_key
    username           = var.automation_username
  }
}
resource "local_file" "autounattend_xml" {
  content  = data.template_file.autounattend.rendered
  filename = "${path.module}/autounattend/Autounattend.xml"
}

resource "null_resource" "build_autounattend_iso" {
  depends_on = [local_file.autounattend_xml]

  provisioner "local-exec" {
    command = <<EOT
      ./${path.module}/scripts/create-autounattend-iso.sh ./${path.module}/autounattend ./${path.module}/images/autounattend.iso
    EOT
  }
}

resource "libvirt_domain" "windows_vm" {
  depends_on = [null_resource.build_autounattend_iso]
  name   = var.vm_name
  memory = var.memory
  vcpu   = var.vcpu


  disk {
    volume_id = libvirt_volume.win_disk.id
  }

  disk {
    file = var.windows_iso_path
  }

  disk {
    file = abspath("${path.module}/images/autounattend.iso")
  }

  disk {
    file = var.virtio_iso_path
  }


  network_interface {
    network_id     = var.network_id
    hostname       = var.vm_hostname
    wait_for_lease = true
  }

  graphics {
    type        = "vnc"
    listen_type = "address"
    autoport    = true
  }

  console {
    type        = "pty"
    target_type = "serial"
    target_port = "0"
  }

  boot_device {
    dev = ["cdrom", "hd"]
  }
}

locals {
  depends_on = [libvirt_domain.windows_vm]
  ip_addresses = libvirt_domain.windows_vm.network_interface[0].addresses
}


# resource "local_file" "ansible_inventory" {
#   filename = "${path.root}/files/${var.module_name}/ansible/inventory.ini"
#   content  = <<EOT
# [windows_servers]
# ${join("\n", [for ip in local.ip_addresses : "${ip} ansible_user=${var.automation_username} ansible_ssh_private_key_file=${var.ssh_private_key_path} ansible_ssh_common_args='-o StrictHostKeyChecking=no'"])}
# EOT
# }


# resource "null_resource" "ansible_playbook" {
#   depends_on = [local_file.ansible_inventory]
#   provisioner "local-exec" {
#     command = <<EOT
#       ANSIBLE_PRIVATE_KEY_FILE=../${var.ssh_private_key_path} \
#       ansible-playbook -i ${local_file.ansible_inventory.filename} ../k8s.yml --tags 'etcd'
#     EOT
#   }
# }


resource "null_resource" "cleanup_on_destroy" {
  provisioner "local-exec" {
    when    = destroy
    command = <<EOT
      rm ./${path.module}/images/autounattend.iso
    EOT
  }

  triggers = {
    domain = libvirt_domain.windows_vm.name
  }
}


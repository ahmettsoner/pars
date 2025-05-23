resource "libvirt_volume" "etcd_volume" {
  count          = var.vm_count
  name           = "${var.vm_name}-volume-${count.index}"
  pool           = var.vm_pool
  base_volume_id = var.base_volume_id
  format         = "qcow2"
}
locals {
  ip_addresses = [
    for i in range(var.vm_count) : replace(var.vm_ip_suffix, "xxx", tostring(i + 171))
  ]
}

data "template_file" "cloudinit" {
  count  = var.vm_count
  template = file("${path.module}/cloudinit.tpl")

  vars = {
    ssh_authorized_key = var.ssh_authorized_key
    hostname           = var.vm_hostname
  }
}
data "template_file" "cloudinit_meta" {
  count  = var.vm_count
  template = file("${path.module}/metadata.tpl")

  vars = {
    hostname           = var.vm_hostname
  }
}
data "template_file" "network_data" {
  count  = var.vm_count
  template = file("${path.module}/network_data.tpl")

  vars = {
    hostname           = var.vm_hostname
    ip_address         = local.ip_addresses[count.index]
    ip_block           = var.vm_ip_block
    gateway            = replace(var.vm_ip_suffix, "xxx", "1")
  }
}

resource "libvirt_cloudinit_disk" "etcd_commoninit" {
  count     = var.vm_count
  name      = "k8s_commoninit_etcd-${count.index}.iso"
  pool      = var.vm_pool
  user_data = data.template_file.cloudinit[count.index].rendered
  meta_data = data.template_file.cloudinit_meta[count.index].rendered
  network_config = data.template_file.network_data[count.index].rendered
}

resource "libvirt_domain" "etcd_vms" {
  count  = var.vm_count
  name   = "${var.vm_name}-${count.index}"
  memory = 1024
  vcpu   = 1

  disk {
    volume_id = libvirt_volume.etcd_volume[count.index].id
  }

  cloudinit = libvirt_cloudinit_disk.etcd_commoninit[count.index].id

  network_interface {
    network_name = "default"
    hostname     = var.vm_hostname
  }

  console {
    type        = "pty"
    target_port = "0"
    target_type = "serial"
  }

  graphics {
    type        = "spice"
    listen_type = "none"
  }

  boot_device {
    dev = ["hd"]
  }
}

resource "local_file" "ansible_inventory" {
  filename = "${path.root}/files/${var.module_name}/ansible/inventory.ini"
  content  = <<EOT
[etcd_servers]
${join("\n", [for ip in local.ip_addresses : "${ip} ansible_user=${var.automation_username} ansible_ssh_private_key_file=${var.ssh_private_key_path} ansible_ssh_common_args='-o StrictHostKeyChecking=no'"])}
EOT
}


resource "null_resource" "ansible_playbook" {
  depends_on = [local_file.ansible_inventory]
  provisioner "local-exec" {
    command = <<EOT
      for ip in ${join(" ", local.ip_addresses)}; do
        ssh-keygen -R $ip;
      done
      ANSIBLE_PRIVATE_KEY_FILE=../${var.ssh_private_key_path} \
      ansible-playbook -i ${local_file.ansible_inventory.filename} ../k8s.yml --tags 'etcd'
    EOT
  }
}

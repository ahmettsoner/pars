resource "libvirt_volume" "worker_volume" {
  count          = var.vm_count
  name           = "${var.vm_name}-volume-${count.index}"
  pool           = var.vm_pool
  base_volume_id = var.base_volume_id
  format         = "qcow2"
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

resource "libvirt_cloudinit_disk" "worker_commoninit" {
  count     = var.vm_count
  name      = "k8s_commoninit_worker-${count.index}.iso"
  pool      = var.vm_pool
  user_data = data.template_file.cloudinit[count.index].rendered
  meta_data = data.template_file.cloudinit_meta[count.index].rendered
}

resource "libvirt_domain" "worker_vms" {
  count  = var.vm_count
  name   = "${var.vm_name}-${count.index}"
  memory = 1024
  vcpu   = 1

  disk {
    volume_id = libvirt_volume.worker_volume[count.index].id
  }

  cloudinit = libvirt_cloudinit_disk.worker_commoninit[count.index].id

  network_interface {
    network_name = "default"
    hostname     = var.vm_hostname
    wait_for_lease = true
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
  depends_on = [libvirt_domain.worker_vms]
  filename = "${path.root}/files/${var.module_name}/ansible/inventory.ini"

  content = join("\n", ["[worker_servers]"], flatten([
    for vm in libvirt_domain.worker_vms : (
      vm.network_interface[0].addresses != null ?
      [for ip in vm.network_interface[0].addresses :
        "${ip} ansible_user=${var.automation_username} ansible_ssh_private_key_file=${var.ssh_private_key_path} ansible_ssh_common_args='-o StrictHostKeyChecking=no'"
      ] : []
    )
  ]))
}


resource "null_resource" "ansible_playbook" {
  depends_on = [local_file.ansible_inventory]
  provisioner "local-exec" {
    command = <<EOT
    for ip in ${join(" ", flatten([
      for vm in libvirt_domain.worker_vms : (
        vm.network_interface[0].addresses != null ? vm.network_interface[0].addresses : []
      )
    ]))}; do
      ssh-keygen -R $ip;
    done
    ANSIBLE_PRIVATE_KEY_FILE=../${var.ssh_private_key_path} \
    ansible-playbook -i ${local_file.ansible_inventory.filename} ../ansible/k8s.yml --tags 'worker,k8s'
EOT
  }
}

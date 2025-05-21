resource "libvirt_volume" "vm_volumes" {
  count          = var.vm_count
  name           = "fedora41-vm-volume-${count.index}"
  pool           = var.vm_pool
  base_volume_id = var.base_volume_id
  format         = "qcow2"
}

data "template_file" "cloudinit" {
  template = file("${path.module}/cloudinit.tpl")

  vars = {
    ssh_authorized_key = pathexpand(var.ssh_authorized_key)
    hostname           = "fedora41-vm"
    username           = var.username
  }
}

resource "libvirt_cloudinit_disk" "commoninit" {
  name      = "commoninit.iso"
  pool      = var.vm_pool
  user_data = data.template_file.cloudinit.rendered
}

resource "libvirt_domain" "fedora_vms" {
  count  = var.vm_count
  depends_on = [libvirt_cloudinit_disk.commoninit]
  name   = "fedora41-vm-${count.index}"
  memory = var.memory
  vcpu   = var.vcpu

  disk {
    volume_id = libvirt_volume.vm_volumes[count.index].id
  }

  cloudinit = libvirt_cloudinit_disk.commoninit.id

  network_interface {
    network_name = var.network_name
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

  # provisioner "local-exec" {
  #   when    = destroy
  #   command = "virsh undefine fedora41-vm-${count.index} --remove-all-storage || true"
  # }
}

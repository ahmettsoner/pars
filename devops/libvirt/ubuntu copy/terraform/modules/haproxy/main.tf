
resource "libvirt_volume" "haproxy_volume" {
  name           = "haproxy-vm-volume"
  pool           = var.vm_pool
  base_volume_id = var.base_volume_id
  format         = "qcow2"
}

data "template_file" "cloudinit" {
  template = file("${path.module}/cloudinit.tpl")

  vars = {
    ssh_authorized_key = pathexpand(var.ssh_authorized_key)
    hostname           = "haproxy-vm"
    username           = var.username
  }
}

resource "libvirt_cloudinit_disk" "commoninit" {
  name      = "commoninit1.iso"
  pool      = var.vm_pool
  user_data = data.template_file.cloudinit.rendered
}


# data "template_file" "haproxy_cloudinit" {
#   template = file("${path.module}/haproxy-cloudinit.yaml")
# }

# resource "libvirt_cloudinit_disk" "haproxy_init" {
#   name      = "haproxy-init.iso"
#   pool      = var.vm_pool
#   user_data = data.template_file.haproxy_cloudinit.rendered
# }

resource "libvirt_domain" "haproxy_vm" {
  name   = "haproxy-vm"
  memory = 1024
  vcpu   = 1

  disk {
    volume_id = libvirt_volume.haproxy_volume.id
  }

  cloudinit = libvirt_cloudinit_disk.commoninit.id

  network_interface {
    network_name = "default"
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
  #   command = "virsh undefine ${self.name} || true"
  # }
}

resource "libvirt_volume" "windows_disk" {
  name   = "${var.vm_name}.qcow2"
  pool   = var.pool_id
  size   = 50 * 1024 * 1024 * 1024  # 50 GB
  format = "qcow2"
}
resource "libvirt_volume" "autounattend_disk" {
  name   = "${var.vm_name}_autounattend.qcow2"
  source = var.autounattend_iso_path
  pool   = var.pool_id
  size   = 50 * 1024 * 1024 * 1024  # 50 GB
  format = "qcow2"
}

resource "libvirt_domain" "windows_vm" {
  name   = var.vm_name
  memory = var.memory
  vcpu   = var.vcpus



  disk {
    volume_id = libvirt_volume.windows_disk.id
  }

  disk {
    file = var.autounattend_iso_path
  }

  disk {
    file = var.windows_iso_path
  }

  disk {
    file = var.virtio_iso_path
  }

  network_interface {
    network_id     = var.network_id
    wait_for_lease = false
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

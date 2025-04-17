provider "libvirt" {
  uri = "qemu:///system"
}

resource "libvirt_volume" "windows_iso" {
  name = "windows_server.iso"
  pool = "default"
  source = "/path/to/your/windows-server.iso"
}

resource "libvirt_volume" "win10_qcow2" {
  name = "windows-server.qcow2"
  pool = "default"
  format = "qcow2"
  size   = 50 # Specify the size as needed
}

resource "libvirt_domain" "windows_server" {
  name   = "windows-server"
  memory = "4096" # 4GB RAM
  vcpu   = 2

  disk {
    volume_id = libvirt_volume.win10_qcow2.id
  }

  cdrom {
    volume_id = libvirt_volume.windows_iso.id
  }

  network_interface {
    network_name = "default"
  }

  graphics {
    type        = "vnc"
    listen_type = "address"
    listen_address = "0.0.0.0"
  }

  console {
    type        = "pty"
    target_type = "serial"
    target_port = "0"
  }

  # UEFI boot, optional
  # uefi {
  #   enabled = true
  # }
}
resource "libvirt_volume" "win_disk" {
  count          = var.vm_count
  name           = "${var.vm_name}-volume-${count.index}"
  pool           = var.vm_pool
  format         = "qcow2"
  size           = 50 * 1024 * 1024 * 1024 # 50 GB
}

data "template_file" "autounattend" {
  template = file("${path.module}/autounattend/Autounattend.tpl")

  vars = {
    ssh_authorized_key = var.ssh_authorized_key
    username           = var.automation_username
  }
}
resource "local_file" "autounattend_xml" {
  count  = var.vm_count
  content  = data.template_file.autounattend.rendered
  filename = "${path.module}/autounattend/${count.index}/Autounattend.xml"
}

resource "null_resource" "build_autounattend_iso" {
  depends_on = [local_file.autounattend_xml]
  count  = var.vm_count

  provisioner "local-exec" {
    command = <<EOT
      ./${path.module}/scripts/create-autounattend-iso.sh ./${path.module}/autounattend/${count.index} ./${path.module}/images/autounattend-${count.index}.iso
    EOT
  }
}

resource "libvirt_domain" "windows_vms" {
  depends_on = [null_resource.build_autounattend_iso]
  count  = var.vm_count
  name   = "${var.vm_name}-${count.index}"
  memory = var.memory
  vcpu   = var.vcpu


  disk {
    volume_id = libvirt_volume.win_disk[count.index].id
  }

  disk {
    file = var.windows_iso_path
  }

  disk {
    file = abspath("${path.module}/images/autounattend-${count.index}.iso")
  }

  disk {
    file = var.virtio_iso_path
  }


  network_interface {
    network_id     = var.network_id
    hostname       = var.vm_hostname
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

resource "null_resource" "wait_for_ip" {
  depends_on = [libvirt_domain.windows_vms]
  # count  = var.vm_count

  provisioner "local-exec" {
    command = <<EOT
#!/bin/bash
for i in {1..60}; do
  IP=$(sudo virsh domifaddr ${var.vm_name}-0| grep ipv4 | awk '{print $4}' | cut -d/ -f1)
  if [[ ! -z "$IP" ]]; then
    echo "VM IP: $IP"
    echo "$IP" > ${path.root}/files/${var.module_name}/vm_ip.txt
    exit 0
  fi
  echo "Waiting for IP..."
  sleep 10
done
echo "Timeout waiting for VM IP"
exit 1
EOT
    interpreter = ["bash", "-c"]
  }
}

data "external" "vm_ip" {
  depends_on = [null_resource.wait_for_ip]

  program = ["bash", "-c", "${abspath(path.module)}/scripts/get_ip.sh ${var.vm_name}-0"]
}


locals {
  depends_on = [null_resource.wait_for_ip]
  
  ip_addresses = [data.external.vm_ip.result.ip]
}



resource "local_file" "ansible_inventory" {
  depends_on = [null_resource.wait_for_ip]
  
  filename = "${path.root}/files/${var.module_name}/ansible/inventory.ini"
  content  = <<EOT
[windows_servers]
${join("\n", [for ip in local.ip_addresses : "${ip} ansible_user=${var.automation_username} ansible_ssh_private_key_file=${var.ssh_private_key_path} ansible_ssh_common_args='-o StrictHostKeyChecking=no'"])}
EOT
}


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
  count  = var.vm_count

  provisioner "local-exec" {
    when    = destroy
    command = <<EOT
      rm ./${path.module}/images/autounattend-${count.index}.iso
    EOT
  }

  triggers = {
    domain = libvirt_domain.windows_vms[count.index].name
  }
}
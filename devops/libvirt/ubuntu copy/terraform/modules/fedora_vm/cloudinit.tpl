#cloud-config
hostname: ${hostname}
users:
  - name: ${username}
    sudo: ALL=(ALL) NOPASSWD:ALL
    groups: wheel
    shell: /bin/bash
    ssh_authorized_keys:
      - ${ssh_authorized_key}

#cloud-config
hostname: ${hostname}
users:
    - name: ${username}
      sudo: ALL=(ALL) NOPASSWD:ALL
      shell: /bin/bash
      ssh-authorized-keys:
          - ${ssh_authorized_key}
      lock_passwd: true

disable_root: false # root erişimi engellenmesin (opsiyonel)

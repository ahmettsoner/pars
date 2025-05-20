#cloud-config
hostname: ${hostname}
users:
  - name: ${username}
    sudo: ALL=(ALL) NOPASSWD:ALL
    groups: wheel
    shell: /bin/bash
    ssh_authorized_keys:
      - ${ssh_authorized_key}

write_files:
  - path: /etc/netplan/01-netcfg.yaml
    content: |
      network:
        version: 2
        ethernets:
          eth0:
            addresses:
              - 192.168.124.100/24
            gateway4: 192.168.124.1
            nameservers:
              addresses: [8.8.8.8, 1.1.1.1]
runcmd:
  - netplan apply

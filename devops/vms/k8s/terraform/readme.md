```bash
sudo usermod -aG libvirt libvirt-qemu
```

```bash
sudo systemctl restart libvirtd
```

```bash
rm -rf ./vms/k8s/terraform/.terraform* ./vms/k8s/terraform/terraform.tfstate*
```

```bash
mkdir -p ./images
```

```bash
wget -P ./images https://dl.fedoraproject.org/pub/fedora/linux/releases/41/Cloud/x86_64/images/Fedora-Cloud-Base-Generic-41-1.4.x86_64.qcow2
```

```bash
wget -P ./images https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img
qemu-img convert -O qcow2 ./images/jammy-server-cloudimg-amd64.img ./images/jammy-server-cloudimg-amd64.qcow2
```

```bash
wget -P ./images https://cloud-images.ubuntu.com/minimal/releases/jammy/release/ubuntu-22.04-minimal-cloudimg-amd64.img
qemu-img convert -O qcow2 ./images/ubuntu-22.04-minimal-cloudimg-amd64.img ./images/ubuntu-22.04-minimal-cloudimg-amd64.qcow2
```

```bash
sudo cp /usr/share/OVMF/OVMF_VARS.fd /var/lib/libvirt/qemu/nvram/windows-server-2022_VARS.fd
sudo chown qemu:kvm /var/lib/libvirt/qemu/nvram/windows-server-2022_VARS.fd
sudo cat /var/lib/libvirt/qemu/nvram/windows-server-2022_VARS.fd
```

```bash
wget -P ./images https://go.microsoft.com/fwlink/p/?LinkID=2195280&clcid=0x409&culture=en-us&country=US
wget -P ./images https://fedorapeople.org/groups/virt/virtio-win/direct-downloads/archive-virtio/virtio-win-0.1.100/virtio-win.iso
```

```bash
sudo apt install wimtools genisoimage cabextract p7zip-full
sudo dnf install wimlib wimlib-utils genisoimage p7zip p7zip-plugins cabextract

sudo chmod +x devops/vms/k8s/terraform/modules/windows/scripts/create-custom-windows-iso.sh
sudo sudo ./devops/vms/k8s/terraform/modules/windows/scripts/create-custom-windows-iso.sh ./devops/images/SERVER_EVAL_x64FRE_en-us.iso ./devops/images/virtio-win-0.1.100.iso ./devops/images/custom_win_server.iso
```

```bash
chmod +x devops/vms/k8s/terraform/modules/windows/scripts/create-autounattend-iso.sh
sudo ./devops/vms/k8s/terraform/modules/windows/scripts/create-autounattend-iso.sh ./devops/vms/k8s/terraform/modules/windows/autounattend ./devops/images/autounattend.iso
```

```bash
chmod +x devops/vms/k8s/terraform/modules/windows/scripts/create-autounattend-iso-xorriso.sh
sudo ./devops/vms/k8s/terraform/modules/windows/scripts/create-autounattend-iso-xorriso.sh ./devops/vms/k8s/terraform/modules/windows/autounattend ./devops/images/autounattend.iso ./devops/images/SERVER_EVAL_x64FRE_en-us.iso
```

```bash
sudo chown vms-qemu:kvm /var/lib/vms/images/
sudo chmod 644 /var/lib/vms/images/
```

## 1. Adım

```bash
terraform -chdir=./devops/vms/k8s/terraform init
```

## 2. Adım

```bash
terraform -chdir=./devops/vms/k8s/terraform apply -auto-approve
```

```bash
terraform -chdir=./devops/vms/k8s/terraform destroy
```

```
sudo apt install ovmf
```

list vms:

```bash
sudo virsh list --all
```

get ip for vm

```bash
sudo virsh domifaddr k8s-haproxy-vm
sudo virsh console k8s-haproxy-vm

```

```bash
sudo virsh shutdown k8s-master-vm-0
sudo virsh shutdown k8s-master-vm-1
sudo virsh shutdown k8s-master-vm-2
sudo virsh shutdown k8s-worker-vm-0
sudo virsh shutdown k8s-worker-vm-1
sudo virsh shutdown k8s-worker-vm-2
sudo virsh shutdown k8s-haproxy-vm

```

```bash
sudo virsh undefine k8s-etcd-vm-0
sudo virsh undefine k8s-etcd-vm-1
sudo virsh undefine k8s-etcd-vm-2
sudo virsh pool-destroy k8s_pool
sudo virsh pool-undefine k8s_pool

```

```bash
sudo virsh undefine k8s-master-vm-0
sudo virsh undefine k8s-master-vm-1
sudo virsh undefine k8s-master-vm-2
sudo virsh undefine k8s-worker-vm-0
sudo virsh undefine k8s-worker-vm-1
sudo virsh undefine k8s-worker-vm-2
sudo virsh undefine k8s-haproxy-vm

```

```bash
sudo virsh vol-delete k8s_commoninit_haproxy.iso
sudo virsh vol-delete k8s_commoninit_master.iso
sudo virsh vol-delete k8s_commoninit_worker.iso

```

```bash
ssh -i ./vms/k8s/terraform/files/ssh_keys/infra_id_rsa automation@192.168.124.171
```

```bash
ANSIBLE_PRIVATE_KEY_FILE=../files/ssh_keys/infra_id_rsa
cd vms/k8s/terraform
ansible-playbook -i ./files/etcd/ansible/inventory.ini ../k8s.yml --tags 'etcd'

```

```bash
sudo virsh pool-list --all

```

```bash
sudo qemu-img create -f qcow2 /var/lib/libvirt/images/win2022.qcow2 50G


sudo virt-install \
  --name win2022 \
  --memory 4096 \
  --vcpus 2 \
  --disk path=/var/lib/libvirt/images/win2022.qcow2,format=qcow2 \
  --disk /home/ahmetsoner/AS/prs/pars/devops/images/SERVER_EVAL_x64FRE_en-us.iso,device=cdrom \
  --disk /home/ahmetsoner/AS/prs/pars/devops/images/autounattend.iso,device=cdrom \
  --disk /home/ahmetsoner/AS/prs/pars/devops/images/virtio-win-0.1.100.iso,device=cdrom \
  --os-variant win2k22 \
  --network network=default \
  --graphics vnc \
  --boot cdrom,hd \
  --noautoconsole



```

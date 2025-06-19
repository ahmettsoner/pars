```bash
sudo usermod -aG libvirt libvirt-qemu
```

```bash
sudo systemctl restart libvirtd
```

```
cp -r ./devops/images/ /var/lib/libvirt
```

```bash
rm -rf ./devops/vms/k8s/terraform/.terraform* ./devops/vms/k8s/terraform/terraform.tfstate*
```


```bash
wget -P /var/lib/libvirt/images https://dl.fedoraproject.org/pub/fedora/linux/releases/41/Cloud/x86_64/images/Fedora-Cloud-Base-Generic-41-1.4.x86_64.qcow2
```

```bash
wget -P /var/lib/libvirt/images https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img
qemu-img convert -O qcow2 /var/lib/libvirt/images/jammy-server-cloudimg-amd64.img /var/lib/libvirt/images/jammy-server-cloudimg-amd64.qcow2
```

```bash
wget -P /var/lib/libvirt/images https://cloud-images.ubuntu.com/minimal/releases/jammy/release/ubuntu-22.04-minimal-cloudimg-amd64.img
qemu-img convert -O qcow2 /var/lib/libvirt/images/ubuntu-22.04-minimal-cloudimg-amd64.img /var/lib/libvirt/images/ubuntu-22.04-minimal-cloudimg-amd64.qcow2
```

```bash
wget -P /var/lib/libvirt/images https://go.microsoft.com/fwlink/p/?LinkID=2195280&clcid=0x409&culture=en-us&country=US
wget -P /var/lib/libvirt/images https://fedorapeople.org/groups/virt/virtio-win/direct-downloads/archive-virtio/virtio-win-0.1.100/virtio-win.iso
```

```bash
sudo apt install wimtools genisoimage cabextract p7zip-full
sudo dnf install wimlib wimlib-utils genisoimage p7zip p7zip-plugins cabextract
```

```bash
chmod +x ./devops/vms/k8s/terraform/modules/windows/scripts/create-autounattend-iso.sh
sudo ./devops/vms/k8s/terraform/modules/windows/scripts/create-autounattend-iso.sh ./devops/vms/k8s/terraform/modules/windows/autounattend ./devops/images/autounattend.iso
```
```bash
chmod +x ./devops/vms/k8s/terraform/modules/windows/scripts/get_ip.sh
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
terraform -chdir=./devops/vms/k8s/terraform destroy -auto-approve
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
ssh -i ./devops/vms/k8s/terraform/files/ssh_keys/infra_id_rsa automation@192.168.124.171
ssh -i ~/.ssh/id_rsa user1@192.168.122.24
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
sudo ./devops/vms/k8s/terraform/modules/windows/scripts/create-autounattend-iso.sh ./devops/vms/k8s/terraform/modules/windows/autounattend ./devops/images/autounattend.iso
```
```bash
sudo rm /var/lib/libvirt/images/win2022.qcow2
```
```bash
sudo qemu-img create -f qcow2 /var/lib/libvirt/images/win2022.qcow2 50G
```
```bash
sudo virt-install \
  --name win2022 \
  --memory 4096 \
  --vcpus 2 \
  --disk $(pwd)/devops/images/SERVER_EVAL_x64FRE_en-us.iso,device=cdrom \
  --disk $(pwd)/devops/images/autounattend.iso,device=cdrom \
  --disk $(pwd)/devops/images/virtio-win.iso,device=cdrom \
  --os-variant win2k22 \
  --network network=default \
  --graphics vnc \
  --boot cdrom,hd \
  --noautoconsole

```
```bash
sudo virt-install \
  --name win2022 \
  --memory 4096 \
  --vcpus 2 \
  --disk path=/var/lib/libvirt/images/win2022.qcow2,format=qcow2 \
  --cdrom $(pwd)/devops/images/SERVER_EVAL_x64FRE_en-us.iso \
  --disk $(pwd)/devops/images/autounattend.iso,device=cdrom \
  --disk $(pwd)/devops/images/virtio-win.iso,device=cdrom \
  --os-variant win2k22 \
  --network network=default \
  --graphics vnc \
  --boot cdrom,hd \
  --autostart \
  --noautoconsole


```



Edit `C:\ProgramData\ssh\sshd_config`





```pwsh
$sshDir = Join-Path $env:USERPROFILE ".ssh"
$authKeysPath = Join-Path $sshDir "authorized_keys"

$publicKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCdyr8788+c/Kl5TGCLO1pXUODWb0OYa4w5cAKoz40227gdlfo0tMXlny8mrUH9KvuKXPtLYKU42vMx+SDJKD8/cN1SN/umAaxwgAGMTIQMfOjo5aAiOwK5jLK1dLr0bDj2AUiqhrwj8WmGJB4Z+GYa0AZHkfrgrgoSW+KN0J+KvkBjamHOu6FUL7DUZC3vnsnZVo4IomRENsxnjIlsVatIzJcgl3b5kAXtKOzLJ5/UTUCNmjRJmbV6u4j8okQAfkFePaotyDTUcKJYgOdP5kHGexXGIjBAodfhbBmxwC2dx4T/KTH9/nqdVXq9mhbArJO70yAkZVeSLYFfNUY0i6GW7FQNrL9yuceG+PnzIpr8hktRopTZWw/vJU60QYbBUhLW4uhWE3aE56eznWA6LNED2DijK5d5U5ZeAIU3i9wcKVP4RzvszTEfoE0jvLl5FkEda7ufmhcqGyOgFHlUlPutct94UFHZ+fTEnSfz+JtP2ADxRZKGBcWvLD8Y813Y0rswT1Ib7Eumi784JvtxbJoIqbuc4EdcUm/SIlHjKTjrZ47sU65dEklNkJMi91/9z3ZnSbydGkl5JobzfH8U6VanwOApEb9Dyun2SgFiNs4Qf/klOhjz57V8qkoBesOk+E4bR1BCHiuQ4Mi+J3e4U9yB6HxAo/uIGlLK5swW4OysNw== ahmet.soner@pusula.int"

New-Item -ItemType Directory -Path $sshDir -Force

Set-Content -Path $authKeysPath -Value $publicKey -Encoding ascii

# # İzinleri ayarla: inheritance kapat, sadece kullanıcı izinli olsun
# icacls $sshDir /inheritance:r
# # icacls $sshDir /remove "Administrators" "Users" "Authenticated Users" "BUILTIN\Users" "BUILTIN\Administrators"
# icacls $sshDir /grant:r "${env:USERNAME}`:(F)"
# # icacls $sshDir /grant:r "SYSTEM:(F)"
# icacls $authKeysPath /inheritance:d
# # icacls $authKeysPath /remove "Administrators" "Users" "Authenticated Users" "BUILTIN\Users" "BUILTIN\Administrators"
# icacls $authKeysPath /grant:r "${env:USERNAME}`:(F)"
# # icacls $authKeysPath /grant:r "SYSTEM:(F)"
```




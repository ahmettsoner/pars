```
rm -rf ./vms/k8s/terraform/.terraform* ./vms/k8s/terraform/terraform.tfstate*
```

```
mkdir -p ./images
```

```
wget -P ./images https://dl.fedoraproject.org/pub/fedora/linux/releases/41/Cloud/x86_64/images/Fedora-Cloud-Base-Generic-41-1.4.x86_64.qcow2
```

```
wget -P ./images https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img
qemu-img convert -O qcow2 ./images/jammy-server-cloudimg-amd64.img ./images/jammy-server-cloudimg-amd64.qcow2
```

```
wget -P ./images https://cloud-images.ubuntu.com/minimal/releases/jammy/release/ubuntu-22.04-minimal-cloudimg-amd64.img
qemu-img convert -O qcow2 ./images/ubuntu-22.04-minimal-cloudimg-amd64.img ./images/ubuntu-22.04-minimal-cloudimg-amd64.qcow2
```

```
sudo chown vms-qemu:kvm /var/lib/vms/images/
sudo chmod 644 /var/lib/vms/images/
```

## 1. Adım

```
terraform -chdir=./vms/k8s/terraform init
```

## 2. Adım

```
terraform -chdir=./vms/k8s/terraform apply -auto-approve
```

```
terraform -chdir=./vms/k8s/terraform destroy
```

list vms:

```
sudo virsh list --all
```

get ip for vm

```
sudo virsh domifaddr k8s-haproxy-vm
sudo virsh console k8s-haproxy-vm
```

```
sudo virsh shutdown k8s-master-vm-0
sudo virsh shutdown k8s-master-vm-1
sudo virsh shutdown k8s-master-vm-2
sudo virsh shutdown k8s-worker-vm-0
sudo virsh shutdown k8s-worker-vm-1
sudo virsh shutdown k8s-worker-vm-2
sudo virsh shutdown k8s-haproxy-vm
```

```
sudo virsh undefine k8s-master-vm-0
sudo virsh undefine k8s-master-vm-1
sudo virsh undefine k8s-master-vm-2
sudo virsh undefine k8s-worker-vm-0
sudo virsh undefine k8s-worker-vm-1
sudo virsh undefine k8s-worker-vm-2
sudo virsh undefine k8s-haproxy-vm
```

```
sudo virsh vol-delete k8s_commoninit_haproxy.iso
sudo virsh vol-delete k8s_commoninit_master.iso
sudo virsh vol-delete k8s_commoninit_worker.iso
```

```
ssh -i ./vms/k8s/terraform/files/ssh_keys/infra_id_rsa automation@192.168.124.171
```

```
ANSIBLE_PRIVATE_KEY_FILE=../files/ssh_keys/infra_id_rsa
cd vms/k8s/terraform
ansible-playbook -i ./files/etcd/ansible/inventory.ini ../k8s.yml --tags 'etcd'

```

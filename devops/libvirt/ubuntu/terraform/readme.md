```
rm -rf ./libvirt/ubuntu/terraform/.terraform* ./libvirt/ubuntu/terraform/terraform.tfstate*
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
sudo chown libvirt-qemu:kvm /var/lib/libvirt/images/
sudo chmod 644 /var/lib/libvirt/images/
```

```
terraform -chdir=./libvirt/ubuntu/terraform init
```

```
terraform -chdir=./libvirt/ubuntu/terraform apply -auto-approve
```

```
terraform -chdir=./libvirt/ubuntu/terraform destroy
```

list vms:

```
sudo virsh list --all
```

get ip for vm

```
sudo virsh domifaddr fedora41-vm-0
```

```
sudo virsh shutdown fedora41-vm-0
sudo virsh shutdown fedora41-vm-1
sudo virsh shutdown fedora41-vm-2
sudo virsh shutdown haproxy-vm
```

```
sudo virsh destroy fedora41-vm-0
sudo virsh destroy fedora41-vm-1
sudo virsh destroy fedora41-vm-2
sudo virsh destroy haproxy-vm
```

```
sudo virsh undefine fedora41-vm-0
sudo virsh undefine fedora41-vm-1
sudo virsh undefine fedora41-vm-2
sudo virsh undefine haproxy-vm
```

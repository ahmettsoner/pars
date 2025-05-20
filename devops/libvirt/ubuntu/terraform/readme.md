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
terraform -chdir=./libvirt/ubuntu/terraform terraform init -reconfigure
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
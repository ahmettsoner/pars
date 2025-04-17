## Ön Gereksinimler

1. KVM ve gerekli araçlar Debian tabanlı bir sistemde kurulu olmalıdır:

```bash
sudo apt update
sudo apt install qemu-kvm libvirt-daemon-system libvirt-clients bridge-utils virt-manager
```

2. Terraform ve Libvirt Provider for Terraform kurulu olmalıdır. Bunun için Terraform binary indirilerek PATH'e eklenmeli ve libvirt plugin'i yüklenmelidir.

3. Terraform İle Uygulama:

```bash
terraform init    # Terraform çalışma dizinini başlatır ve gerekli eklentileri yükler
terraform apply   # Sanal makineyi oluşturur
```

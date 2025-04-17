## Ön Gereksinimler

1. Ansible'ı Kur

```bash
sudo apt update
sudo apt install ansible -y
```

2. Ansible ile sanal makine oluşturma işlemi genellikle bir imaj veya bulut ortamı üzerinden olur. Direkt KVM üzerinde ISO'dan Ansible ile VM oluşturmak karmaşık olabilir. virt-install gibi komutlarla bu işlemi basitleştirebiliriz.

3. Playbook'u Çalıştırma

```bash
ansible-playbook node_go_setup.yml
```

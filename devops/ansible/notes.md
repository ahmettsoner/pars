install
fedora

```
sudo dnf install ansible -y
```

run

```
ansible-playbook -i inventory.ini playbook.yml --ask-become-pass
```

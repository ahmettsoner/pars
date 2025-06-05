## yum client settings

```
cat <<EOF | sudo tee /etc/yum.repos.d/pars-dev.repo > /dev/null
[pars]
name=Pars Repository
baseurl=http://localhost:8070/releases/pars/dev/yum/
enabled=1
gpgcheck=1
gpgkey=http://localhost:8070/downloads/public-keys/gpg/public-rpm.gpg
EOF
```

```
sudo yum install pars
```

```
sudo dnf install pars
```

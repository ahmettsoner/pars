## yum client settings

```
cat <<EOF | sudo tee /etc/yum.repos.d/pars-dev.repo > /dev/null
[pars]
name=Pars Repository
baseurl=http://localhost:8070/releases/pars/dev/yum/
enabled=1
gpgcheck=1
gpgkey=http://localhost:8070/downloads/keys/gpg/public-rpm-dev.gpg
EOF
```

```
sudo yum install pars
```

```
sudo dnf install pars
```

## apt client settings

```
curl -fsSL http://localhost:8070/downloads/keys/gpg/public-deb-dev.gpg | sudo tee /usr/share/keyrings/pars-dev-archive-keyring.gpg > /dev/null
```

```
echo "deb [signed-by=/usr/share/keyrings/pars-dev-archive-keyring.gpg] http://localhost:8070/releases/pars/dev/apt/ universal main" | sudo tee /etc/apt/sources.list.d/pars-dev.list
```

```
sudo apt update
```

```
sudo apt install pars
```

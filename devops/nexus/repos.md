## yum client settings

```
cat <<EOF | sudo tee /etc/yum.repos.d/pars.repo > /dev/null
[pars]
name=Pars YUM Repo
baseurl=http://localhost:8081/repository/yum/
enabled=1
gpgcheck=1
gpgkey=http://localhost:8081/repository/public-keys/rpm/gpg/public-rpm.gpg
EOF
```

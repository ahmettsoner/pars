## yum client settings

```
cat <<EOF | sudo tee /etc/yum.repos.d/pars.repo > /dev/null
[pars]
name=Pars Repository
baseurl=http://localhost:8081/repository/yum/
enabled=1
gpgcheck=1
gpgkey=http://localhost:8081/repository/keys/rpm/gpg/public-rpm.gpg
EOF
```

YUM REPO
create yum hosted repo
default ayarlarlar

-   Repodata Depth: 0

```bash
NEXUS_URL="http://localhost:8081"
USERNAME="admin"
PASSWORD="your-admin-password"
REPO_NAME="pars-apt"

curl -u "$USERNAME:$PASSWORD" -X POST "$NEXUS_URL/service/rest/v1/repositories/apt/hosted" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "'"$REPO_NAME"'",
    "online": true,
    "storage": {
      "blobStoreName": "default",
      "strictContentTypeValidation": true,
      "writePolicy": "ALLOW"
    },
    "apt": {
      "distribution": "stable",
      "flat": false,
      "signing": {
        "keypair": "pars-sign-key",  // daha önce Nexus'a tanımlanmışsa
        "passphrase": "gizli-sifre"  // Nexus içindeki GPG key'in şifresi
      }
    }
  }'

```

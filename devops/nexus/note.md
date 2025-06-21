```
mkdir -p ./devops/nexus/data/nexus
mkdir -p ./devops/nexus/data/nginx
sudo chown -R $(whoami):$(whoami) ./devops/nexus/data
sudo chown -R 200:200 ./devops/nexus/data/nexus
```

```
docker compose -f ./devops/nexus/docker-compose.yaml up -d
```

get admin pass

```
docker exec -it nexus cat /nexus-data/admin.password
```

## Nexus UI: http://localhost:8081

Public RPM Package Repository

```
gpg --full-generate-key

gpg (GnuPG) 2.4.5; Copyright (C) 2024 g10 Code GmbH
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Please select what kind of key you want:
   (1) RSA and RSA
   (2) DSA and Elgamal
   (3) DSA (sign only)
   (4) RSA (sign only)
   (9) ECC (sign and encrypt) *default*
  (10) ECC (sign only)
  (14) Existing key from card
Your selection? 0
Invalid selection.
Your selection? 1
RSA keys may be between 1024 and 4096 bits long.
What keysize do you want? (3072) 4096
Requested keysize is 4096 bits
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 0
Key does not expire at all
Is this correct? (y/N) y

GnuPG needs to construct a user ID to identify your key.

Real name: ParsDevKit
Email address: support@parsdevkit.net
Comment: Pars Repo Key
You selected this USER-ID:
    "ParsDevKit (Pars Repo Key) <support@parsdevkit.net>"

Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? O
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
gpg: directory '/home/ahmetsoner/.gnupg/openpgp-revocs.d' created
gpg: revocation certificate stored as '/home/ahmetsoner/.gnupg/openpgp-revocs.d/9BB6C221634CFCF0CFE58AF0B5978DE2981A203F.rev'
public and secret key created and signed.

pub   rsa4096 2025-05-27 [SC]
      9BB6C221634CFCF0CFE58AF0B5978DE2981A203F
uid                      ParsDevKit (Pars Repo Key) <support@parsdevkit.net>
sub   rsa4096 2025-05-27 [E]


```

Pass Phrase: `rNM0b8Bu7q0wwkb`

```
gpg --list-keys --keyid-format LONG

```

```bash
# Public key
gpg --export --armor YOUR_KEY_ID > public.gpg

# Private key
gpg --export-secret-keys --armor YOUR_KEY_ID > private.gpg

# Base64 Encode
base64 private.gpg > private.gpg.b64
```

# dearmor public key

```
gpg --dearmor devops/public.gpg > devops/public-da.gpg
```

## APT REPO
create apt hosted repo
name: apt-universial
distribution: universial
signin key: private.gpg
passphrase: 33KX4gaYW7kd


## YUM REPO
create yum hosted repo
name: yum
repodata Depth: 0

yum.repo.parsdevkit.net
rpm.repo.parsdevkit.net

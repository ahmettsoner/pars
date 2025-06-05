## Jenkins

-   add `shared-lib` to `Jenkins Dashboard → Manage Jenkins → System → Global Trusted Pipeline Libraries` as trusted library
-   add Credetials

    -   GITEA_TOKEN:
        -   Secret text
    -   nexus-credential:
        -   Username with password
    -   public-rpm.gpg
        -   Secret File
    -   private-rpm.gpg
        -   Secret File
    -   GPG_PASSPHRASE:
        -   Secret Text

-   add apt-jammy hosted repository on nexus
-   add apt-bionic hosted repository on nexus
-   add yum hosted repository on nexus
    -   Repodata Depth: 0
-

`ssh-keygen -t rsa -b 4096 -C "admin@local.git"`
`passphrase`: `33KX4gaYW7kd`
`git remote set-url local ssh://admin@192.168.118.47:2222/admin/pars.git`

`/releases/{platform}/{artifact}/{type}/{project}/{distribution}/{env}/{version}|latest/{arch}/{app}.{ext}`

# 1. Yol Yapısı Açılımı (Terminoloji)

-   **platform:** linux, windows, macos, android, ios
-   **artifact:** packages, binaries, archives, installers (dosya türü genel grubu)
-   **type:** apt, yum, raw, installer, archive, pkg, msi, dmg vb alt kategoriler
-   **project:** uygulama/proje adı, örn: pars
-   **distribution:** dağıtım türü veya versiyonu (apt için jammy, focal, yum için centos7, almalinux8 vb) veya all
-   **env:** dev, test, staging, prod gibi ortamlar
-   **version:** versiyon numarası veya `latest`
-   **arch:** amd64, x86_64, arm64, armv7 vb mimari
-   **app:** uygulama dosya adı
-   **ext:** dosya uzantısı (rpm, deb, tar.gz, run, exe vb)

yum

`/releases/{platform}/{artifact}/{type}/{project}/{distribution}/{env}`

-   proxy: /releases/linux/packages/yum/pars/dev/version|latest/arch/app.rpm
    -   /releases/linux/packages/yum/pars/dev/latest/x86_64/app.rpm
-   repo: /repository/yum/

apt

-   proxy: /releases/linux/packages/apt/pars/distribution/dev/version|latest/arch/app.rpm
    -   /releases/linux/packages/apt/pars/jammy/dev/1.0.0/amd64/app.deb
    -   /releases/linux/packages/apt/pars/jammy/dev/latest/amd64/app.deb

binary/executables

-   proxy: /releases/linux/binaries/raw/pars/dev/version|latest/linux/arch/app
-   proxy: /releases/linux/binaries/package/pars/dev/version|latest/linux/arch/app
-   proxy: /releases/linux/binaries/installer/pars/dev/version|latest/linux/arch/app
-   proxy: /releases/linux/binaries/archive/pars/dev/version|latest/linux/arch/app

`http://localhost:8070/releases/linux/binaries/raw/pars/all/dev/v1.0.0-dev.22/x86_64/pars`

---

`http://localhost:8070/`
`http://localhost:8070/releases/pars/dev/yum`
`http://localhost:8070/releases/pars/dev/apt/dists/universal/InRelease`
`http://localhost:8070/releases/pars/dev/v1.0.0-dev.22/changelog.md`
`http://localhost:8070/releases/pars/dev/v1.0.0-dev.23/linux/binaries/raw/all/x86_64/pars`
`http://localhost:8070/downloads/keys/gpg/public-deb-dev.gpg`

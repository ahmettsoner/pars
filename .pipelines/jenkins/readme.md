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

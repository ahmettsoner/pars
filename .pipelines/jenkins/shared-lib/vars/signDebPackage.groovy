def call(String OS, String ARCH, String DIST_CODENAME){
    withCredentials([
        file(credentialsId: 'private-rpm.gpg', variable: 'PRIVATE_GPG'),
        file(credentialsId: 'public-rpm.gpg', variable: 'PUBLIC_GPG'),
        string(credentialsId: 'GPG_PASSPHRASE', variable: 'GPG_PASSPHRASE')
    ]) {

        def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
        def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
        def gpgIdentity = "ParsDevKit (Pars Repo Key) <support@parsdevkit.net>"
        def repoRoot = "${env.WORKSPACE}/apt-repo"
        def poolPath = "${repoRoot}/pool/main/${APPNAME[0]}/${APPNAME}"
        def distPath = "${repoRoot}/dists/${DIST_CODENAME}/main/binary-${ARCH}"

        sh """
            set -e
            mkdir -p ~/.gnupg
            chmod 700 ~/.gnupg
            gpg --import $PRIVATE_GPG
            gpg --import $PUBLIC_GPG

            echo "Signing .deb package..."
            dpkg-sig -k '${gpgIdentity}' --sign builder '${debOutputPath}'

            echo "Building APT repository..."
            mkdir -p '${poolPath}'
            mkdir -p '${distPath}'

            cp '${debOutputPath}' '${poolPath}/'

            cd '${repoRoot}'
            dpkg-scanpackages pool /dev/null | gzip -9c > '${distPath}/Packages.gz'

            cd '${repoRoot}/dists/${DIST_CODENAME}'
            apt-ftparchive release . > Release

            echo "Signing Release file (detached)..."
            gpg --batch --yes --default-key '${gpgIdentity}' \\
                --passphrase '${GPG_PASSPHRASE}' \\
                --pinentry-mode loopback \\
                -abs -o Release.gpg Release

            echo "Creating InRelease file (clearsigned)..."
            gpg --batch --yes --default-key '${gpgIdentity}' \\
                --passphrase '${GPG_PASSPHRASE}' \\
                --pinentry-mode loopback \\
                --clearsign -o InRelease Release
        """
    }
}

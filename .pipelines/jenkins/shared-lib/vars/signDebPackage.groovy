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
        def poolPath = "${repoRoot}/pool/main/${APPNAME}"
        def distPath = "${repoRoot}/dists/${DIST_CODENAME}/main/binary-${ARCH}"

        sh '''#!/bin/bash -e
            mkdir -p ~/.gnupg
            chmod 700 ~/.gnupg

            echo "use-agent" > ~/.gnupg/gpg.conf
            echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
            echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf

            gpgconf --kill gpg-agent
            gpgconf --launch gpg-agent

            gpg --batch --import "$PRIVATE_GPG"
            gpg --batch --import "$PUBLIC_GPG"

            echo "Setting GPG trust..."
            FPR=$(gpg --list-keys --with-colons | grep '^fpr' | head -n1 | cut -d':' -f10)
            echo "$FPR:6:" > trust.txt
            gpg --import-ownertrust trust.txt

            echo "Signing .deb..."
            echo "$GPG_PASSPHRASE" | dpkg-sig -k "$FPR" --sign builder "$debOutputPath"

            echo "Creating APT repository..."
            mkdir -p "${poolPath}" "${distPath}"
            cp "$debOutputPath" "${poolPath}/"

            cd "${repoRoot}"
            dpkg-scanpackages pool /dev/null | gzip -9c > "${distPath}/Packages.gz"

            cd "${repoRoot}/dists/${DIST_CODENAME}"
            apt-ftparchive release . > Release

            echo "$GPG_PASSPHRASE" | gpg --batch --yes --pinentry-mode loopback \\
                --passphrase-fd 0 -u "$FPR" -abs -o Release.gpg Release

            echo "$GPG_PASSPHRASE" | gpg --batch --yes --pinentry-mode loopback \\
                --passphrase-fd 0 -u "$FPR" --clearsign -o InRelease Release

            echo "Cleanup"
            rm -rf ~/.gnupg trust.txt
        '''
    }
}

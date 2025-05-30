def call(String OS, String ARCH, String DIST_CODENAME){
    withCredentials([
        file(credentialsId: 'public-rpm.gpg', variable: 'GPG_PUBLIC'),
        file(credentialsId: 'private-rpm.gpg', variable: 'GPG_PRIVATE'),
        string(credentialsId: 'GPG_PASSPHRASE', variable: 'GPG_PASSPHRASE'),
        string(credentialsId: 'GPG_FINGERPRINT', variable: 'GPG_FINGERPRINT')
    ]) {

        def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
        def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

        def repoRoot = "${env.WORKSPACE}/apt-repo"
        def poolPath = "${repoRoot}/pool/main/${APPNAME}"
        def distPath = "${repoRoot}/dists/${DIST_CODENAME}/main/binary-${ARCH}"

        sh '''
            mkdir -p ~/.gnupg
            chmod 700 ~/.gnupg

            # Import GPG keys
            gpg --batch --import "$GPG_PUBLIC"
            gpg --batch --import "$GPG_PRIVATE"

            # Get fingerprint
            FPR=$(gpg --list-keys --with-colons | grep '^fpr' | head -n1 | cut -d':' -f10)
            echo "$FPR" > fpr.txt

            # Trust key
            echo "$FPR:6:" > trust.txt
            gpg --import-ownertrust trust.txt

            # GPG config for loopback
            echo "use-agent" > ~/.gnupg/gpg.conf
            echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
            echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf

            # Restart gpg-agent
            gpgconf --kill gpg-agent
            export GPG_TTY=$(tty || true)
            gpgconf --launch gpg-agent
        '''

        def FPR = sh(script: 'cat fpr.txt', returnStdout: true).trim()
        sh """
            echo GELDİ
            echo GPG_PASSPHRASE: ${GPG_PASSPHRASE}
            echo FPR: ${FPR}
            echo debOutputPath: ${debOutputPath}
            
        """

        // sh '''
        //     gpg --batch --yes --pinentry-mode loopback \
        //         --passphrase "$GPG_PASSPHRASE" \
        //         -u "$FPR" \
        //         --output "${debOutputPath}.gpg" \
        //         --detach-sign "${debOutputPath}"

        // '''
        sh 'ls -lah "$debOutputPath"'

        sh 'echo "$GPG_PASSPHRASE" | dpkg-sig --sign builder -k "$GPG_FINGERPRINT" "$debOutputPath"''
        sh """
            echo "[*] Creating APT repo structure..."
            mkdir -p "${poolPath}" "${distPath}"
            cp "${debOutputPath}" "${poolPath}/"

            cd "${repoRoot}"
            dpkg-scanpackages pool /dev/null | gzip -9c > "${distPath}/Packages.gz"

            cd "${repoRoot}/dists/${DIST_CODENAME}"
            apt-ftparchive release . > Release

            echo "[*] Signing Release files..."
            gpg --batch --yes --pinentry-mode loopback \\
                --passphrase-fd 0 -u "\$FPR" -abs -o Release.gpg Release

            gpg --batch --yes --pinentry-mode loopback \\
                --passphrase-fd 0 -u "\$FPR" --clearsign -o InRelease Release

        """
        sh 'rm -rf ~/.gnupg trust.txt'
    }
}

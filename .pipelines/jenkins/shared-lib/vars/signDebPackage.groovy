def call(String OS, String ARCH, String DIST_CODENAME){
    withCredentials([
        file(credentialsId: 'private-rpm.gpg', variable: 'PRIVATE_GPG'),
        file(credentialsId: 'public-rpm.gpg', variable: 'PUBLIC_GPG'),
        string(credentialsId: 'GPG_PASSPHRASE', variable: 'GPG_PASSPHRASE'),
        string(credentialsId: 'GPG_FINGERPRINT', variable: 'GPG_FINGERPRINT')
    ]) {

        def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
        def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

        def repoRoot = "${env.WORKSPACE}/apt-repo"
        def poolPath = "${repoRoot}/pool/main/${APPNAME}"
        def distPath = "${repoRoot}/dists/${DIST_CODENAME}/main/binary-${ARCH}"

        sh """
            set -e

            echo "[*] Check .deb exists..."
            if [[ ! -f "${debOutputPath}" ]]; then
                echo "ERROR: .deb file not found at ${debOutputPath}"
                exit 1
            fi

            echo "[*] Setting up GPG environment..."
            export GNUPGHOME=\$(mktemp -d)
            chmod 700 "\$GNUPGHOME"

            echo "use-agent" > "\$GNUPGHOME/gpg.conf"
            echo "pinentry-mode loopback" >> "\$GNUPGHOME/gpg.conf"
            echo "allow-loopback-pinentry" > "\$GNUPGHOME/gpg-agent.conf"

            gpgconf --kill gpg-agent
            gpgconf --launch gpg-agent

            gpg --batch --import "${PRIVATE_GPG}"
            gpg --batch --import "${PUBLIC_GPG}"

            echo "[*] Trusting GPG key..."
            echo "${GPG_FINGERPRINT}:6:" > "\$GNUPGHOME/trust.txt"
            gpg --import-ownertrust "\$GNUPGHOME/trust.txt"

            echo "[*] Signing .deb with dpkg-sig..."
            GPG_TTY=\$(tty)
            export GPG_TTY
            echo "${GPG_PASSPHRASE}" | dpkg-sig --sign builder -k "${GPG_FINGERPRINT}" "${debOutputPath}"

            echo "[*] Creating APT repo structure..."
            mkdir -p "${poolPath}" "${distPath}"
            cp "${debOutputPath}" "${poolPath}/"

            cd "${repoRoot}"
            dpkg-scanpackages pool /dev/null | gzip -9c > "${distPath}/Packages.gz"

            cd "${repoRoot}/dists/${DIST_CODENAME}"
            apt-ftparchive release . > Release

            echo "[*] Signing Release files..."
            echo "${GPG_PASSPHRASE}" | gpg --batch --yes --pinentry-mode loopback \\
                --passphrase-fd 0 -u "${GPG_FINGERPRINT}" -abs -o Release.gpg Release

            echo "${GPG_PASSPHRASE}" | gpg --batch --yes --pinentry-mode loopback \\
                --passphrase-fd 0 -u "${GPG_FINGERPRINT}" --clearsign -o InRelease Release

            echo "[*] Cleaning up..."
            rm -rf "\$GNUPGHOME"
        """
    }
}

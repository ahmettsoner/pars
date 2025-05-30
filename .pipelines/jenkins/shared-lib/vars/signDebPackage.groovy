def call(String OS, String ARCH, String DIST_CODENAME) {
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

        // GNUPG setup and trust
        sh '''
            mkdir -p ~/.gnupg
            chmod 700 ~/.gnupg

            gpg --batch --import "$GPG_PUBLIC"
            gpg --batch --import "$GPG_PRIVATE"

            FPR=$(gpg --list-keys --with-colons | grep '^fpr' | head -n1 | cut -d':' -f10)
            echo "$FPR" > fpr.txt
            echo "$FPR:6:" > trust.txt
            gpg --import-ownertrust trust.txt

            echo "use-agent" > ~/.gnupg/gpg.conf
            echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
            echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf

            gpgconf --kill gpg-agent
            export GPG_TTY=$(tty || true)
            gpgconf --launch gpg-agent
        '''

        // Sign .deb using expect script
        sh '''
            cat > sign.expect << 'EOF'
#!/usr/bin/expect -f

set passphrase [lindex $argv 0]
set key [lindex $argv 1]
set deb [lindex $argv 2]

spawn dpkg-sig --sign builder -k $key $deb
expect {
    "Enter passphrase:" {
        send "$passphrase\\r"
    }
}
expect eof
EOF

            chmod +x sign.expect
            ./sign.expect "$GPG_PASSPHRASE" "$GPG_FINGERPRINT" "${debOutputPath}"
        '''

        // APT repo structure and signing
        sh """
            echo "[*] Creating APT repo structure..."
            mkdir -p "${poolPath}" "${distPath}"
            cp "${debOutputPath}" "${poolPath}/"

            cd "${repoRoot}"
            dpkg-scanpackages pool /dev/null | gzip -9c > "${distPath}/Packages.gz"

            cd "${repoRoot}/dists/${DIST_CODENAME}"
            apt-ftparchive release . > Release

            echo "[*] Signing Release files..."
            FPR=\$(cat ../../fpr.txt)

            echo "$GPG_PASSPHRASE" | gpg --batch --yes --pinentry-mode loopback \
                --passphrase-fd 0 -u "$FPR" -abs -o Release.gpg Release

            echo "$GPG_PASSPHRASE" | gpg --batch --yes --pinentry-mode loopback \
                --passphrase-fd 0 -u "$FPR" --clearsign -o InRelease Release
        """

        // Cleanup
        sh 'rm -rf ~/.gnupg trust.txt fpr.txt sign.expect'
    }
}

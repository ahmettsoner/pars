def call(String OS, String ARCH, String DIST_CODENAME) {
    unstash "${OS}-${ARCH}-deb-package-artifacts"

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
    def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
    def gpgIdentity = "ParsDevKit (Pars Repo Key) <support@parsdevkit.net>"
    def repoRoot = "${env.ARTIFACT_PATH}/apt-repo"
    def poolPath = "${repoRoot}/pool/main/${APPNAME}"

    withCredentials([
        file(credentialsId: 'public-rpm.gpg', variable: 'GPG_PUBLIC'),
        file(credentialsId: 'private-rpm.gpg', variable: 'GPG_PRIVATE'),
        string(credentialsId: 'GPG_PASSPHRASE', variable: 'GPG_PASSPHRASE')
    ]) {
        sh '''
        mkdir -p ~/.gnupg
        chmod 700 ~/.gnupg
        
        FPR=$(gpg --list-secret-keys --with-colons | awk -F: '/^fpr/ { print $10; exit }')

        # GPG agent yapılandır
        echo "use-agent" > ~/.gnupg/gpg.conf
        echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
        echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf

        gpgconf --kill gpg-agent
        gpgconf --launch gpg-agent
        export GPG_TTY=$(tty || true)

        # Trust ve anahtar import işlemleri
        gpg --batch --import "$GPG_PUBLIC"
        gpg --batch --import "$GPG_PRIVATE"
        echo "$FPR:6:" > trust.txt
        gpg --import-ownertrust trust.txt

        # Parolayı önceden cache’e al
        echo "$GPG_PASSPHRASE" | /usr/lib/gnupg/gpg-preset-passphrase --preset "$FPR"
        '''

        // dpkg-sig ile imzala
        sh """
        dpkg-sig -k "${gpgIdentity}" --sign builder "${debOutputPath}"
        """

        // Prepare pool path and copy signed .deb
        sh """
        mkdir -p ${poolPath}
        cp ${debOutputPath} ${poolPath}/
        """

        def distPath = "${repoRoot}/dists/${DIST_CODENAME}/main/binary-${ARCH}"
        sh """
        mkdir -p ${distPath}
        cd ${poolPath}
        dpkg-scanpackages . /dev/null | gzip -9c > ${distPath}/Packages.gz

        cd ${repoRoot}/dists/${DIST_CODENAME}
        apt-ftparchive release . > Release
        gpg --batch --yes --passphrase "${GPG_PASSPHRASE}" --pinentry-mode loopback -u "${gpgIdentity}" -abs -o Release.gpg Release
        gpg --batch --yes --passphrase "${GPG_PASSPHRASE}" --pinentry-mode loopback -u "${gpgIdentity}" --clearsign -o InRelease Release
        """

        sh 'rm -rf ~/.gnupg trust.txt'
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}

return this

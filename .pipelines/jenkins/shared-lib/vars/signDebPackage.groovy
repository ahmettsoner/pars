def call(String DIST, String CODENAME, String ARCH) {
    def gpgIdentity = "ParsDevKit (Pars Repo Key) <support@parsdevkit.net>"
    def debFileName = "${APPNAME}_${VERSION}_${ARCH}.deb"
    def repoRoot = "${env.ARTIFACT_PATH}/apt"

    def poolPath = "${repoRoot}/pool/main/p/${APPNAME}"
    def distPath = "${repoRoot}/dists/${CODENAME}/main/binary-${ARCH}"

    sh "mkdir -p ${poolPath} ${distPath}"
    sh "cp ${env.ARTIFACT_PATH}/${debFileName} ${poolPath}/"

    withCredentials([
        file(credentialsId: 'public-deb.gpg', variable: 'GPG_PUBLIC'),
        file(credentialsId: 'private-deb.gpg', variable: 'GPG_PRIVATE'),
        string(credentialsId: 'GPG_PASSPHRASE', variable: 'GPG_PASSPHRASE')
    ]) {
        sh '''
        mkdir -p ~/.gnupg
        chmod 700 ~/.gnupg

        gpg --batch --import "$GPG_PUBLIC"
        gpg --batch --import "$GPG_PRIVATE"

        FPR=$(gpg --list-keys --with-colons | grep '^fpr' | head -n1 | cut -d':' -f10)
        echo "$FPR:6:" > trust.txt
        gpg --import-ownertrust trust.txt

        echo "use-agent" > ~/.gnupg/gpg.conf
        echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
        echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf

        gpgconf --kill gpg-agent
        export GPG_TTY=$(tty || true)
        gpgconf --launch gpg-agent
        '''

        // Generate Packages.gz
        sh """
        cd ${repoRoot}
        dpkg-scanpackages --arch ${ARCH} pool > ${distPath}/Packages
        gzip -kf ${distPath}/Packages
        """

        // Create Release file
        sh """
        cd ${distPath}
        apt-ftparchive release . > Release
        """

        // Sign the Release file
        sh """
        cd ${distPath}
        gpg --batch --yes --pinentry-mode loopback --passphrase "${GPG_PASSPHRASE}" -u "${gpgIdentity}" -abs -o Release.gpg Release
        gpg --batch --yes --pinentry-mode loopback --passphrase "${GPG_PASSPHRASE}" -u "${gpgIdentity}" --clearsign -o InRelease Release
        """

        // Cleanup
        sh 'rm -rf ~/.gnupg ~/.rpmmacros trust.txt'
    }

    stash includes: 'apt/**/*', name: "${DIST}-${ARCH}-deb-repo"
}

return this

def call(String OS, String ARCH, List<String> DIST_CODENAMES) {
    unstash "${OS}-${ARCH}-deb-package-artifacts"

    def newBaseName = "${APPNAME}_${APPVERSION}_${ARCH}.deb"
    def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
    def gpgIdentity = "ParsDevKit (Pars Repo Key) <support@parsdevkit.net>"
    def repoRoot = "${env.ARTIFACT_PATH}/apt-repo"
    def poolPath = "${repoRoot}/pool/main/${APPNAME}"

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

        // Sign the .deb file
        sh """
        echo "${GPG_PASSPHRASE}" | dpkg-sig -k "${gpgIdentity}" --sign builder "${debOutputPath}"
        """

        // Prepare pool path and copy signed .deb
        sh """
        mkdir -p ${poolPath}
        cp ${debOutputPath} ${poolPath}/
        """

        // Loop through each codename and build Packages, Release, etc.
        DIST_CODENAMES.each { codename ->
            def distPath = "${repoRoot}/dists/${codename}/main/binary-${ARCH}"
            sh """
            mkdir -p ${distPath}
            cd ${poolPath}
            dpkg-scanpackages . /dev/null | gzip -9c > ${distPath}/Packages.gz

            cd ${repoRoot}/dists/${codename}
            apt-ftparchive release . > Release
            gpg --batch --yes --passphrase "${GPG_PASSPHRASE}" --pinentry-mode loopback -u "${gpgIdentity}" -abs -o Release.gpg Release
            gpg --batch --yes --passphrase "${GPG_PASSPHRASE}" --pinentry-mode loopback -u "${gpgIdentity}" --clearsign -o InRelease Release
            """
        }

        // Cleanup GPG
        sh 'rm -rf ~/.gnupg trust.txt'
    }

    // Re-stash final APT repo
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}

return this

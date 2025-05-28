def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

    // GPG Key yükleme ve imzalama
    withCredentials([
        file(credentialsId: 'private-rpm.gpg', variable: 'GPG_PRIVATE_KEY_FILE'),
        file(credentialsId: 'public-rpm.gpg', variable: 'GPG_PUBLIC_KEY_FILE')
    ]) {
        // GPG anahtarlarını import et
        sh '''
        mkdir -p ~/.gnupg
        chmod 700 ~/.gnupg

        gpg --batch --import "$GPG_PUBLIC_KEY_FILE"
        gpg --batch --import "$GPG_PRIVATE_KEY_FILE"
        '''

        // Fingerprint al
        def keyFpr = sh(script: "gpg --list-secret-keys --with-colons | awk -F: '/^fpr/ { print \$10; exit }'", returnStdout: true).trim()

        // Güven düzeyini ultimate (6) yap
        writeFile file: 'trust.txt', text: "${keyFpr}:6:\n"
        sh "gpg --import-ownertrust trust.txt"

        // .rpmmacros dosyasını yaz
        sh """
        echo '%_signature gpg' > ~/.rpmmacros
        echo '%_gpg_name ParsDevKit (Pars Repo Key) <support@parsdevkit.net>' >> ~/.rpmmacros
        """

        // RPM paketini imzala
        sh "rpm --addsign ${rpmOutputPath}"
    }

    // Temizlik
    sh 'rm -rf ~/.gnupg ~/.rpmmacros trust.txt'

    // Tekrar stash
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

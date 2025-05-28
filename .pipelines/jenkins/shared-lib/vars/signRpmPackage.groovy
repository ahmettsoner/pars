def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

    // Secret file kullanımı (gpg-private-key ID'li secret file)
    withCredentials([file(credentialsId: 'public-rpm.gpg', variable: 'GPG_KEY_FILE')]) {
        // GPG key import
        sh "gpg --batch --import ${GPG_KEY_FILE}"

        // Get key fingerprint
        def keyFpr = sh(script: "gpg --list-keys --with-colons | grep fpr | head -n1 | cut -d':' -f10", returnStdout: true).trim()

        // Set trust level to ultimate (trust level 6)
        writeFile file: 'trust.txt', text: "${keyFpr}:6:\n"
        sh "gpg --import-ownertrust trust.txt"

        // GPG agent ayarları
        sh '''
        echo "use-agent" >> ~/.gnupg/gpg.conf
        echo "allow-loopback-pinentry" >> ~/.gnupg/gpg-agent.conf
        gpgconf --kill gpg-agent
        export GPG_TTY=$(tty || true)
        '''
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
    sh 'rm -rf ~/.gnupg trust.txt'
}

return this

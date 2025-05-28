def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

    // Secret file kullanımı (gpg-private-key ID'li secret file)
    withCredentials([file(credentialsId: 'gpg-private-key', variable: 'GPG_KEY_FILE')]) {
        sh "gpg --batch --import ${GPG_KEY_FILE}"
        sh '''
        echo "%_signature gpg" > ~/.rpmmacros
        echo "%_gpg_name Your Name <your.email@example.com>" >> ~/.rpmmacros
        '''
        sh "rpm --addsign ${rpmOutputPath}"
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

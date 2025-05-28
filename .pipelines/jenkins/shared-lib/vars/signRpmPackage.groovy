def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

    withCredentials([
        file(credentialsId: 'private-rpm.gpg', variable: 'GPG_PRIVATE_KEY_FILE'),
        file(credentialsId: 'public-rpm.gpg', variable: 'GPG_PUBLIC_KEY_FILE')
    ]) {
        sh '''
        mkdir -p ~/.gnupg
        chmod 700 ~/.gnupg

        gpg --batch --import "$GPG_PUBLIC_KEY_FILE"
        gpg --batch --import "$GPG_PRIVATE_KEY_FILE"

        # GPG agent yapılandır
        echo "use-agent" > ~/.gnupg/gpg.conf
        echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
        echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf

        gpgconf --kill gpg-agent
        export GPG_TTY=$(tty || true)
        gpgconf --launch gpg-agent
        '''

        def keyFpr = sh(script: "gpg --list-secret-keys --with-colons | awk -F: '/^fpr/ { print \$10; exit }'", returnStdout: true).trim()

        writeFile file: 'trust.txt', text: "${keyFpr}:6:\n"
        sh "gpg --import-ownertrust trust.txt"

        // rpmmacros yaz
        sh """
        echo '%_signature gpg' > ~/.rpmmacros
        echo '%_gpg_name ParsDevKit (Pars Repo Key) <support@parsdevkit.net>' >> ~/.rpmmacros
        echo '%__gpg /usr/bin/gpg' >> ~/.rpmmacros
        echo '%__gpg_check_password_cmd %{nil}' >> ~/.rpmmacros
        echo '%__gpg_sign_cmd %{__gpg} \\\n\
        --batch --verbose --pinentry-mode loopback \\\n\
        --no-armor --no-tty --yes \\\n\
        -u "%{_gpg_name}" -sbo %{__signature_filename} %{__plaintext_filename}' >> ~/.rpmmacros
        """

        // RPM imzalama
        sh "rpm --addsign ${rpmOutputPath}"
    }

    sh 'rm -rf ~/.gnupg ~/.rpmmacros trust.txt'

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

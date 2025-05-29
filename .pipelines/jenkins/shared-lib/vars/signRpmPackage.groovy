def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
    def gpgIdentity = "ParsDevKit (Pars Repo Key) <support@parsdevkit.net>"

    withCredentials([
        file(credentialsId: 'public-rpm.gpg', variable: 'GPG_PUBLIC'),
        file(credentialsId: 'private-rpm.gpg', variable: 'GPG_PRIVATE'),
        string(credentialsId: 'GPG_PASSPHRASE', variable: 'GPG_PASSPHRASE'),
        string(credentialsId: 'GPG_FINGERPRINT', variable: 'GPG_FINGERPRINT')
    ]) {
        // Prepare ~/.rpmmacros (DO NOT use %% here, just single %)
        writeFile file: "${env.HOME}/.rpmmacros", text: """
%_signature gpg
%_gpg_name ${gpgIdentity}
%__gpg /usr/bin/gpg
%__gpg_sign_cmd %{__gpg} \\
  --batch \\
  --yes \\
  --no-armor \\
  --pinentry-mode loopback \\
  --passphrase "${GPG_PASSPHRASE}" \\
  -u "%{_gpg_name}" \\
  -sbo %{__signature_filename} %{__plaintext_filename}
"""

        sh """
        mkdir -p ~/.gnupg
        chmod 700 ~/.gnupg

        # Import GPG keys
        gpg --batch --import "$GPG_PUBLIC"
        gpg --batch --import "$GPG_PRIVATE"

        # Trust key
        echo "${GPG_FINGERPRINT}:6:" > trust.txt
        gpg --import-ownertrust trust.txt

        # GPG config for loopback
        echo "use-agent" > ~/.gnupg/gpg.conf
        echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
        echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf

        # Restart gpg-agent
        gpgconf --kill gpg-agent
        export GPG_TTY=$(tty || true)
        gpgconf --launch gpg-agent
        """

        // Sign the RPM
        sh "rpm --addsign ${rpmOutputPath}"

        // Clean up GPG traces
        sh 'rm -rf ~/.gnupg ~/.rpmmacros trust.txt'
    }

    // Re-stash signed package
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

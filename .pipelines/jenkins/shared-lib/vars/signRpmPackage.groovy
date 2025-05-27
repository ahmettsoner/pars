def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"



    
    sh "echo \"${GPG_PRIVATE_KEY_B64}\" | base64 -d > /tmp/private.gpg"
    sh "gpg --batch --import /tmp/private.gpg"

    sh "KEY_FPR=$(gpg --list-keys --with-colons | grep fpr | head -n1 | cut -d':' -f10)"
    sh "echo -e \"5\ny\n\" | gpg --command-fd 0 --expert --edit-key \"$KEY_FPR\" trust"
    sh "echo \"use-agent\" >> ~/.gnupg/gpg.conf"
    sh "echo \"allow-loopback-pinentry\" >> ~/.gnupg/gpg-agent.conf"
    sh "gpgconf --kill gpg-agent"
    sh "GPG_TTY=$(tty)"
    sh "export GPG_TTY"

    sh "rpm --define \"%_gpg_name ${GPG_SIGN_KEY}\" --addsign ${rpmOutputPath}"
    sh "rpm -K ${rpmOutputPath}"

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

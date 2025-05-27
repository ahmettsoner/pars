def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"



    
    sh "echo \"${GPG_PRIVATE_KEY_B64}\" | base64 -d > /tmp/private.gpg"
    sh "gpg --batch --import /tmp/private.gpg"

    sh "KEY_FPR=\"$(gpg --list-keys --with-colons | grep fpr | head -n1 | cut -d':' -f10)\""
    sh "echo -e \"5\ny\n\" | gpg --command-fd 0 --expert --edit-key \"$KEY_FPR\" trust"
    sh "echo \"use-agent\" >> ~/.gnupg/gpg.conf"
    sh "echo \"allow-loopback-pinentry\" >> ~/.gnupg/gpg-agent.conf"
    sh "gpgconf --kill gpg-agent"
    sh "GPG_TTY=$(tty)"
    sh "export GPG_TTY"

    // sh "curl -v -u \"$NEXUS_USER:$NEXUS_PASS\" --upload-file ~/rpmbuild/RPMS/x86_64/my-package.rpm \"$NEXUS_URL/my-package.rpm\""

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

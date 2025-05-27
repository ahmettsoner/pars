def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

    sh "echo \"%_gpg_name ${GPG_SIGN_KEY}\" >> ~/.rpmmacros"
    sh "rpm --addsign ${rpmOutputPath}"
    sh "rpm -K ${rpmOutputPath}"

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

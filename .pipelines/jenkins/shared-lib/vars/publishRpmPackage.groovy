def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"




    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

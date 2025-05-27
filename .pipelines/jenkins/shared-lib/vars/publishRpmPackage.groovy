def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"


    sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file ${rpmOutputPath} \"${NEXUS_URL}/${APPNAME}.rpm\""


    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

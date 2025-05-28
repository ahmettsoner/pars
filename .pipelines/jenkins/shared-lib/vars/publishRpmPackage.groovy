def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
    def RPM_FILENAME = "${APPNAME}-${CURRENT_BASE_VERSION_RAW}-${BUILD_VERSION_RELEASE_NUMBER}.${ARCH}.rpm"


    withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file ${rpmOutputPath} \"${NEXUS_URL}/yum2/${RPM_FILENAME}\""
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

def call(String OS, String ARCH, String DIST_CODENAME) {
    unstash "${OS}-${ARCH}-deb-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
    def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
    def DEB_FILENAME = "${APPNAME}-${CURRENT_BASE_VERSION_RAW}-${BUILD_VERSION_RELEASE_NUMBER}.${ARCH}.deb"

    // Nexus APT upload path:
    def firstChar = APPNAME[0].toLowerCase()
    def uploadUrl = "${NEXUS_URL}/apt-${DIST_CODENAME}/pool/main/${firstChar}/${APPNAME}/${DEB_FILENAME}"
    
    echo "Uploading DEB package: ${debOutputPath}"
    echo "Upload URL: ${uploadUrl}"
    
    withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file ${debOutputPath} \"${uploadUrl}\""
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}

return this

def call(String OS, String ARCH, String DIST_CODENAME) {
    unstash "${OS}-${ARCH}-deb-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
    def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
    def VERSION = env.CURRENT_BASE_VERSION_RAW + (env.CURRENT_VERSION_RELEASE_NUMBER ? "-${env.CURRENT_VERSION_RELEASE_NUMBER}" : "")

    def utils = new com.parsdevkit.Utils(this)
    def platformArch = utils.mapArch("debian", ARCH)
    def DEB_FILENAME = "${APPNAME}_${VERSION}_${platformArch}.deb"
  
    // Nexus APT upload path:
    def firstChar = APPNAME[0].toLowerCase()
    def uploadUrl = "${NEXUS_URL}/repository/apt-${DIST_CODENAME}/pool/${firstChar}/${APPNAME}/${DEB_FILENAME}"
    
    echo "Uploading DEB package: ${debOutputPath}"
    echo "Upload URL: ${uploadUrl}"
    
    withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
        // sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file ${debOutputPath} \"${uploadUrl}\""
        sh """
            curl -v -u "${NEXUS_USER}:${NEXUS_PASS}" \\
                -H "Content-Type: multipart/form-data" \\
                --data-binary "@${debOutputPath}" \\
                "${NEXUS_URL}/repository/apt-${DIST_CODENAME}/"
        """
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}

return this

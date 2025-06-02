def call(String OS, String ARCH, String DIST_CODENAME) {
    unstash "${OS}-${ARCH}-deb-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
    def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
    def VERSION = env.CURRENT_BASE_VERSION_RAW + (env.BUILD_VERSION_RELEASE_NUMBER ? "-${env.BUILD_VERSION_RELEASE_NUMBER}" : "")

    def debArch = ""
    switch (ARCH) {
        case "x86":
            debArch = "i386"
            break
        case "x86_64":
            debArch = "amd64"
            break
        case "arm":
            debArch = "armhf"
            break
        case "arm64":
            debArch = "arm64"
            break
        default:
            error "Unsupported architecture: ${debArch}"
    }
    def DEB_FILENAME = "${APPNAME}_${VERSION}_${debArch}.deb"
  
    // Nexus APT upload path:
    def firstChar = APPNAME[0].toLowerCase()
    def uploadUrl = "${NEXUS_URL}/apt-${DIST_CODENAME}/pool/${firstChar}/${APPNAME}/${DEB_FILENAME}"
    
    echo "Uploading DEB package: ${debOutputPath}"
    echo "Upload URL: ${uploadUrl}"
    
    withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file ${debOutputPath} \"${uploadUrl}\""
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}

return this

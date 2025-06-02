def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

    def rpmArch = ""
    switch (ARCH) {
        case "x86":
            rpmArch = "i386"
            break
        case "x86_64":
            rpmArch = "x86_64"
            break
        case "arm":
            rpmArch = "armv7hl"
            break
        case "arm64":
            rpmArch = "aarch64"
            break
        default:
            error "Unsupported architecture: ${ARCH}"
    }
    def RPM_FILENAME = "${APPNAME}-${CURRENT_BASE_VERSION_RAW}-${BUILD_VERSION_RELEASE_NUMBER}.${rpmArch}.rpm"


    withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file ${rpmOutputPath} \"${NEXUS_URL}/repository/yum/${RPM_FILENAME}\""
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-msi-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.msi"
    def msiOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

    def utils = new com.parsdevkit.Utils(this)
    def platformArch = utils.mapArch("windows", ARCH)
    def MSI_FILENAME = "${APPNAME}-${CURRENT_VERSION}.${platformArch}.msi"

    withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${msiOutputPath}\" \"${NEXUS_URL}/repository/msi-${CHANNEL}/${MSI_FILENAME}\""
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-msi-package-artifacts"
}

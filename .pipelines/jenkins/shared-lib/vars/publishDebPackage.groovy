def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-deb-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
    def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

    withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
        sh """
            curl -v -u "${NEXUS_USER}:${NEXUS_PASS}" \\
                -H "Content-Type: multipart/form-data" \\
                --data-binary "@${debOutputPath}" \\
                "${NEXUS_URL}/repository/apt-${CHANNEL}/"
        """
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}

return this

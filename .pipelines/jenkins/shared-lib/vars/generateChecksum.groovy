def call(String OS, String ARCH) {
    echo "Build Version: ${env.BUILD_VERSION}"

    env.ARTIFACT_CHECKSUM_MD5_PATH = "${env.ARTIFACT_PATH}/Checksums.md"
    if (!env.BUILD_VERSION) {
        error "BUILD_VERSION is empty or null"
    }

    sh """
        echo "| Platform | Architecture | Type      | File Name | SHA-256 Checksum |" > ${env.ARTIFACT_CHECKSUM_MD5_PATH}
        echo "|----------|--------------|-----------|-----------|------------------|" >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}
    """


    def ext = ""
    if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()

    def type = "Binary"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"
    sh "echo '${line}' >> ${env.ARTIFACT_PATH}/${newBaseName}-checksum.txt"

    archiveArtifacts artifacts: 'dist/artifacts/**/*', fingerprint: true
}

return this
def call(String OS, String ARCH) {
    def ext = ""
    if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()

    def type = "Binary"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"
    sh "echo '${line}' >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}"

    archiveArtifacts artifacts: 'dist/artifacts/**/*', fingerprint: true
}

return this
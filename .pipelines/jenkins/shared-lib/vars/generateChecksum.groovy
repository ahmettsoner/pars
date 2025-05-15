def WriteChecksumForBinary(OS, ARCH) {
    unstash "${OS}-${ARCH}-artifacts"

    def ext = OS.toLowerCase() == 'windows' ? '.exe' : ''
    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()
    def type = "Binary"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"

    sh "echo \"${line}\" >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}"
}

def WriteChecksumForArchive(OS, ARCH) {
    unstash "${OS}-${ARCH}-archive-artifacts"

    def archiveFormat = ""
    switch (OS) {
        case "windows":
            archiveFormat = "zip"
            break
        case "linux":
            archiveFormat = "tar.gz"
            break
        default:
            archiveFormat = "tar.gz"
    }
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.bin.${archiveFormat}"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()
    def type = "Archive"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"

    sh "echo \"${line}\" >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}"
}

return this

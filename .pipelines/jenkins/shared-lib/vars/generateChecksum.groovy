def WriteChecksumForBinary(OS, ARCH) {
    unstash "${OS}-${ARCH}-artifacts"

    def utils = new com.parsdevkit.Utils(this)
    def ext = utils.appExt(OS)
    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()
    def type = "Binary"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"

    sh "echo \"${line}\" >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}"
}

def WriteChecksumForArchive(OS, ARCH) {
    unstash "${OS}-${ARCH}-archive-artifacts"


    def utils = new com.parsdevkit.Utils(this)
    def archiveFormat = utils.archiveFormat(OS)
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.bin.${archiveFormat}"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()
    def type = "Archive"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"

    sh "echo \"${line}\" >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}"
}
def WriteChecksumForDEBPackage(ARCH) {
    def OS = 'linux'
    unstash "${OS}-${ARCH}-deb-package-artifacts"

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()
    def type = "DEB"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"

    sh "echo \"${line}\" >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}"
}
def WriteChecksumForRPMPackage(ARCH) {
    def OS = 'linux'
    unstash "${OS}-${ARCH}-rpm-package-artifacts"

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()
    def type = "RPM"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"

    sh "echo \"${line}\" >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}"
}
def WriteChecksumForMSIInstaller(ARCH) {
    def OS = 'windows'
    unstash "${OS}-${ARCH}-msi-package-artifacts"

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.msi"
    def artifactPath = "${env.ARTIFACT_PATH}\\${newBaseName}"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()

    def type = "MSI"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"

    powershell """
        Add-Content -Path '${env.ARTIFACT_CHECKSUM_MD5_PATH}' -Value '${line}'
    """
}

return this

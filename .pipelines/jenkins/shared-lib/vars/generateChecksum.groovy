def WriteChecksumeForBinary(String OS, String ARCH) {
    def ext = OS.toLowerCase() == 'windows' ? '.exe' : ''
    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"
    def checksum = sh(script: "sha256sum ${env.ARTIFACT_PATH}/${newBaseName} | awk '{print \$1}'", returnStdout: true).trim()
    def type = "Binary"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"

    // Write to a temporary file
    writeFile file: "checksum-${OS}-${ARCH}.txt", text: "${line}\n"
    stash includes: "checksum-${OS}-${ARCH}.txt", name: "checksum-${OS}-${ARCH}"
}
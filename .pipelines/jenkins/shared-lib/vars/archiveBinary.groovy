def call(String OS, String ARCH, String versionFilePath) {
    def versionOutput = readFile(versionFilePath)
    
    def appName = versionOutput.split('\n').find { it.startsWith('APPNAME=') }?.split('=')[1]?.trim()
    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()

    def extLine = versionOutput.split('\n').find { it.startsWith('EXT=') }
    def ext = ""
    
    if (extLine?.contains('=') && extLine.split('=').length > 1) {
        ext = extLine.split('=')[1].trim()
    } else if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    def originalFileName = "${appName}${ext}"
    def newBaseName = "${appName}-${OS}-${ARCH}${ext}"

    def binaryOutputPathBase = "${WORKSPACE}/dist/${buildVersion}/${OS}/bin/${ARCH}"
    def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"
    def binaryChecksumPath = "${binaryOutputPathBase}/checksum.txt"

    def artifactPath = "dist/artifacts/${buildVersion}"

    // Hedef klasörü oluştur
    sh "mkdir -p ${artifactPath}"

    // Dosyayı kopyala
    sh "cp ${binaryOutputPathBase}/${originalFileName} ${artifactPath}/${newBaseName}"

    // Checksum oku
    def binaryChecksum = readFile("${binaryChecksumPath}").trim()

    // Checksums.md'ye yaz
    def type = "Binary"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${binaryChecksum} |"
    sh "echo '${line}' >> ${artifactPath}/Checksums.md"

    // Stash dosyalar
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-artifacts"
}

return this

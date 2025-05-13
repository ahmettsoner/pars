def archiveBinary(String OS, String ARCH) {
    def ext = ""
    if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    def originalFileName = "${APPNAME}${ext}"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"

    def binaryOutputPathBase = "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}"
    def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"
    def binaryChecksumPath = "${binaryOutputPathBase}/checksum.txt"

    def artifactPath = "dist/artifacts/${env.BUILD_VERSION}"

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

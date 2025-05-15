def call(String OS, String ARCH) {
    def ext = ""
    if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    def originalFileName = "${APPNAME}${ext}"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"

    def binaryOutputPathBase = "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}"
    def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"
    def binaryChecksumPath = "${binaryOutputPathBase}/checksum.txt"

    // Hedef klasörü oluştur
    sh "mkdir -p ${env.ARTIFACT_PATH}"

    // Dosyayı kopyala
    sh "cp ${binaryOutputPathBase}/${originalFileName} ${env.ARTIFACT_PATH}/${newBaseName}"

    // // Checksum oku
    // def checksum = readFile("${binaryChecksumPath}").trim()

    // // Checksums.md'ye yaz
    // def type = "Binary"
    // def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"
    // sh "echo '${line}' >> ${env.ARTIFACT_PATH}/${newBaseName}-checksum.txt"

    // Stash dosyalar
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-artifacts"
    stash includes: "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}/**/*", name: "${OS}-${ARCH}-dist-bin"
}

return this

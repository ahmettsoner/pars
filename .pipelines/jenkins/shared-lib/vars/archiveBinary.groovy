def call(String OS, String ARCH) {
    def ext = OS.toLowerCase() == 'windows' ? '.exe' : ''

    def originalFileName = "${APPNAME}${ext}"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"

    unstash "${OS}-${ARCH}-dist-bin"
    def binaryOutputPathBase = "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}"

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
}

return this

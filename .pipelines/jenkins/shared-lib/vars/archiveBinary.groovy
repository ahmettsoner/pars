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

    // Stash dosyalar
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-artifacts"
}

return this

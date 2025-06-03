def call(String OS, String ARCH) {
    def utils = new com.parsdevkit.Utils(this)
    def platformArch = utils.mapArch("debian", ARCH)

    def plainVersion = env.BUILD_VERSION.replaceFirst(/^v/, "")
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
    def packageOutputBase = "dist/${env.BUILD_VERSION}/${OS}/pkg/deb/${ARCH}"
    def debOutputBase = "${packageOutputBase}/output"
    def debOutputPath = "${debOutputBase}/${APPNAME}_${plainVersion}_${platformArch}.deb"
    def checksumFilePath = "${debOutputBase}/checksum.txt"
    // Dosya var mı kontrolü
    if (!fileExists(debOutputPath)) {
        error "No DEB file found at ${debOutputPath}"
    }

    // Artifact dizinine kopyala
    sh "cp '${debOutputPath}' '${env.ARTIFACT_PATH}/${newBaseName}'"

    // Stash işlemi
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}


return this
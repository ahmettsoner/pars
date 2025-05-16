def call(String OS, String ARCH) {
    def debArch = ""
    switch (ARCH) {
        case "x86":
            debArch = "i386"
            break
        case "x86_64":
            debArch = "amd64"
            break
        case "arm":
            debArch = "armhf"
            break
        case "arm64":
            debArch = "arm64"
            break
        default:
            error "Unsupported architecture: ${ARCH}"
    }

    def plainVersion = env.BUILD_VERSION.replaceFirst(/^v/, "")
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
    def packageOutputBase = "dist/${env.BUILD_VERSION}/${OS}/pkg/deb/${ARCH}"
    def debOutputBase = "${packageOutputBase}/output"
    def debOutputPath = "${debOutputBase}/${APPNAME}_${plainVersion}_${debArch}.deb"
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
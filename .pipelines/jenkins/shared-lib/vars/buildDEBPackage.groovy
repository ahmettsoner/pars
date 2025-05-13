def call(String OS, String ARCH) {
    def binaryOutputPathBase = "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}"

    def ext = ""
    if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }
    def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"

    def pckgPath = "usr/bin"
    def packageOutputBase = "dist/${env.BUILD_VERSION}/${OS}/pkg/deb/${ARCH}"
    def packageBinaryPath = "${packageOutputBase}/${APPNAME}/${pckgPath}"
    // def newBaseName = "${APPNAME}-${env.CURRENT_BASE_VERSION}.tar.gz"
    // def originalFileName = "${APPNAME}${ext}"

    sh """
        mkdir -p '${packageBinaryPath}'
        cp '${binaryOutputPath}' '${packageBinaryPath}'
        export GO111MODULE=on
        make build.deb.package.${ARCH}.configuration VERSION=${env.BUILD_VERSION}
        make build.deb.package.${ARCH}.package VERSION=${env.BUILD_VERSION}
    """

    def debArch = ""
    if (ARCH == "x86") {
        debArch = "i386"
    } else if (ARCH == "x86_64") {
        debArch = "amd64"
    } else if (ARCH == "arm") {
        debArch = "armhf"
    } else if (ARCH == "arm64") {
        debArch = "arm64"
    } else {
        error "Unsupported architecture: ${ARCH}"
    }

    def plainVersion = env.BUILD_VERSION.replaceFirst(/^v/, "")
    def debOutputBase = "${packageOutputBase}/output"
    def debOutputPath = "${debOutputBase}/pars_${plainVersion}_${debArch}.deb"
    def debChecksumPath = "${debOutputBase}/checksum.txt"

    def debChecksum = sh(script: "sha256sum '${debOutputPath}' | awk '{print \$1}'", returnStdout: true).trim()
    echo "Checksum for ${debOutputPath}: ${debChecksum}"

    writeFile file: debChecksumPath, text: debChecksum

}

def copyDebAndUpdateChecksums(String OS, String ARCH) {
    def artifactPath = "dist/artifacts/${env.BUILD_VERSION}"

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
    def checksumsMdPath = "${artifactPath}/Checksums.md"

    // Dosya var mı kontrolü
    if (!fileExists(debOutputPath)) {
        error "No DEB file found at ${debOutputPath}"
    }

    // Artifact dizinine kopyala
    sh "mkdir -p '${artifactPath}'"
    sh "cp '${debOutputPath}' '${artifactPath}/${newBaseName}'"

    // Checksum oku
    def checksum = readFile(checksumFilePath).trim()

    // Checksums.md'ye satır ekle
    def line = "| ${OS} | ${ARCH} | DEB | ${newBaseName} | ${checksum} |"
    // writeFile file: checksumsMdPath, text: "${line}\n", encoding: "UTF-8", append: true
    sh "echo '${line}' >> ${artifactPath}/Checksums.md"

    echo "Added checksum line to Checksums.md: ${line}"

    // Stash işlemi
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}


return this
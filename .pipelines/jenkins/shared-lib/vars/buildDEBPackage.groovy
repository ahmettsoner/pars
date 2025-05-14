def call(String OS, String ARCH) {

    unstash "${OS}-${ARCH}-dist-bin"
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

return this
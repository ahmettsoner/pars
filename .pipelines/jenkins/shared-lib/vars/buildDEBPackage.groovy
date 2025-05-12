def call(String OS, String ARCH) {
    def versionFilePath = "version_output.txt"
    def versionOutput = readFile(versionFilePath)

    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()
    def baseVersion = versionOutput.split('\n').find { it.startsWith('CURRENT_BASE_VERSION=') }?.split('=')[1]?.trim()

    def appName = "pars"
    def binaryOutputPathBase = "dist/${buildVersion}/${OS}/bin/${ARCH}"

    def extLine = versionOutput.split('\n').find { it.startsWith('EXT=') }
    def ext = ""
    if (extLine?.contains('=') && extLine.split('=').length > 1) {
        ext = extLine.split('=')[1].trim()
    } else if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }
    def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"

    def pckgPath = "usr/bin"
    def packageOutputBase = "dist/${buildVersion}/${OS}/pkg/deb/${ARCH}"
    def packageBinaryPath = "${packageOutputBase}/${appName}/${pckgPath}"
    // def newBaseName = "${appName}-${baseVersion}.tar.gz"
    // def originalFileName = "${appName}${ext}"

    sh """
        mkdir -p '${packageBinaryPath}'
        cp '${binaryOutputPath}' '${packageBinaryPath}'
        export GO111MODULE=on
        make build.deb.package.${ARCH}.configuration VERSION=${buildVersion}
        make build.deb.package.${ARCH}.package VERSION=${buildVersion}
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

    def plainVersion = version.replaceFirst(/^v/, "")
    def debOutputBase = "${packageOutputBase}/output"
    def debOutputPath = "${debOutputBase}/pars_${plainVersion}_${debArch}.deb"
    def debChecksumPath = "${debOutputBase}/checksum.txt"

    def debChecksum = sh(script: "sha256sum '${debOutputPath}' | awk '{print \$1}'", returnStdout: true).trim()
    echo "Checksum for ${debOutputPath}: ${debChecksum}"

    writeFile file: debChecksumPath, text: debChecksum

}

def copyDebAndUpdateChecksums(String OS, String ARCH) {
    def versionFilePath = "version_output.txt"
    def versionOutput = readFile(versionFilePath)

    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()
    def baseVersion = versionOutput.split('\n').find { it.startsWith('CURRENT_BASE_VERSION=') }?.split('=')[1]?.trim()
    def rawBaseVersion = baseVersion?.startsWith('v') ? baseVersion.substring(1) : baseVersion
    def appName = "pars"
    def artifactPath = "dist/artifacts/${buildVersion}"

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

    def plainVersion = rawBaseVersion
    def newBaseName = "${appName}-${OS}-${ARCH}.deb"
    def packageOutputBase = "dist/${buildVersion}/${OS}/pkg/deb/${ARCH}"
    def debOutputBase = "${packageOutputBase}/output"
    def debOutputPath = "${debOutputBase}/${appName}_${plainVersion}_${debArch}.deb"
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
    def checksumLine = "| ${OS} | ${ARCH} | DEB | ${newBaseName} | ${checksum} |"
    writeFile file: checksumsMdPath, text: "${checksumLine}\n", encoding: "UTF-8", append: true

    echo "Added checksum line to Checksums.md: ${checksumLine}"

    // Stash işlemi
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
}


return this
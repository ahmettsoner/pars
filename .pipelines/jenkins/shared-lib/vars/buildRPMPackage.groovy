def call(String OS, String ARCH) {
    def versionFilePath = "version_output.txt"
    def versionOutput = readFile(versionFilePath)

    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()
    def baseVersion = versionOutput.split('\n').find { it.startsWith('CURRENT_BASE_VERSION=') }?.split('=')[1]?.trim()
    def rawBaseVersion = baseVersion?.startsWith('v') ? baseVersion.substring(1) : baseVersion
    def appName = "pars"

    def extLine = versionOutput.split('\n').find { it.startsWith('EXT=') }
    def ext = ""
    if (extLine?.contains('=') && extLine.split('=').length > 1) {
        ext = extLine.split('=')[1].trim()
    } else if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    // unstash 'linux-x86_64-artifacts'
    def binaryOutputBase = "dist/${buildVersion}/${OS}/bin/${ARCH}"
    def originalFileName = "${appName}${ext}"
    def newBaseName = "${appName}-${rawBaseVersion}.tar.gz"

    def packageOutputBase = "dist/${buildVersion}/${OS}/pkg/rpm/${ARCH}/${appName}"
    def packageSourceDir = "${packageOutputBase}/SOURCES"
    def tarPath = "${binaryOutputBase}/${newBaseName}"

    // Create tar.gz
    sh """
        mkdir -p '${packageSourceDir}'
        cd '${binaryOutputBase}'
        tar -czf '${newBaseName}' . --warning=no-file-changed || true
        if ! tar -tzf '${newBaseName}' > /dev/null; then
        echo "Error with archive: Unable to list contents for validation."
        fi
        cp '${newBaseName}' '${WORKSPACE}/${packageSourceDir}/'
    """

    // Build RPM package
    sh """
        export GO111MODULE=on
        make build.rpm.package.${ARCH}.configuration VERSION=${buildVersion}
        make build.rpm.package.${ARCH}.package VERSION=${buildVersion}
    """

    // Map ARCH to RPM arch
    def rpmArch = ""
    switch (ARCH) {
        case "x86":
            rpmArch = "i386"
            break
        case "x86_64":
            rpmArch = "x86_64"
            break
        case "arm":
            rpmArch = "armv7hl"
            break
        case "arm64":
            rpmArch = "aarch64"
            break
        default:
            error "Unsupported architecture: ${ARCH}"
    }

    def rpmOutputBase = "${packageOutputBase}/RPMS/${rpmArch}"
    def checksumPath = "${rpmOutputBase}/checksum.txt"

    // Find .rpm file and calculate checksum
    def rpmFile = sh(script: "ls ${rpmOutputBase}/${appName}*.${rpmArch}.rpm | head -n1", returnStdout: true).trim()
    def checksum = sh(script: "sha256sum '${rpmFile}' | awk '{print \$1}'", returnStdout: true).trim()

    echo "Checksum for ${rpmFile}: ${checksum}"
    writeFile file: checksumPath, text: checksum

}

def copyRpmAndUpdateChecksums(String OS, String ARCH) {
    def versionFilePath = "version_output.txt"
    def versionOutput = readFile(versionFilePath)

    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()
    def baseVersion = versionOutput.split('\n').find { it.startsWith('CURRENT_BASE_VERSION=') }?.split('=')[1]?.trim()
    def rawBaseVersion = baseVersion?.startsWith('v') ? baseVersion.substring(1) : baseVersion
    def appName = "pars"
    def artifactPath = "dist/artifacts/${buildVersion}"

    def extLine = versionOutput.split('\n').find { it.startsWith('EXT=') }
    def ext = ""
    if (extLine?.contains('=') && extLine.split('=').length > 1) {
        ext = extLine.split('=')[1].trim()
    } else if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    def rpmArch = ""

    switch (ARCH) {
        case "x86":
            rpmArch = "i386"
            break
        case "x86_64":
            rpmArch = "x86_64"
            break
        case "arm":
            rpmArch = "armv7hl"
            break
        case "arm64":
            rpmArch = "aarch64"
            break
        default:
            error "Unsupported architecture: ${ARCH}"
    }

    def newBaseName = "${appName}-${OS}-${ARCH}.rpm"
    def packageOutputBase = "dist/${buildVersion}/${OS}/pkg/rpm/${ARCH}/${appName}"
    def rpmOutputBase = "${packageOutputBase}/RPMS/${rpmArch}"
    def checksumFilePath = "${rpmOutputBase}/checksum.txt"
    def checksumsMdPath = "${artifactPath}/Checksums.md"

    // Find RPM file
    def rpmFiles = sh(script: "ls ${rpmOutputBase}/${appName}*.${rpmArch}.rpm", returnStdout: true).trim().split("\n")

    if (rpmFiles.size() == 0) {
        error "No RPM file found for ${appName} with arch ${rpmArch} in ${rpmOutputBase}"
    }

    def rpmOutputPath = rpmFiles[0]

    // Copy RPM file to artifact path with new name
    sh "cp '${rpmOutputPath}' '${artifactPath}/${newBaseName}'"

    // Read checksum
    def checksum = readFile(checksumFilePath).trim()

    // Append to Checksums.md
    def checksumLine = "| ${OS} | ${ARCH} | RPM | ${newBaseName} | ${checksum} |"
    writeFile file: checksumsMdPath, text: "${checksumLine}\n", encoding: "UTF-8", append: true

    echo "Added checksum line to Checksums.md: ${checksumLine}"


    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

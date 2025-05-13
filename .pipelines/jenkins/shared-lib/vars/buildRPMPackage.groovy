def call(String OS, String ARCH) {
    def ext = ""
    if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    // unstash 'linux-x86_64-artifacts'
    def binaryOutputBase = "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}"
    def originalFileName = "${APPNAME}${ext}"
    def newBaseName = "${APPNAME}-${env.CURRENT_BASE_VERSION_RAW}.tar.gz"

    def packageOutputBase = "dist/${env.BUILD_VERSION}/${OS}/pkg/rpm/${ARCH}/${APPNAME}"
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
        make build.rpm.package.${ARCH}.configuration VERSION=${env.BUILD_VERSION}
        make build.rpm.package.${ARCH}.package VERSION=${env.BUILD_VERSION}
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
    def rpmFile = sh(script: "ls ${rpmOutputBase}/${APPNAME}*.${rpmArch}.rpm | head -n1", returnStdout: true).trim()
    def checksum = sh(script: "sha256sum '${rpmFile}' | awk '{print \$1}'", returnStdout: true).trim()

    echo "Checksum for ${rpmFile}: ${checksum}"
    writeFile file: checksumPath, text: checksum

}


return this

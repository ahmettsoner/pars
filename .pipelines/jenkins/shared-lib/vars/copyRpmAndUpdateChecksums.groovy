def call(String OS, String ARCH) {
    def artifactPath = "dist/artifacts/${env.BUILD_VERSION}"

    def ext = ""
    if (OS.toLowerCase() == 'windows') {
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

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def packageOutputBase = "dist/${env.BUILD_VERSION}/${OS}/pkg/rpm/${ARCH}/${APPNAME}"
    def rpmOutputBase = "${packageOutputBase}/RPMS/${rpmArch}"
    def checksumFilePath = "${rpmOutputBase}/checksum.txt"

    // Find RPM file
    def rpmFiles = sh(script: "ls ${rpmOutputBase}/${APPNAME}*.${rpmArch}.rpm", returnStdout: true).trim().split("\n")

    if (rpmFiles.size() == 0) {
        error "No RPM file found for ${APPNAME} with arch ${rpmArch} in ${rpmOutputBase}"
    }

    def rpmOutputPath = rpmFiles[0]

    // Copy RPM file to artifact path with new name
    sh "cp '${rpmOutputPath}' '${env.ARTIFACT_PATH}/${newBaseName}'"

    // // Read checksum
    // def checksum = readFile(checksumFilePath).trim()

    // // Append to Checksums.md
    // def type = "RPM"
    // def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"
    // sh "echo '${line}' >> ${env.ARTIFACT_PATH}/${newBaseName}-checksum.txt"

    // echo "Added checksum line to Checksums.md: ${line}"


    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

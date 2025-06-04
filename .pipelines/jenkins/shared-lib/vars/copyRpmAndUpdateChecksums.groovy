def call(String OS, String ARCH) {
    def artifactPath = "dist/artifacts/${env.CURRENT_VERSION}"

    def utils = new com.parsdevkit.Utils(this)
    def platformArch = utils.mapArch("rhel", ARCH)

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def packageOutputBase = "dist/${env.CURRENT_VERSION}/${OS}/pkg/rpm/${ARCH}/${APPNAME}"
    def rpmOutputBase = "${packageOutputBase}/RPMS/${platformArch}"
    def checksumFilePath = "${rpmOutputBase}/checksum.txt"

    // Find RPM file
    def rpmFiles = sh(script: "ls ${rpmOutputBase}/${APPNAME}*.${platformArch}.rpm", returnStdout: true).trim().split("\n")

    if (rpmFiles.size() == 0) {
        error "No RPM file found for ${APPNAME} with arch ${platformArch} in ${rpmOutputBase}"
    }

    def rpmOutputPath = rpmFiles[0]

    // Copy RPM file to artifact path with new name
    sh "cp '${rpmOutputPath}' '${env.ARTIFACT_PATH}/${newBaseName}'"

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

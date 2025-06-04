def call(String OS, String ARCH) {
    def utils = new com.parsdevkit.Utils(this)
    def ext = utils.appExt(OS)

    unstash "${OS}-${ARCH}-dist-bin"
    def binaryOutputBase = "dist/${env.CURRENT_VERSION}/${OS}/bin/${ARCH}"
    def originalFileName = "${APPNAME}${ext}"
    def newBaseName = "${APPNAME}-${env.CURRENT_BASE_VERSION_RAW}.tar.gz"

    def packageOutputBase = "dist/${env.CURRENT_VERSION}/${OS}/pkg/rpm/${ARCH}/${APPNAME}"
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
        make build.rpm.package.${ARCH}.configuration VERSION=${env.CURRENT_VERSION}
        make build.rpm.package.${ARCH}.package VERSION=${env.CURRENT_VERSION}
    """
}


return this

def call(String OS, String ARCH) {

    unstash "${OS}-${ARCH}-dist-bin"
    def binaryOutputPathBase = "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}"

    def utils = new com.parsdevkit.Utils(this)
    def ext = utils.appExt(OS)
    def binaryOutputPath = "${binaryOutputPathBase}/${APPNAME}${ext}"

    def pckgPath = "usr/bin"
    def packageOutputBase = "dist/${env.BUILD_VERSION}/${OS}/pkg/deb/${ARCH}"
    def packageBinaryPath = "${packageOutputBase}/${APPNAME}/${pckgPath}"

    sh """
        mkdir -p '${packageBinaryPath}'
        cp '${binaryOutputPath}' '${packageBinaryPath}'
        export GO111MODULE=on
        make build.deb.package.${ARCH}.configuration VERSION=${env.BUILD_VERSION}
        make build.deb.package.${ARCH}.package VERSION=${env.BUILD_VERSION}
    """

}

return this
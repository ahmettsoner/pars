def call(String OS, String ARCH) {
    def ext = OS.toLowerCase() == 'windows' ? '.exe' : ''

    // Make komutunu çalıştırarak binary dosyasını oluşturuyoruz
    sh """
        make build.binary.${OS}.${ARCH} VERSION=$env.BUILD_VERSION
    """
    stash includes: "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}/**/*", name: "${OS}-${ARCH}-dist-bin"

}
return this

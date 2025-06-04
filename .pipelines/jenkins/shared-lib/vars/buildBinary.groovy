def call(String OS, String ARCH) {
    // Make komutunu çalıştırarak binary dosyasını oluşturuyoruz
    sh """
        make build.binary.${OS}.${ARCH} VERSION=$env.CURRENT_VERSION
    """
    stash includes: "dist/${env.CURRENT_VERSION}/${OS}/bin/${ARCH}/**/*", name: "${OS}-${ARCH}-dist-bin"

}
return this

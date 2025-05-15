def call(String OS, String ARCH) {
    def ext = ""
    if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

    // Make komutunu çalıştırarak binary dosyasını oluşturuyoruz
    sh """
        make build.binary.${OS}.${ARCH} VERSION=$env.BUILD_VERSION
    """

    // // Binary dosyasının çıkış yolunu belirliyoruz
    // def binaryOutputPathBase = "dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}"
    // def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"
    // def binaryChecksumPath = "${binaryOutputPathBase}/checksum.txt"

    // // Binary dosyasının checksum'unu hesaplıyoruz
    // def binaryChecksum = sh(script: "sha256sum ${binaryOutputPath} | awk '{print \$1}'", returnStdout: true).trim()

    // // Checksum değerini ekrana yazdırıyoruz
    // echo "Checksum for ${binaryOutputPath}: ${binaryChecksum}"

    // // Checksum ve dosya yolunu checksums.txt dosyasına yazıyoruz
    // sh """
    //     echo "${binaryChecksum}" > ${binaryChecksumPath}
    // """
}
return this

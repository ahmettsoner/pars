def build() {
    def versionOutput = readFile("${WORKSPACE}/version_output.txt")
    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()
    def extLine = versionOutput.split('\n').find { it.startsWith('EXT=') }
    def ext = (extLine?.contains('=') && extLine.split('=').length > 1) ? extLine.split('=')[1].trim() : ""

    // Make komutunu çalıştırarak binary dosyasını oluşturuyoruz
    sh """
        make build.binary.${OS}.${ARCH} VERSION=$buildVersion
    """

    // Binary dosyasının çıkış yolunu belirliyoruz
    def binaryOutputPathBase = "dist/${buildVersion}/${OS}/bin/${ARCH}"
    def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"
    def binaryChecksumPath = "${binaryOutputPathBase}/checksum.txt"

    // Binary dosyasının checksum'unu hesaplıyoruz
    def binaryChecksum = sh(script: "sha256sum ${binaryOutputPath} | awk '{print \$1}'", returnStdout: true).trim()

    // Checksum değerini ekrana yazdırıyoruz
    echo "Checksum for ${binaryOutputPath}: ${binaryChecksum}"

    // Checksum ve dosya yolunu checksums.txt dosyasına yazıyoruz
    sh """
        echo "${binaryChecksum}" > ${binaryChecksumPath}
    """
    

    // // Binary dosyasının yolunu environment değişkeni olarak ekliyoruz
    // env.BINARY_OUTPUT_PATH_BASE = binaryOutputPathBase
    // env.BINARY_OUTPUT_PATH = binaryOutputPath
    // env.BINARY_CHECKSUM_PATH = binaryChecksumPath

    // // Environment değişkenlerini GitHub Actions ortamına aktarıyoruz (isteğe bağlı)
    // echo "BINARY_OUTPUT_PATH_BASE=\${binaryOutputPathBase}" >> $GITHUB_ENV
    // echo "BINARY_OUTPUT_PATH=\${binaryOutputPath}" >> $GITHUB_ENV
    // echo "BINARY_CHECKSUM_PATH=\${binaryChecksumPath}" >> $GITHUB_ENV
}
return this
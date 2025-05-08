def build(String OS, String ARCH, String versionFilePath) {
    def versionOutput = readFile(versionFilePath)
    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()
    def extLine = versionOutput.split('\n').find { it.startsWith('EXT=') }
    def ext = ""
    
    if (extLine?.contains('=') && extLine.split('=').length > 1) {
        ext = extLine.split('=')[1].trim()
    } else if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }

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
    
    // Binary dosyasının yolunu environment değişkeni olarak ekliyoruz
    // env.BINARY_OUTPUT_PATH_BASE = binaryOutputPathBase
    // env.BINARY_OUTPUT_PATH = binaryOutputPath
    // env.BINARY_CHECKSUM_PATH = binaryChecksumPath
}
return this

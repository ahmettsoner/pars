def call(String OS, String ARCH) {
    def utils = new com.parsdevkit.Utils(this)
    def platformArch = utils.mapArch("windows", ARCH)

    def plainVersion = env.CURRENT_VERSION.replaceFirst(/^v/, "")
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.msi"
    def packageOutputBase = "dist\\${env.CURRENT_VERSION}\\${OS}\\pkg\\msi\\${ARCH}"
    def msiOutputBase = "${packageOutputBase}\\output"
    def msiOutputPath = "${msiOutputBase}\\${APPNAME}_${plainVersion}_${platformArch}.msi"
    def checksumFilePath = "${msiOutputBase}\\checksum.txt"

    // Dosya var mı kontrolü
    if (!fileExists(msiOutputPath)) {
        error "No MSI file found at ${msiOutputPath}"
    }

    // Artifact dizinine kopyala (PowerShell)
    powershell """
        Copy-Item -Path '${msiOutputPath}' -Destination 'dist\\artifacts\\${env.CURRENT_VERSION}\\${newBaseName}' -Force
    """

    // Stash işlemi
    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-msi-package-artifacts"
}

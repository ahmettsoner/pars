def call(String OS, String ARCH) {
    def utils = new com.parsdevkit.Utils(this)
    def ext = utils.appExt(OS)

    unstash "${OS}-${ARCH}-dist-bin"
    def binaryOutputBase = "dist\\${env.CURRENT_VERSION}\\${OS}\\bin\\${ARCH}"
    def originalFileName = "${APPNAME}${ext}"
    def newBaseName = "${APPNAME}-${env.CURRENT_BASE_VERSION_RAW}.zip"

    def packageOutputBase = "dist\\${env.CURRENT_VERSION}\\${OS}\\pkg\\msi\\${ARCH}\\${APPNAME}"
    def packageSourceDir = "${packageOutputBase}\\SOURCES"

    powershell """
            Write-Host ">>>>>: ${packageSourceDir}'"
    """

    // Create zip (Windows PowerShell)
    powershell """
        \$ErrorActionPreference = 'Stop'
        New-Item -ItemType Directory -Force -Path '${packageSourceDir}' | Out-Null
        Push-Location '${binaryOutputBase}'
        try {
            Compress-Archive -Path * -DestinationPath '${newBaseName}' -Force
        } catch {
            Write-Host "Compress-Archive failed: $($_.Exception.Message)"
        }
        if (!(Test-Path '${newBaseName}')) {
            Write-Host "Error with archive: Unable to validate archive contents."
        }
        Copy-Item -Path '${newBaseName}' -Destination '${WORKSPACE}\\${packageSourceDir}' -Force
        Pop-Location
    """

    // Build MSI package
    powershell """
        \$env:GO111MODULE = "on"
        make build.msi.package.${ARCH}.configuration VERSION=${env.CURRENT_VERSION}
        make build.msi.package.${ARCH}.package VERSION=${env.CURRENT_VERSION}
    """
}

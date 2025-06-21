def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-msi-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.msi"
    def msiOutputPath = "dist\\artifacts\\${env.CURRENT_VERSION}\\${newBaseName}"

    withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
        powershell """
            \$ProgressPreference = 'SilentlyContinue'
            Invoke-RestMethod -Uri "${NEXUS_URL}/repository/apt-${CHANNEL}/" `
                              -Method Post `
                              -Headers @{"Content-Type" = "multipart/form-data"} `
                              -InFile "${msiOutputPath}" `
                              -Credential (New-Object System.Management.Automation.PSCredential('${NEXUS_USER}', (ConvertTo-SecureString '${NEXUS_PASS}' -AsPlainText -Force)))
        """
    }

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-msi-package-artifacts"
}

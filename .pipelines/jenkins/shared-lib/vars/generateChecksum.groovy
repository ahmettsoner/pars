def generateChecksum(String label) {
    node(label) {
        stage("Create ${label.toUpperCase()} markdown file for checksums") {
            steps {
                script {
                    def versionOutput = readFile("${WORKSPACE}/version_output.txt")
                    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()

                    echo "Build Version: ${buildVersion}"

                    if (!buildVersion) {
                        error "BUILD_VERSION is empty or null"
                    }

                    def ARTIFACT_PATH = "dist/artifacts/${buildVersion}"

                    sh """
                        echo "Creating directory: ${ARTIFACT_PATH}"
                        mkdir -p "${ARTIFACT_PATH}"

                        echo "| Platform | Architecture | Type      | File Name | SHA-256 Checksum |" > ${ARTIFACT_PATH}/Checksums.md
                        echo "|----------|--------------|-----------|-----------|------------------|" >> ${ARTIFACT_PATH}/Checksums.md
                    """
                }
            }
        }
    }
}

return this
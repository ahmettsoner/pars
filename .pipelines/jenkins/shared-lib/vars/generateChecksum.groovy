def generateChecksum() {
    echo "Build Version: ${env.BUILD_VERSION}"

    if (!env.BUILD_VERSION) {
        error "BUILD_VERSION is empty or null"
    }

    def ARTIFACT_PATH = "dist/artifacts/${env.BUILD_VERSION}"

    sh """
        echo "Creating directory: ${ARTIFACT_PATH}"
        mkdir -p "${ARTIFACT_PATH}"

        echo "| Platform | Architecture | Type      | File Name | SHA-256 Checksum |" > ${ARTIFACT_PATH}/Checksums.md
        echo "|----------|--------------|-----------|-----------|------------------|" >> ${ARTIFACT_PATH}/Checksums.md
    """
}

return this
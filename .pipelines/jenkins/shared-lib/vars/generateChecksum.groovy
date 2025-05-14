def call() {
    echo "Build Version: ${env.BUILD_VERSION}"

    env.ARTIFACT_CHECKSUM_MD5_PATH = "${env.ARTIFACT_PATH}/Checksums.md"
    if (!env.BUILD_VERSION) {
        error "BUILD_VERSION is empty or null"
    }

    sh """
        echo "| Platform | Architecture | Type      | File Name | SHA-256 Checksum |" > ${env.ARTIFACT_CHECKSUM_MD5_PATH}
        echo "|----------|--------------|-----------|-----------|------------------|" >> ${env.ARTIFACT_CHECKSUM_MD5_PATH}
    """
}

return this
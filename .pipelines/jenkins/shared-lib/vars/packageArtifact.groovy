def packageArtifact(String OS, String ARCH, String versionFilePath, String archiveFormat) {
    def versionOutput = readFile(versionFilePath)
    def appName = versionOutput.split('\n').find { it.startsWith('APPNAME=') }?.split('=')[1]?.trim()
    def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()
    def html_output_dir = versionOutput.split('\n').find { it.startsWith('HTML_OUTPUT_DIR=') }?.split('=')[1]?.trim()
    def extLine = versionOutput.split('\n').find { it.startsWith('EXT=') }
    def ext = ""
    
    if (extLine?.contains('=') && extLine.split('=').length > 1) {
        ext = extLine.split('=')[1].trim()
    } else if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }
    
    def binaryOutputPathBase = "${WORKSPACE}/dist/${buildVersion}/${OS}/bin/${ARCH}"
    def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"
    def binaryChecksumPath = "${binaryOutputPathBase}/checksum.txt"
    def artifactPath = "dist/artifacts/${buildVersion}"
    
    if (!fileExists(binaryOutputPathBase) || !fileExists(binaryChecksumPath)) {
        echo "Required files do not exist. Skipping archive step."
        return
    }

    echo "Binary and Checksum files are present."

    def newBaseName = "${appName}-${OS}-${ARCH}.bin.${archiveFormat}"
    def binaryTempPath = "${WORKSPACE}/dist/temp/${buildVersion}/${appName}-${OS}-${ARCH}-${archiveFormat}"

    sh """
        mkdir -p ${binaryTempPath}/bin
        cp -r ${binaryOutputPathBase} ${binaryTempPath}/bin/

        mkdir -p ${binaryTempPath}/meta
        cp -r ${binaryChecksumPath} ${binaryTempPath}/meta/

        mkdir -p ${binaryTempPath}/docs
        cp -r ${html_output_dir}/ ${binaryTempPath}/docs/
    """

    if (archiveFormat == "zip") {
        sh "cd ${binaryTempPath} && zip -r '${newBaseName}' ."
    } else if (archiveFormat == "tar.gz") {
        sh """
            cd ${binaryTempPath} && tar -czf '${newBaseName}' . --warning=no-file-changed || true
            if ! tar -tzf '${binaryTempPath}/${newBaseName}' > /dev/null; then
                echo "Error: Cannot list contents of archive."
            fi
        """
    } else if (archiveFormat == "7z") {
        sh "cd ${binaryTempPath} && 7z a '${newBaseName}' *"
    } else if (archiveFormat == "rar") {
        sh "cd ${binaryTempPath} && rar a '${newBaseName}' *"
    } else {
        error "Unsupported archive format: ${archiveFormat}"
    }

    sh "mkdir -p ${artifactPath}"
    sh "cp '${binaryTempPath}/${newBaseName}' '${artifactPath}/${newBaseName}'"

    def archiveChecksum = sh(script: "sha256sum '${binaryTempPath}/${newBaseName}' | awk '{print \$1}'", returnStdout: true).trim()
    echo "Checksum for ${binaryTempPath}/${newBaseName}: ${archiveChecksum}"

    def type = "Archive"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${archiveChecksum} |"
    sh "echo '${line}' >> ${artifactPath}/Checksums.md"

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-archive-artifacts"
}

return this
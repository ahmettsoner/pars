def call(String OS, String ARCH, String archiveFormat) {
    def ext = ""
    if (OS.toLowerCase() == 'windows') {
        ext = ".exe"
    }
    
    def binaryOutputPathBase = "${WORKSPACE}/dist/${env.BUILD_VERSION}/${OS}/bin/${ARCH}"
    def binaryOutputPath = "${binaryOutputPathBase}/pars${ext}"
    def binaryChecksumPath = "${binaryOutputPathBase}/checksum.txt"
    
    if (!fileExists(binaryOutputPathBase) || !fileExists(binaryChecksumPath)) {
        echo "Required files do not exist. Skipping archive step."
        return
    }

    echo "Binary and Checksum files are present."

    def newBaseName = "${APPNAME}-${OS}-${ARCH}.bin.${archiveFormat}"
    def binaryTempPath = "${WORKSPACE}/dist/temp/${env.BUILD_VERSION}/${APPNAME}-${OS}-${ARCH}-${archiveFormat}"


    unstash "html-outputdir"
    sh """
        mkdir -p ${binaryTempPath}/bin
        cp -r ${binaryOutputPathBase} ${binaryTempPath}/bin/

        mkdir -p ${binaryTempPath}/meta
        cp -r ${binaryChecksumPath} ${binaryTempPath}/meta/

        mkdir -p ${binaryTempPath}/docs
        cp -r ${env.HTML_OUTPUT_DIR}/ ${binaryTempPath}/docs/
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

    sh "cp '${binaryTempPath}/${newBaseName}' '${env.ARTIFACT_PATH}/${newBaseName}'"

    def checksum = sh(script: "sha256sum '${binaryTempPath}/${newBaseName}' | awk '{print \$1}'", returnStdout: true).trim()
    echo "Checksum for ${binaryTempPath}/${newBaseName}: ${checksum}"

    def type = "Archive"
    def line = "| ${OS} | ${ARCH} | ${type} | ${newBaseName} | ${checksum} |"
    sh "echo '${line}' >> ${env.ARTIFACT_PATH}/${newBaseName}-checksum.txt"

    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-archive-artifacts"
}

return this
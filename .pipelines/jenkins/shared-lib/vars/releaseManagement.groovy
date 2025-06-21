def releaseRepo(List<String> osList, List<String> archList) {
    unstash "artifacts-checksums"
    def checksumTest = readFile(env.ARTIFACT_CHECKSUM_MD5_PATH).trim()

    sh(
        script: "grm changelog generate --from $env.CURRENT_BASE_VERSION-dev.1 --to $env.CURRENT_VERSION --environment $env.CHANNEL --merge-all --output ${env.CHANGELOG_PATH}",
        returnStdout: true
    ).trim()

    def changelogText = readFile(env.CHANGELOG_PATH).trim()

    def changelog = "${changelogText}\n\n---\n\n${checksumTest}"

    // downloads/pars/dev/v1.4.0-dev.3/windows/x86_64/pars.exe
    // downloads/pars/dev/latest/windows/x86_64/pars.exe
    // downloads/pars/stable/v1.3.0/linux/amd64/pars
    // downloads/pars/test/v1.4.0-rc1/darwin/arm64/pars

    withCredentials([
        string(credentialsId: 'GITEA_TOKEN', variable: 'GITEA_TOKEN'),
        usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')
    ]) {


        def remoteChangelogFilePath = "${NEXUS_URL}/repository/release-dev/${APPNAME}/${CHANNEL}/${env.CURRENT_VERSION}/changelog.md"
        writeFile file: "${env.ARTIFACT_PATH}/changelog.md", text: changelog
        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${env.ARTIFACT_PATH}/changelog.md\" \"${remoteChangelogFilePath}\""


        def remoteChecksumFilePath = "${NEXUS_URL}/repository/release-dev/${APPNAME}/${CHANNEL}/${env.CURRENT_VERSION}/checksums.md"
        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${env.ARTIFACT_PATH}/Checksums.md\" \"${remoteChecksumFilePath}\""

        osList.each { OS ->
            archList.each { ARCH ->
                if (!(OS == 'netbsd' && ARCH == 'arm64')) {
                    unstash "${OS}-${ARCH}-artifacts"
                    unstash "${OS}-${ARCH}-archive-artifacts"

                    def utils = new com.parsdevkit.Utils(this)
                    def platformArch = utils.mapArch("rhel", ARCH)
                    def ext = utils.appExt(OS)
                    def archiveFormat = utils.archiveFormat(OS)

                    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"
                    def remoteFilePath = "${NEXUS_URL}/repository/binary-dev/${APPNAME}/${CHANNEL}/${env.CURRENT_VERSION}/${OS}/${platformArch}/${APPNAME}${ext}"

                    sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${env.ARTIFACT_PATH}/${newBaseName}\" \"${remoteFilePath}\""


                    def newArchiveBaseName = "${APPNAME}-${OS}-${ARCH}.bin.${archiveFormat}"
                    def remoteArchiveFilePath = "${NEXUS_URL}/repository/binary-dev/${APPNAME}/${CHANNEL}/${env.CURRENT_VERSION}/${OS}/${platformArch}/${APPNAME}.${archiveFormat}"

                    sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${env.ARTIFACT_PATH}/${newArchiveBaseName}\" \"${remoteArchiveFilePath}\""


                    if (OS == 'linux') {
                        unstash "${OS}-${ARCH}-rpm-package-artifacts"
                        unstash "${OS}-${ARCH}-deb-package-artifacts"


                        def newRPMBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
                        def remoteRPMFilePath = "${NEXUS_URL}/repository/yum-dev/${APPNAME}/${CHANNEL}/${env.CURRENT_VERSION}/${OS}/${platformArch}/${APPNAME}.rpm"

                        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${env.ARTIFACT_PATH}/${newRPMBaseName}\" \"${remoteRPMFilePath}\""

                        publishRpmPackage(OS, ARCH)

                        def newDebBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
                        def remoteDebFilePath = "${NEXUS_URL}/repository/apt-dev/${APPNAME}/${CHANNEL}/${env.CURRENT_VERSION}/${OS}/${platformArch}/${APPNAME}.deb"

                        sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${env.ARTIFACT_PATH}/${newDebBaseName}\" \"${remoteDebFilePath}\""

                        
                        publishDebPackage(OS, ARCH, "universial")
                    } else if (OS == 'windows') {
                        unstash "${OS}-${ARCH}-msi-package-artifacts"

                        def newMsiBaseName = "${APPNAME}-${OS}-${ARCH}.msi"
                        def remoteMsiFilePath = "${NEXUS_URL}/repository/msi-dev/${APPNAME}/${CHANNEL}/${env.CURRENT_VERSION}/${OS}/${platformArch}/${APPNAME}.msi"

                        withCredentials([usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
                            sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${env.ARTIFACT_PATH}/${newMsiBaseName}\" \"${remoteMsiFilePath}\""
                        }

                        publishMsiInstaller(OS, ARCH)
                    }
                }
            }
        }
    }
}
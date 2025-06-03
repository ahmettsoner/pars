def release(List<String> osList, List<String> archList) {

    osList.each { OS ->
        archList.each { ARCH ->
            if (!(OS == 'netbsd' && ARCH == 'arm64')) {
                unstash "${OS}-${ARCH}-artifacts"
                unstash "${OS}-${ARCH}-archive-artifacts"
                if (OS == 'linux') {
                    unstash "${OS}-${ARCH}-rpm-package-artifacts"
                    unstash "${OS}-${ARCH}-deb-package-artifacts"
                }
            }
        }
    }
    unstash "artifacts-checksums"

    sh(
        script: "grm changelog generate --from $env.CURRENT_BASE_VERSION-dev.1 --to $env.BUILD_VERSION --environment $env.CHANNEL --merge-all --output ${env.CHANGELOG_PATH}",
        returnStdout: true
    ).trim()

    def changelogText = readFile(env.CHANGELOG_PATH).trim()
    def checksumTest = readFile(env.ARTIFACT_CHECKSUM_MD5_PATH).trim()

    def changelog = "${changelogText}\n\n---\n\n${checksumTest}"

    withCredentials([string(credentialsId: 'GITEA_TOKEN', variable: 'GITEA_TOKEN')]) {

        def releaseJson = """{
            "tag_name": "${env.BUILD_VERSION}",
            "target": "dev",
            "name": "${env.BUILD_VERSION} Release",
            "body": ${groovy.json.JsonOutput.toJson(changelog)},
            "draft": false,
            "prerelease": false
        }"""

        writeFile file: 'release.json', text: releaseJson

        sh """
            curl -X POST "$GITEA_URL/api/v1/repos/$GITEA_OWNER/$GITEA_REPO/releases" \\
                -H "Content-Type: application/json" \\
                -H "Authorization: token $GITEA_TOKEN" \\
                -d @release.json
        """

        def releaseInfo = sh(
            script: """curl -s -H "Authorization: token $GITEA_TOKEN" \\
                "$GITEA_URL/api/v1/repos/$GITEA_OWNER/$GITEA_REPO/releases/tags/${env.BUILD_VERSION}" """,
            returnStdout: true
        ).trim()

        def releaseId = new groovy.json.JsonSlurper().parseText(releaseInfo).id

        def fileList = sh(
            script: "find ${env.ARTIFACT_PATH} -type f",
            returnStdout: true
        ).trim().split('\n')
        fileList.each { filePath ->
            def fileName = filePath.tokenize('/').last()
            echo "Uploading artifact: ${fileName}"
            sh """
                curl -X POST "$GITEA_URL/api/v1/repos/$GITEA_OWNER/$GITEA_REPO/releases/$releaseId/assets?name=${fileName}" \\
                    -H "Authorization: token $GITEA_TOKEN" \\
                    -H "Content-Type: application/octet-stream" \\
                    --data-binary @${filePath}
            """
        }
    }
}

def releaseRepo() {
    unstash "artifacts-checksums"
    def checksumTest = readFile(env.ARTIFACT_CHECKSUM_MD5_PATH).trim()

    sh(
        script: "grm changelog generate --from $env.CURRENT_BASE_VERSION-dev.1 --to $env.BUILD_VERSION --environment $env.CHANNEL --merge-all --output ${env.CHANGELOG_PATH}",
        returnStdout: true
    ).trim()

    def changelogText = readFile(env.CHANGELOG_PATH).trim()

    def changelog = "${changelogText}\n\n---\n\n${checksumTest}"

    def osList = ['linux']//, 'windows', 'darwin', 'openbsd', 'netbsd', 'freebsd']
    def archList = ['x86_64']//, 'arm64']

    // downloads/pars/dev/v1.4.0-dev.3/windows/x86_64/pars.exe
    // downloads/pars/dev/latest/windows/x86_64/pars.exe
    // downloads/pars/stable/v1.3.0/linux/amd64/pars
    // downloads/pars/test/v1.4.0-rc1/darwin/arm64/pars

    withCredentials([
        string(credentialsId: 'GITEA_TOKEN', variable: 'GITEA_TOKEN'),
        usernamePassword(credentialsId: 'nexus-credential', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')
        ]) {

        // def releaseJson = """{
        //     "tag_name": "${env.BUILD_VERSION}",
        //     "target": "dev",
        //     "name": "${env.BUILD_VERSION} Release",
        //     "body": ${groovy.json.JsonOutput.toJson(changelog)},
        //     "draft": false,
        //     "prerelease": false
        // }"""

        // writeFile file: 'release.json', text: releaseJson

        // sh """
        //     curl -X POST "$GITEA_URL/api/v1/repos/$GITEA_OWNER/$GITEA_REPO/releases" \\
        //         -H "Content-Type: application/json" \\
        //         -H "Authorization: token $GITEA_TOKEN" \\
        //         -d @release.json
        // """

        // def releaseInfo = sh(
        //     script: """curl -s -H "Authorization: token $GITEA_TOKEN" \\
        //         "$GITEA_URL/api/v1/repos/$GITEA_OWNER/$GITEA_REPO/releases/tags/${env.BUILD_VERSION}" """,
        //     returnStdout: true
        // ).trim()

        // def releaseId = new groovy.json.JsonSlurper().parseText(releaseInfo).id

        def String[] fileList = new String[0] 
        osList.each { OS ->
            archList.each { ARCH ->
                if (!(OS == 'netbsd' && ARCH == 'arm64')) {
                    unstash "${OS}-${ARCH}-artifacts"
                    unstash "${OS}-${ARCH}-archive-artifacts"



                    def ext = OS.toLowerCase() == 'windows' ? '.exe' : ''
                    def originalFileName = "${APPNAME}${ext}"
                    def newBaseName = "${APPNAME}-${OS}-${ARCH}${ext}"
                    def remoteFilePath = "${NEXUS_URL}/repository/raw/downloads/${APPNAME}/${CHANNEL}/${env.BUILD_VERSION}/${OS}/${ARCH}/${APPNAME}${ext}"

                    sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file \"${env.ARTIFACT_PATH}/${newBaseName}\" \"${remoteFilePath}\""


                    if (OS == 'linux') {
                        unstash "${OS}-${ARCH}-rpm-package-artifacts"
                        unstash "${OS}-${ARCH}-deb-package-artifacts"
                    }
                }
            }
        }

        // def fileList = sh(
        //     script: "find ${env.ARTIFACT_PATH} -type f",
        //     returnStdout: true
        // ).trim().split('\n')
        // fileList.each { filePath ->
        //     def fileName = filePath.tokenize('/').last()
        //     echo "Uploading artifact: ${fileName}"
        //     sh "curl -v -u \"${NEXUS_USER}:${NEXUS_PASS}\" --upload-file ${filePath} \"${NEXUS_URL}/repository/raw/\""
        // }
    }
}
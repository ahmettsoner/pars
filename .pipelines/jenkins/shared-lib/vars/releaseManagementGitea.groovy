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
def call(String OS, String ARCH) {
    unstash "${OS}-${ARCH}-rpm-package-artifacts"
    def newBaseName = "${APPNAME}-${OS}-${ARCH}.rpm"
    def rpmOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"


    sh "curl -v -u \"$NEXUS_USER:$NEXUS_PASS\" --upload-file ~/rpmbuild/RPMS/x86_64/my-package.rpm \"$NEXUS_URL/my-package.rpm\""


    stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-rpm-package-artifacts"
}

return this

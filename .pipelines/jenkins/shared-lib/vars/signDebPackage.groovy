def call(String OS, String ARCH, List<String> DIST_CODENAMES){
    withCredentials([
        file(credentialsId: 'public-rpm.gpg', variable: 'GPG_PUBLIC'),
        file(credentialsId: 'private-rpm.gpg', variable: 'GPG_PRIVATE'),
        string(credentialsId: 'GPG_PASSPHRASE', variable: 'GPG_PASSPHRASE'),
        string(credentialsId: 'GPG_FINGERPRINT', variable: 'GPG_FINGERPRINT')
    ]) {

        unstash "${OS}-${ARCH}-deb-package-artifacts"

        def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
        def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"
        def aptlyRepoName = "${APPNAME}-${OS}-${ARCH}"
        def debRepoOutputPath = "${env.ARTIFACT_PATH}/${aptlyRepoName}"

        def publishDir = "${debRepoOutputPath}/publish"

        sh '''
            set -e

            # Setup GPG
            mkdir -p ~/.gnupg
            chmod 700 ~/.gnupg

            gpg --batch --import "$GPG_PUBLIC"
            gpg --batch --import "$GPG_PRIVATE"

            echo "$GPG_FINGERPRINT:6:" > trust.txt
            gpg --import-ownertrust trust.txt

            echo "use-agent" > ~/.gnupg/gpg.conf
            echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
            echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf
            gpgconf --kill gpg-agent || true
            gpgconf --launch gpg-agent || true
        '''

        sh "ar t ${debOutputPath}"

        sh """
            aptly repo drop -force ${aptlyRepoName} || true
            aptly repo create -component=main ${aptlyRepoName}
            aptly repo add ${aptlyRepoName} ${debOutputPath}
        """

        DIST_CODENAMES.each { dist ->
            sh """
                aptly publish drop ${dist} || true

                echo "$GPG_PASSPHRASE" | aptly publish repo \
                    -distribution=${dist} \
                    -passphrase="$GPG_PASSPHRASE" \
                    -architectures=${ARCH} \
                    ${aptlyRepoName}
            """
        }

        sh """
            mkdir -p ${publishDir}
            cp -r ~/.aptly/public/* ${publishDir}/
        """

        sh 'rm -rf ~/.gnupg trust.txt ~/.aptly'

        stash includes: 'dist/artifacts/**/*', name: "${OS}-${ARCH}-deb-package-artifacts"
    }
}

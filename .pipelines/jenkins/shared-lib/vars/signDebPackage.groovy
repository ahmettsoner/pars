def call(String OS, String ARCH, String DIST_CODENAME){
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
        def publishDir = "${env.WORKSPACE}/aptly-publish"

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

        // Create aptly repo and publish
        sh """
            # Clean any existing repo with the same name
            aptly repo drop -force ${aptlyRepoName} || true

            # Create a new local repo
            aptly repo create -distribution=${DIST_CODENAME} -component=main ${aptlyRepoName}
            aptly repo add ${aptlyRepoName} ${debOutputPath}

            # Publish the repo with GPG signing
            aptly publish drop ${DIST_CODENAME} || true

            echo "$GPG_PASSPHRASE" | aptly publish repo \
                -distribution=${DIST_CODENAME} \
                -passphrase="$GPG_PASSPHRASE" \
                -architectures=${ARCH} \
                ${aptlyRepoName}
        """

        // Optionally archive output for upload to Nexus
        sh """
            mkdir -p ${publishDir}
            cp -r ~/.aptly/public/* ${publishDir}/
        """

        // Cleanup
        sh 'rm -rf ~/.gnupg trust.txt ~/.aptly'

        stash name: "${OS}-${ARCH}-aptly-repo", includes: 'aptly-publish/**/*'
    }
}

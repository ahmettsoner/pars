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
        def aptlyRepoName = "${APPNAME}-repo"
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

        // Repo var mı kontrol et, yoksa oluştur
        def repoExists = sh(script: "aptly repo list | grep -w ${aptlyRepoName} || true", returnStdout: true).trim()
        if (!repoExists) {
            sh "aptly repo create -component=main ${aptlyRepoName}"
        }

        // Paketi repo'ya ekle
        sh "aptly repo add ${aptlyRepoName} ${debOutputPath}"

        // Önceki publish edilmiş mimarileri al
        def existingArchitectures = ""
        DIST_CODENAMES.each { dist ->
            def archs = sh(script: "aptly publish list | grep '^${dist}' | awk '{print \$NF}' | grep ${aptlyRepoName} || true", returnStdout: true).trim()
            if (archs) {
                existingArchitectures += archs + ","
            }
        }

        // Mevcut mimarileri topla, tekrar eklemeden birleştir
        existingArchitectures = existingArchitectures.tokenize(',').unique().findAll { it }.join(',')
        if (existingArchitectures) {
            existingArchitectures += ",${ARCH}"
        } else {
            existingArchitectures = ARCH
        }
        existingArchitectures = existingArchitectures.tokenize(',').unique().join(',')

        // Her dist için publish işlemi (önce drop yapılıyor, sonra publish)
        DIST_CODENAMES.each { dist ->
            sh """
                aptly publish drop ${dist} || true

                echo "$GPG_PASSPHRASE" | aptly publish repo \
                    -distribution=${dist} \
                    -passphrase="$GPG_PASSPHRASE" \
                    -architectures=${existingArchitectures} \
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

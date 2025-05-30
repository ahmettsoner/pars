def call(String OS, String ARCH, String DIST_CODENAME){
    withCredentials([
        file(credentialsId: 'public-rpm.gpg', variable: 'GPG_PUBLIC'),
        file(credentialsId: 'private-rpm.gpg', variable: 'GPG_PRIVATE'),
        string(credentialsId: 'GPG_PASSPHRASE', variable: 'GPG_PASSPHRASE'),
        string(credentialsId: 'GPG_FINGERPRINT', variable: 'GPG_FINGERPRINT')
    ]) {

        def newBaseName = "${APPNAME}-${OS}-${ARCH}.deb"
        def debOutputPath = "${env.ARTIFACT_PATH}/${newBaseName}"

        def repoRoot = "${env.WORKSPACE}/apt-repo"
        def poolPath = "${repoRoot}/pool/main/${APPNAME}"
        def distPath = "${repoRoot}/dists/${DIST_CODENAME}/main/binary-${ARCH}"

        sh '''
            mkdir -p ~/.gnupg
            chmod 700 ~/.gnupg

            # Import GPG keys
            gpg --batch --import "$GPG_PUBLIC"
            gpg --batch --import "$GPG_PRIVATE"

            # Get fingerprint
            FPR=$(gpg --list-keys --with-colons | grep '^fpr' | head -n1 | cut -d':' -f10)
            echo "$FPR" > fpr.txt

            # Trust key
            echo "$FPR:6:" > trust.txt
            gpg --import-ownertrust trust.txt

            # GPG config for loopback
            echo "use-agent" > ~/.gnupg/gpg.conf
            echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
            echo "allow-loopback-pinentry" > ~/.gnupg/gpg-agent.conf

            # Restart gpg-agent
            gpgconf --kill gpg-agent
            export GPG_TTY=$(tty || true)
            gpgconf --launch gpg-agent
        '''

        writeFile file: 'sign_deb.expect', text: """
        #!/usr/bin/expect -f

        set timeout -1
        set passphrase [lindex \$argv 0]
        set key [lindex \$argv 1]
        set deb [lindex \$argv 2]

        spawn dpkg-sig --sign builder -k \$key \$deb
        expect {
            "Enter passphrase:" {
                send "\$passphrase\\r"
                exp_continue
            }
            eof
        }
        """

        sh 'chmod +x sign_deb.expect'

        sh "./sign_deb.expect \"$GPG_PASSPHRASE\" \"$GPG_FINGERPRINT\" \"$debOutputPath\""


        // sh "echo \"$GPG_PASSPHRASE\" | dpkg-sig --sign builder -k \"$GPG_FINGERPRINT\" \"$debOutputPath\""
        
        // sh 'rm -rf ~/.gnupg trust.txt'
    }
}

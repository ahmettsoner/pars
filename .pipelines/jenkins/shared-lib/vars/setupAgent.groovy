def setupAgent(String label) {
      stage("Setup ${label} Agent") {
          checkout scm
          sh 'git config --global --add safe.directory $(pwd)'

          sh """
          git config --global user.email "${GIT_USER_EMAIL}"
          git config --global user.name "${GIT_USER_NAME}"

          CURRENT_BASE_VERSION=\$(grm flow phase "${CHANNEL}" --next --print=base)
          TAG_NAME=\$(grm flow phase "${CHANNEL}" --current)
          BUILD_VERSION=\$TAG_NAME

          make build.cmake.\$BUILD_VERSION

          echo "BUILD_VERSION=\$BUILD_VERSION" > version_output.txt
          echo "CURRENT_BASE_VERSION=\$CURRENT_BASE_VERSION" >> version_output.txt
          echo "CHANGELOG_PATH=CHANGELOG/\$BUILD_VERSION.md" >> version_output.txt
          echo "APPNAME=pars" >> version_output.txt
          """
      }
}

return this
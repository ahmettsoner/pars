def call(String BRANCH) {
  checkout scm
  sh 'git config --global --add safe.directory $(pwd)'

  sh """
    git config --global user.email "${GIT_USER_EMAIL}"
    git config --global user.name "${GIT_USER_NAME}"
    git fetch --tags
  """
    // Remote branch ve tag'leri fetch et
    sh 'git fetch --all --tags'

    def branchExists = sh(script: "git ls-remote --heads origin ${BRANCH}", returnStatus: true) == 0
    if (branchExists) {
        sh "git checkout ${BRANCH}"
    } else {
        sh "git checkout -b dev"
    }

    // Git tag'lerini listele (opsiyonel log)
    sh 'git tag -l'

  def currentBaseVersion = sh(
    script: 'grm flow phase "${CHANNEL}" --next --print=base',
    returnStdout: true
  ).trim()

  def tagName = sh(
    script: 'grm flow phase "${CHANNEL}" --current',
    returnStdout: true
  ).trim()

  def buildVersion = tagName

  env.BUILD_VERSION = tagName
  env.CURRENT_BASE_VERSION = currentBaseVersion
  env.CURRENT_BASE_VERSION_RAW = env.CURRENT_BASE_VERSION?.startsWith('v') ? env.CURRENT_BASE_VERSION.substring(1) : env.CURRENT_BASE_VERSION
  env.CHANGELOG_PATH = "CHANGELOG/${buildVersion}.md"
  env.HTML_OUTPUT_DIR = "temp/${env.BUILD_VERSION}/html_docs"
  env.DIST_PATH = "dist"
  env.ARTIFACT_PATH = "dist/artifacts/${env.BUILD_VERSION}"
  env.ARTIFACT_CHECKSUM_MD5_PATH = "${env.ARTIFACT_PATH}/Checksums.md"

  echo "Creating directory: ${env.ARTIFACT_PATH}"
  sh "mkdir -p '${env.ARTIFACT_PATH}'"

  sh "make build.cmake.${buildVersion}"
}


return this
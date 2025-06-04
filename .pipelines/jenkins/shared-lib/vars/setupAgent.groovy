def call(String BRANCH) {
  checkout scm
  sh "git -C '${env.WORKSPACE}' config --global --add safe.directory '${env.WORKSPACE}'"

  sh """
    git config --global user.email "${GIT_USER_EMAIL}"
    git config --global user.name "${GIT_USER_NAME}"
    git fetch --all --tags
  """

  def branchExists = sh(script: "git ls-remote --heads origin ${BRANCH}", returnStatus: true) == 0
  if (branchExists) {
      // BRANCH'i doğrudan remote'dan sabitle
      sh "git checkout -B ${BRANCH} origin/${BRANCH}"
  } else {
      sh "git checkout -b dev"
  }

  // Artık BRANCH güncel, merged tag'leri görmek mümkün
  sh "git tag --merged ${BRANCH}"


  def currentBaseVersion = sh(
    script: 'grm flow phase "${CHANNEL}" --next --print=base',
    returnStdout: true
  ).trim()

  def tagName = sh(
    script: 'grm flow phase "${CHANNEL}" --next',
    returnStdout: true
  ).trim()


  env.BUILD_VERSION = tagName
  def suffix = tagName.replaceFirst("^${currentBaseVersion}-?", "")
  env.BUILD_VERSION_RELEASE_NUMBER = suffix
  env.CURRENT_BASE_VERSION = currentBaseVersion
  env.CURRENT_BASE_VERSION_RAW = env.CURRENT_BASE_VERSION?.startsWith('v') ? env.CURRENT_BASE_VERSION.substring(1) : env.CURRENT_BASE_VERSION
  env.HTML_OUTPUT_DIR = "temp/${env.BUILD_VERSION}/html_docs"
  env.DIST_PATH = "dist"
  env.ARTIFACT_PATH = "dist/artifacts/${env.BUILD_VERSION}"
  env.CHANGES_PATH = ".changes"
  env.CHANGELOG_PATH = "${env.CHANGES_PATH}/${env.BUILD_VERSION}.md"

  echo "BUILD_VERSION: ${env.BUILD_VERSION}"


  echo "Creating directory: ${env.ARTIFACT_PATH}"
  sh "mkdir -p '${env.CHANGES_PATH}'"
  sh "mkdir -p '${env.ARTIFACT_PATH}'"

  sh "make build.cmake.${env.BUILD_VERSION}"
}


return this
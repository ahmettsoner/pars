def setupAgent() {
  checkout scm
  sh 'git config --global --add safe.directory $(pwd)'

  sh """
    git config --global user.email "${GIT_USER_EMAIL}"
    git config --global user.name "${GIT_USER_NAME}"
  """

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

  sh "make build.cmake.${buildVersion}"
}


return this
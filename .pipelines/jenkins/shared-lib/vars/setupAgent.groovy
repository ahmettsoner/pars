def call(String BRANCH) {
  checkout scm
  sh "git -C '${env.WORKSPACE}' config --global --add safe.directory '${env.WORKSPACE}'"

  sh """
    git config --global user.email "${GIT_USER_EMAIL}"
    git config --global user.name "${GIT_USER_NAME}"
    git fetch --prune origin "+refs/tags/*:refs/tags/*" || echo "No tags to fetch or fetch failed, skipping"
  """

  def branchExists = sh(script: "git ls-remote --heads origin ${BRANCH}", returnStatus: true) == 0
  if (branchExists) {
      // BRANCH'i doğrudan remote'dan sabitle
      sh "git checkout -B ${BRANCH} origin/${BRANCH}"
  } else {
      sh "git checkout -b dev"
  }


  if (!env.CURRENT_VERSION) {
    def currentBaseVersion = sh(
      script: 'grm flow phase "${CHANNEL}" --next --print=base',
      returnStdout: true
    ).trim()

    def tagName = sh(
      script: 'grm flow phase "${CHANNEL}" --next',
      returnStdout: true
    ).trim()

    sh "git tag ${tagName} -m \"release ${tagName}\""
    withCredentials([usernamePassword(credentialsId: 'gitea-creds', usernameVariable: 'GIT_USER', passwordVariable: 'GIT_PASS')]) {
        sh """
            git push http://${GIT_USER}:${GIT_PASS}@192.168.118.47:3030/admin/pars.git ${tagName}
        """
    }
    env.CURRENT_VERSION = tagName
    env.CURRENT_BASE_VERSION = currentBaseVersion
  }

  def suffix = CURRENT_VERSION.replaceFirst("^${env.CURRENT_BASE_VERSION}-?", "")
  env.CURRENT_VERSION_RELEASE_NUMBER = suffix
  env.CURRENT_BASE_VERSION_RAW = env.CURRENT_BASE_VERSION?.startsWith('v') ? env.CURRENT_BASE_VERSION.substring(1) : env.CURRENT_BASE_VERSION
  env.HTML_OUTPUT_DIR = "temp/${env.CURRENT_VERSION}/html_docs"
  env.DIST_PATH = "dist"
  env.ARTIFACT_PATH = "dist/artifacts/${env.CURRENT_VERSION}"
  env.CHANGES_PATH = ".changes"
  env.CHANGELOG_PATH = "${env.CHANGES_PATH}/${env.CURRENT_VERSION}.md"

  echo "CURRENT_VERSION: ${env.CURRENT_VERSION}"


  echo "Creating directory: ${env.ARTIFACT_PATH}"
  sh "mkdir -p '${env.CHANGES_PATH}'"
  sh "mkdir -p '${env.ARTIFACT_PATH}'"

  sh "make build.cmake.${env.CURRENT_VERSION}"
}


return this
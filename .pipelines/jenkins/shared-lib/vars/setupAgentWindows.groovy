def call(String BRANCH) {
  checkout scm
  powershell """
    git -C "${env.WORKSPACE}" config --global --add safe.directory "${env.WORKSPACE}"
    git config --global user.email "${GIT_USER_EMAIL}"
    git config --global user.name "${GIT_USER_NAME}"
    try {
      git fetch --prune origin "+refs/tags/*:refs/tags/*"
    } catch {
      Write-Host "No tags to fetch or fetch failed, skipping"
    }
  """

  def branchExists = powershell(
    script: "git ls-remote --heads origin ${BRANCH}; exit \$LASTEXITCODE",
    returnStatus: true
  ) == 0

  if (branchExists) {
    powershell "git checkout -B ${BRANCH} origin/${BRANCH}"
  } else {
    powershell "git checkout -b dev"
  }

  if (!env.CURRENT_VERSION) {
    def currentBaseVersion = powershell(
      script: 'grm flow phase "${CHANNEL}" --next --print=base',
      returnStdout: true
    ).trim()

    def tagName = powershell(
      script: 'grm flow phase "${CHANNEL}" --next',
      returnStdout: true
    ).trim()

    powershell "git tag ${tagName} -m \"release ${tagName}\""

    withCredentials([usernamePassword(credentialsId: 'gitea-creds', usernameVariable: 'GIT_USER', passwordVariable: 'GIT_PASS')]) {
      powershell """
        git push http://${env.GIT_USER}:${env.GIT_PASS}@${env.GITEA_BASE_URL}/admin/pars.git ${tagName}
      """
    }

    env.CURRENT_VERSION = tagName
    env.CURRENT_BASE_VERSION = currentBaseVersion
  }

  def suffix = env.CURRENT_VERSION.replaceFirst("^${env.CURRENT_BASE_VERSION}-?", "")
  env.CURRENT_VERSION_RELEASE_NUMBER = suffix
  env.CURRENT_BASE_VERSION_RAW = env.CURRENT_BASE_VERSION?.startsWith('v') ? env.CURRENT_BASE_VERSION.substring(1) : env.CURRENT_BASE_VERSION
  env.HTML_OUTPUT_DIR = "temp/${env.CURRENT_VERSION}/html_docs"
  env.DIST_PATH = "dist"
  env.ARTIFACT_PATH = "dist/artifacts/${env.CURRENT_VERSION}"
  env.CHANGES_PATH = ".changes"
  env.CHANGELOG_PATH = "${env.CHANGES_PATH}/${env.CURRENT_VERSION}.md"

  echo "CURRENT_VERSION: ${env.CURRENT_VERSION}"

  echo "Creating directory: ${env.ARTIFACT_PATH}"
  powershell "New-Item -ItemType Directory -Force -Path '${env.CHANGES_PATH}'"
  powershell "New-Item -ItemType Directory -Force -Path '${env.ARTIFACT_PATH}'"

  powershell "make build.cmake.${env.CURRENT_VERSION}"
}

return this

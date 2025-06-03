def initOnce(String BRANCH) {
  def marker = "${env.WORKSPACE}/.initdone"
  if (!fileExists(marker)) {
    echo "First time init on this agent"
    // setupAgent.call(BRANCH)
    writeFile file: marker, text: 'done'
  } else {
    echo "Init already done, skipping"
  }
  setupAgent.call(BRANCH)
}

def cleanupOnce() {
  def marker = "${env.WORKSPACE}/.cleanupdone"
  if (!fileExists(marker)) {
    echo "Performing cleanup"
    // Cleanup işlemleri
    sh 'echo "Cleaning up workspace..."'
    writeFile file: marker, text: 'done'
  } else {
    echo "Cleanup already done, skipping"
  }
  sh "rm -rf ${env.DIST_PATH}/*"
}

def initOnce(String BRANCH) {
  def marker = "${env.WORKSPACE}/.initdone"
  if (!fileExists(marker)) {
    echo "First time init on this agent"
    // Buraya init adımlarını ekleyin, örnek:
    sh 'echo "Installing dependencies..."'
    sh setupAgent.call(BRANCH)
    writeFile file: marker, text: 'done'
  } else {
    echo "Init already done, skipping"
  }
}

def cleanupOnce() {
  def marker = "${env.WORKSPACE}/.cleanupdone-${name}"
  if (!fileExists(marker)) {
    echo "Performing cleanup"
    // Cleanup işlemleri
    sh 'echo "Cleaning up workspace..."'
    sh "rm -rf ${env.DIST_PATH}/*"
    writeFile file: marker, text: 'done'
  } else {
    echo "Cleanup already done, skipping"
  }
}

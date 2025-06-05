def initOnce(String BRANCH) {
  def marker = "${env.WORKSPACE}/.initdone"
  if (!fileExists(marker)) {
    echo "First time init on this agent"
    // Buraya init adımlarını ekleyin, örnek:
    // setupAgent.call(BRANCH)
    sh 'echo "Installing dependencies..."'
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
def initOnceWindows(String BRANCH) {
  def marker = "${env.WORKSPACE}\\.initdone"
  if (!fileExists(marker)) {
    echo "First time init on this agent"
    // Buraya init adımlarını ekleyin, örnek:
    // setupAgent.call(BRANCH)
    powershell 'Write-Host "Installing dependencies..."'
    writeFile file: marker, text: 'done'
  } else {
    echo "Init already done, skipping"
  }
  setupAgentWindows.call(BRANCH)
}

def cleanupOnceWindows() {
  def marker = "${env.WORKSPACE}\\.cleanupdone"
  if (!fileExists(marker)) {
    echo "Performing cleanup"
    // Cleanup işlemleri
    powershell 'Write-Host "Cleaning up workspace..."'
    writeFile file: marker, text: 'done'
  } else {
    echo "Cleanup already done, skipping"
  }

  // DIST_PATH içeriğini temizle
  powershell """
    \$distPath = "${env.DIST_PATH}"
    if (Test-Path \$distPath) {
      Get-ChildItem -Path \$distPath -Recurse -Force | Remove-Item -Force -Recurse
    }
  """
}

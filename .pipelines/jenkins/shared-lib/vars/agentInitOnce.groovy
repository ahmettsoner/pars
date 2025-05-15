def initOnce(String name = 'default') {
  def marker = "${env.WORKSPACE}/.initdone-${name}"
  if (!fileExists(marker)) {
    echo "[${name}] First time init on this agent"
    // Buraya init adımlarını ekleyin, örnek:
    sh '''
      echo "Installing dependencies..."
      sleep 1
    '''
    writeFile file: marker, text: 'done'
  } else {
    echo "[${name}] Init already done, skipping"
  }
}

def cleanupOnce(String name = 'default') {
  def marker = "${env.WORKSPACE}/.cleanupdone-${name}"
  if (!fileExists(marker)) {
    echo "[${name}] Performing cleanup"
    // Cleanup işlemleri
    sh '''
      echo "Cleaning up workspace..."
      sleep 1
    '''
    writeFile file: marker, text: 'done'
  } else {
    echo "[${name}] Cleanup already done, skipping"
  }
}

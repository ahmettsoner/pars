def generateDocs(String label) {
  def versionOutput = readFile("version_output.txt")
  def buildVersion = versionOutput.split('\n').find { it.startsWith('BUILD_VERSION=') }?.split('=')[1]?.trim()

  if (!buildVersion) {
      error "BUILD_VERSION is empty or null"
  }

  def HTML_OUTPUT_DIR = "temp/${buildVersion}/html_docs"

  sh """
      mkdir -p "${HTML_OUTPUT_DIR}"
      cp -r docs/assets/ "${HTML_OUTPUT_DIR}"

      find ./docs -name '*.md' | while read -r FILE; do
          RELATIVE_PATH="\${FILE#./docs/}"
          HTML_FILE="\${RELATIVE_PATH%.md}.html"
          DIR_PATH="\$(dirname \${RELATIVE_PATH})"
          mkdir -p "${HTML_OUTPUT_DIR}/\${DIR_PATH}"
          pandoc --standalone --css="../assets/style/style.css" -o "${HTML_OUTPUT_DIR}/\${HTML_FILE}" --lua-filter=scripts/change_links.lua \${FILE}
      done

      echo "HTML docs generated at: ${HTML_OUTPUT_DIR}"
      echo "HTML_OUTPUT_DIR=${HTML_OUTPUT_DIR}" >> version_output.txt
  """
}

return this
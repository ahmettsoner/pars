def call() {
    if (!env.BUILD_VERSION) {
        error "BUILD_VERSION is empty or null"
    }


    sh """
        mkdir -p "${env.HTML_OUTPUT_DIR}"
        cp -r docs/assets/ "${env.HTML_OUTPUT_DIR}"

        find ./docs -name '*.md' | while read -r FILE; do
            RELATIVE_PATH="\${FILE#./docs/}"
            HTML_FILE="\${RELATIVE_PATH%.md}.html"
            DIR_PATH="\$(dirname \${RELATIVE_PATH})"
            mkdir -p "${env.HTML_OUTPUT_DIR}/\${DIR_PATH}"
            pandoc --standalone --css="../assets/style/style.css" -o "${env.HTML_OUTPUT_DIR}/\${HTML_FILE}" --lua-filter=scripts/change_links.lua \${FILE}
        done

        echo "HTML docs generated at: ${env.HTML_OUTPUT_DIR}"
    """

    stash includes: "${env.HTML_OUTPUT_DIR}/**/*", name: "${OS}-${ARCH}-html-outputdir"
}

return this
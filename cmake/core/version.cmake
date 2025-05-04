function(generate_version_info custom_version version_tag)
    if(NOT custom_version STREQUAL "NONE")
        set(${version_tag} "${custom_version}" PARENT_SCOPE)
    else()
        set(DEFAULT_VERSION "v1.0.0-test.1")

        execute_process(
            COMMAND date "+%Y.%m.%d"
            OUTPUT_VARIABLE CURRENT_DATE
            OUTPUT_STRIP_TRAILING_WHITESPACE
        )
        execute_process(
            COMMAND date "+%Y.%m.%d"
            OUTPUT_VARIABLE CURRENT_DATE
            OUTPUT_STRIP_TRAILING_WHITESPACE
        )
        set(DEFAULT_DATE ${CURRENT_DATE})

        execute_process(
            COMMAND git describe --tags --abbrev=0
            RESULT_VARIABLE GIT_TAG_FOUND
            OUTPUT_VARIABLE GIT_TAG
            OUTPUT_STRIP_TRAILING_WHITESPACE
        )

        if(NOT GIT_TAG_FOUND EQUAL 0)
            message(STATUS "No Git tag found, using default values.") 
            set(${version_tag} ${DEFAULT_VERSION} PARENT_SCOPE)
        else()
            set(${version_tag} "${GIT_TAG}" PARENT_SCOPE)
        endif()
    endif()
endfunction()

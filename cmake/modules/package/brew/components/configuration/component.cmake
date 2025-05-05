set(CMAKE_SOURCE_DIR_PATH ${CMAKE_SOURCE_DIR})
set(COMMON_VARIABLES 
    PROJECT_NAME
    APP_NAME
    APP_TAG
    RAW_VERSION
    BREW_SERIES
    BREW_BASE
    CHANGELOG_PATH              # d
    BREW_PACK_TYPE
    PROJECT_GIT
    RELEASE_DATE_BREW
    PROJECT_HOMEPAGE
    PROJECT_LICENCE_TYPE
    PROJECT_DESCRIPTION
    PROJECT_SUMMARY
    LINUX_APP_BINARY_DIR
    LINUX_APP_DATA_DATABASE_DIR
    CMAKE_SOURCE_DIR_PATH
    DIST_ROOT_DIR
    GOOS
    EXT
    )

file(GLOB_RECURSE BREW_FILES "${CMAKE_CURRENT_LIST_DIR}/brew-files/*")

foreach(BREWARCH ${ALL_BREWARCH_LIST_LINUX})
    map_brewarch_to_arch_all(${BREWARCH} APP_ARCH)
    
    set(BREW_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${OS_MACOS}/pkg/${BREW_PACKAGE_NAME}/${APP_ARCH})
    set(BREW_PAYLOAD_DIR ${BREW_ROOT_DIR}/${APP_NAME})
    set(BREW_OUTPUT_DIR ${BREW_ROOT_DIR}/output)
    set(BREW_CONF_DIR ${BREW_ROOT_DIR}/${APP_NAME}/brew)

    if(${BREWARCH} STREQUAL ${BREW_ARCH_ALL})
        get_host_arch(HOST_ARCH)
        set(BIN_OUTPUT_FULL_PATH ${BREW_OUTPUT_DIR}/${APP_NAME}/${DIST_ROOT_DIR}/${APP_TAG}/${OS_MACOS}/bin/${HOST_ARCH}/${APP_NAME}${EXT})
    else()
        set(BIN_OUTPUT_FULL_PATH ${BREW_OUTPUT_DIR}/${APP_NAME}/${DIST_ROOT_DIR}/${APP_TAG}/${OS_MACOS}/bin/${APP_ARCH}/${APP_NAME}${EXT})
    endif()

    list(APPEND COMMON_VARIABLES APP_ARCH)
    list(APPEND COMMON_VARIABLES BREWARCH)
    list(APPEND COMMON_VARIABLES BIN_OUTPUT_FULL_PATH)

    set(BREW_FILE_NAMES "")
    foreach(BREWFILE ${BREW_FILES})
        file(RELATIVE_PATH REL_FILE_PATH "${CMAKE_CURRENT_LIST_DIR}/brew-files" ${BREWFILE})

        if (REL_FILE_PATH STREQUAL "Formula/config.rb")
            if (IS_PRERELEASE)
                set(CONFIG_FILE_PATH "${BREW_CONF_DIR}/Formula/${APP_NAME}@${VERSION_SEMVER_BUILD}.rb")
            else()
                if(VERSION_PATCH GREATER 0)
                    set(CONFIG_FILE_PATH "${BREW_CONF_DIR}/Formula/${APP_NAME}@${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH}.rb")
                else()
                    set(CONFIG_FILE_PATH "${BREW_CONF_DIR}/Formula/${APP_NAME}@${VERSION_MAJOR}.${VERSION_MINOR}.rb")
                endif()
            endif()
        else()
            set(CONFIG_FILE_PATH "${BREW_CONF_DIR}/${REL_FILE_PATH}")
        endif()
        list(APPEND BREW_FILE_NAMES ${CONFIG_FILE_PATH})
        list(APPEND COMMON_VARIABLES CONFIG_FILE_PATH)

        var_list_to_cmake_args(VARIABLES_TO_PASS "${COMMON_VARIABLES}")
        add_custom_command(
            OUTPUT ${CONFIG_FILE_PATH}
            COMMAND ${CMAKE_COMMAND} -E echo "Generating ${BREWFILE} file..."
            COMMAND ${CMAKE_COMMAND} ${VARIABLES_TO_PASS}  -P "${BREWFILE}"
            COMMENT "Generating ${BREWFILE} to ${CONFIG_FILE_PATH}"
        )
    endforeach()

    add_custom_target(build.brew.package.${APP_ARCH}.configuration DEPENDS check_env_for_brew_packing ${BREW_FILE_NAMES})
endforeach()
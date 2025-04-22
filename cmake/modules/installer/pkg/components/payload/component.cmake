set_os_ext(${OS_MACOS} EXT)
set(PAYLOADS 
    CMakeLists.txt
    .config
    cmake
    Makefile
    src
    docs
)

foreach(PKGARCH ${ALL_PKGARCH_LIST_OS_MACOS})
    map_pkgarch_to_arch_all(${PKGARCH} APP_ARCH)
    set(PKG_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${OS_MACOS}/ins/${PKG_PACKAGE_NAME}/${APP_ARCH})
    set(PKG_PAYLOAD_DIR ${PKG_ROOT_DIR}/${APP_NAME})
    set(PKG_TEMP_DIR ${PKG_ROOT_DIR}/temp)


    set(PAYLOAD_OUTPUTS "")

    command_for_shell("powershell" "if (-not (Test-Path \"${PKG_TEMP_DIR}\")) { New-Item -Path \"${PKG_TEMP_DIR}\" -ItemType Directory }" SHELL_GO_BUILD_COMMAND_CREATE_FOLDER)
    add_custom_command(
        OUTPUT ${PKG_TEMP_DIR}
        COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_FOLDER}
        COMMENT "Creating payloads folder ${PKG_PAYLOAD_DIR}"
    )


    foreach(PAYLOAD ${PAYLOADS})
        list(APPEND PAYLOAD_OUTPUTS ${PKG_TEMP_DIR}/${PAYLOAD})
        command_for_shell("powershell" "Copy-Item -Recurse -Force '${SOURCE_ROOT_DIR}/${PAYLOAD}' '${PKG_TEMP_DIR}'" SHELL_GO_BUILD_COMMAND)
        add_custom_command(
            OUTPUT ${PKG_TEMP_DIR}/${PAYLOAD}
            COMMAND ${SHELL_GO_BUILD_COMMAND}
            VERBATIM
            COMMENT "Copying ${PAYLOAD} to ${PKG_TEMP_DIR}"
        )
    endforeach()

    add_custom_command(
        OUTPUT ${PKG_TEMP_DIR}/src/vendor
        COMMAND cd ${PKG_TEMP_DIR}/src && go mod tidy
        COMMAND cd ${PKG_TEMP_DIR}/src && go mod vendor
        COMMENT "Preparing payloads to ${PKG_TEMP_DIR}"
    )

    add_custom_command(
        OUTPUT ${PKG_TEMP_DIR}/${APP_NAME}${EXT}
        COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_SOURCES_FOLDER}
        COMMAND make build.cmake.macos VERSION=${APP_TAG}
        COMMAND make build.binary.${APP_ARCH} OUTPUT=${PKG_PAYLOAD_DIR}/${APP_NAME}${EXT}
        VERBATIM
        WORKING_DIRECTORY ${PKG_TEMP_DIR}
        COMMENT "Creating binary to ${PKG_PAYLOAD_DIR}"
    )
add_custom_target(build.pkg.package.${APP_ARCH}.payload DEPENDS check_env_for_pkg_packing ${PKG_TEMP_DIR} ${PAYLOAD_OUTPUTS} ${PKG_TEMP_DIR}/src/vendor ${PKG_TEMP_DIR}/${APP_NAME}${EXT})
endforeach()
get_host_os(HOST_OS)
set_os_ext(${HOST_OS} EXT)
set(PAYLOADS 
    CMakeLists.txt
    .config
    cmake
    Makefile
    src
    docs
)

foreach(MSIARCH ${ALL_MSIARCH_LIST_WINDOWS})
    map_msiarch_to_arch_all(${MSIARCH} APP_ARCH)
    set(MSI_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/ins/${MSI_PACKAGE_NAME}/${APP_ARCH})
    set(MSI_PAYLOAD_DIR ${MSI_ROOT_DIR}/${APP_NAME})
    set(MSI_TEMP_DIR ${MSI_ROOT_DIR}/temp)


    set(PAYLOAD_OUTPUTS "")

    command_for_shell("powershell" "if (-not (Test-Path \"${MSI_TEMP_DIR}\")) { New-Item -Path \"${MSI_TEMP_DIR}\" -ItemType Directory }" SHELL_GO_BUILD_COMMAND_CREATE_FOLDER)
    add_custom_command(
        OUTPUT ${MSI_TEMP_DIR}
        COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_FOLDER}
        COMMENT "Creating payloads folder ${MSI_PAYLOAD_DIR}"
    )


    foreach(PAYLOAD ${PAYLOADS})
        list(APPEND PAYLOAD_OUTPUTS ${MSI_TEMP_DIR}/${PAYLOAD})
        command_for_shell("powershell" "Copy-Item -Recurse -Force '${SOURCE_ROOT_DIR}/${PAYLOAD}' '${MSI_TEMP_DIR}'" SHELL_GO_BUILD_COMMAND)
        add_custom_command(
            OUTPUT ${MSI_TEMP_DIR}/${PAYLOAD}
            COMMAND ${SHELL_GO_BUILD_COMMAND}
            VERBATIM
            COMMENT "Copying ${PAYLOAD} to ${MSI_TEMP_DIR}"
        )
    endforeach()

    add_custom_command(
        OUTPUT ${MSI_TEMP_DIR}/src/vendor
        COMMAND cd ${MSI_TEMP_DIR}/src && go mod tidy
        COMMAND cd ${MSI_TEMP_DIR}/src && go mod vendor
        COMMENT "Preparing payloads to ${MSI_TEMP_DIR}"
    )

    add_custom_command(
        OUTPUT ${MSI_TEMP_DIR}/${APP_NAME}${EXT}
        COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_SOURCES_FOLDER}
        COMMAND make build.cmake.windows VERSION=${APP_TAG}
        COMMAND make build.binary.${APP_ARCH} OUTPUT=${MSI_PAYLOAD_DIR}/${APP_NAME}${EXT}
        VERBATIM
        WORKING_DIRECTORY ${MSI_TEMP_DIR}
        COMMENT "Creating binary to ${MSI_PAYLOAD_DIR}"
    )
add_custom_target(build.msi.package.${APP_ARCH}.payload DEPENDS check_env_for_msi_packing ${MSI_TEMP_DIR} ${PAYLOAD_OUTPUTS} ${MSI_TEMP_DIR}/src/vendor ${MSI_TEMP_DIR}/${APP_NAME}${EXT})
endforeach()
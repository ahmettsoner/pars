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

foreach(CHOCOARCH ${ALL_CHOCOARCH_LIST_WINDOWS})
    map_chocoarch_to_arch_all(${CHOCOARCH} APP_ARCH)
    set(CHOCO_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${CHOCO_PACKAGE_NAME}/${APP_ARCH})
    set(CHOCO_PAYLOAD_DIR ${CHOCO_ROOT_DIR}/${APP_NAME}/tools)
    set(CHOCO_TEMP_DIR ${CHOCO_ROOT_DIR}/temp)


    set(PAYLOAD_OUTPUTS "")

    command_for_shell("powershell" "if (-not (Test-Path \"${CHOCO_TEMP_DIR}\")) { New-Item -Path \"${CHOCO_TEMP_DIR}\" -ItemType Directory }" SHELL_GO_BUILD_COMMAND_CREATE_FOLDER)
    add_custom_command(
        OUTPUT ${CHOCO_TEMP_DIR}
        COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_FOLDER}
        COMMENT "Creating payloads folder ${CHOCO_PAYLOAD_DIR}"
    )


    foreach(PAYLOAD ${PAYLOADS})
        list(APPEND PAYLOAD_OUTPUTS ${CHOCO_TEMP_DIR}/${PAYLOAD})
        command_for_shell("powershell" "Copy-Item -Recurse -Force '${SOURCE_ROOT_DIR}/${PAYLOAD}' '${CHOCO_TEMP_DIR}'" SHELL_GO_BUILD_COMMAND)
        add_custom_command(
            OUTPUT ${CHOCO_TEMP_DIR}/${PAYLOAD}
            COMMAND ${SHELL_GO_BUILD_COMMAND}
            VERBATIM
            COMMENT "Copying ${PAYLOAD} to ${CHOCO_TEMP_DIR}"
        )
    endforeach()

    add_custom_command(
        OUTPUT ${CHOCO_TEMP_DIR}/src/vendor
        COMMAND cd ${CHOCO_TEMP_DIR}/src && go mod tidy
        COMMAND cd ${CHOCO_TEMP_DIR}/src && go mod vendor
        COMMENT "Preparing payloads to ${CHOCO_TEMP_DIR}"
    )

    set(ALL_BUILD_OUTPUTS "")
    if(${APP_ARCH} STREQUAL ${CHOCO_ARCH_ALL})
        foreach(CHOCOARCH ${CHOCOARCH_LIST_WINDOWS})
            map_chocoarch_to_arch(${CHOCOARCH} APP_ARCH_FOR_ALL)
            list(APPEND ALL_BUILD_OUTPUTS ${CHOCO_TEMP_DIR}/${APP_NAME}-${APP_ARCH_FOR_ALL}${EXT})
            add_custom_command(
                OUTPUT ${CHOCO_TEMP_DIR}/${APP_NAME}-${APP_ARCH_FOR_ALL}${EXT}
                COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_SOURCES_FOLDER}
                COMMAND make build.cmake.windows VERSION=${APP_TAG}
                COMMAND make build.binary.${APP_ARCH_FOR_ALL} OUTPUT=${CHOCO_PAYLOAD_DIR}/${APP_NAME}-${APP_ARCH_FOR_ALL}${EXT}
                VERBATIM
                WORKING_DIRECTORY ${CHOCO_TEMP_DIR}
                COMMENT "Creating binary to ${CHOCO_PAYLOAD_DIR}"
            )
        endforeach()
    else()
        list(APPEND ALL_BUILD_OUTPUTS ${CHOCO_TEMP_DIR}/${APP_NAME}${EXT})
        add_custom_command(
            OUTPUT ${CHOCO_TEMP_DIR}/${APP_NAME}${EXT}
            COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_SOURCES_FOLDER}
            COMMAND make build.cmake.windows VERSION=${APP_TAG}
            COMMAND make build.binary.${APP_ARCH} OUTPUT=${CHOCO_PAYLOAD_DIR}/${APP_NAME}${EXT}
            VERBATIM
            WORKING_DIRECTORY ${CHOCO_TEMP_DIR}
            COMMENT "Creating binary to ${CHOCO_PAYLOAD_DIR}"
        )
    endif()

    add_custom_target(build.choco.package.${APP_ARCH}.payload DEPENDS check_env_for_choco_packing ${CHOCO_TEMP_DIR} ${PAYLOAD_OUTPUTS} ${CHOCO_TEMP_DIR}/src/vendor ${ALL_BUILD_OUTPUTS})
endforeach()
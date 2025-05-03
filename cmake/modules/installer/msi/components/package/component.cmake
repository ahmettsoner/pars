get_host_os(HOST_OS)

foreach(MSIARCH ${ALL_MSIARCH_LIST_WINDOWS})
    map_msiarch_to_arch_all(${MSIARCH} APP_ARCH)
    set(MSI_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/ins/${MSI_PACKAGE_NAME}/${APP_ARCH})
    set(MSI_PAYLOAD_DIR ${MSI_ROOT_DIR}/${APP_NAME})
    set(MSI_OUTPUT_DIR ${MSI_ROOT_DIR}/output)

    command_for_shell("powershell" "if (-not (Test-Path \"${MSI_OUTPUT_DIR}\")) { New-Item -Path \"${MSI_OUTPUT_DIR}\" -ItemType Directory }" SHELL_GO_BUILD_COMMAND_CREATE_FOLDER)
    add_custom_command(
        OUTPUT ${MSI_OUTPUT_DIR}
        COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_FOLDER}
        COMMENT "Creating payloads folder ${MSI_PAYLOAD_DIR}"
    )

    add_custom_command(
        OUTPUT ${MSI_OUTPUT_DIR}/${APP_NAME}.msi
        COMMAND ${CMAKE_COMMAND} -E echo "Building source files."
        COMMAND wix build -o ${MSI_OUTPUT_DIR}/${APP_NAME}.msi ./config.wxs
        WORKING_DIRECTORY ${MSI_PAYLOAD_DIR}
        COMMENT "Building .msi installer"
    )

    add_custom_target(build.msi.package.${APP_ARCH}.package DEPENDS check_env_for_msi_packing ${MSI_OUTPUT_DIR} ${MSI_OUTPUT_DIR}/${APP_NAME}.msi)

    add_custom_target(build.msi.package.${APP_ARCH})
    add_dependencies(build.msi.package.${APP_ARCH} 
        # build.msi.package.setup
        build.msi.package.${APP_ARCH}.configuration
        build.msi.package.${APP_ARCH}.payload
        build.msi.package.${APP_ARCH}.package
    )
endforeach()
foreach(PKGARCH ${ALL_PKGARCH_LIST_MACOS})
    map_pkgarch_to_arch_all(${PKGARCH} APP_ARCH)
    set(PKG_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${OS_MACOS}/ins/${PKG_PACKAGE_NAME}/${APP_ARCH})
    set(PKG_PAYLOAD_DIR ${PKG_ROOT_DIR}/${APP_NAME})
    set(PKG_OUTPUT_DIR ${PKG_ROOT_DIR}/output)

    command_for_shell("powershell" "if (-not (Test-Path \"${PKG_OUTPUT_DIR}\")) { New-Item -Path \"${PKG_OUTPUT_DIR}\" -ItemType Directory }" SHELL_GO_BUILD_COMMAND_CREATE_FOLDER)
    add_custom_command(
        OUTPUT ${PKG_OUTPUT_DIR}
        COMMAND ${SHELL_GO_BUILD_COMMAND_CREATE_FOLDER}
        COMMENT "Creating payloads folder ${PKG_PAYLOAD_DIR}"
    )

    add_custom_command(
        OUTPUT ${PKG_OUTPUT_DIR}/${APP_NAME}.pkg
        COMMAND ${CMAKE_COMMAND} -E echo "Building source files."
        COMMAND wix build -o ${PKG_OUTPUT_DIR}/${APP_NAME}.pkg ./config.wxs
        WORKING_DIRECTORY ${PKG_PAYLOAD_DIR}
        COMMENT "Building .pkg installer"
    )

    add_custom_target(build.pkg.package.${APP_ARCH}.package DEPENDS check_env_for_pkg_packing ${PKG_OUTPUT_DIR} ${PKG_OUTPUT_DIR}/${APP_NAME}.pkg)

    add_custom_target(build.pkg.package.${APP_ARCH})
    add_dependencies(build.pkg.package.${APP_ARCH} 
        # build.pkg.package.setup
        build.pkg.package.${APP_ARCH}.configuration
        # build.pkg.package.${APP_ARCH}.payload
        # build.pkg.package.${APP_ARCH}.package
    )
endforeach()
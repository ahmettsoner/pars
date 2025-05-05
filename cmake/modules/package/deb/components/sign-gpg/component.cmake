get_host_os(HOST_OS)

foreach(DEBARCH ${ALL_DEBARCH_LIST_LINUX})
    map_debarch_to_arch_all(${DEBARCH} APP_ARCH)
    set(DEB_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${DEB_PACKAGE_NAME}/${APP_ARCH})
    set(DEB_PAYLOAD_DIR ${DEB_ROOT_DIR}/${APP_NAME})
    set(DEB_OUTPUT_DIR ${DEB_ROOT_DIR}/output)

    add_custom_command(
        OUTPUT ${DEB_OUTPUT_DIR}/${APP_NAME}.gpg.sign
        WORKING_DIRECTORY ${DEB_OUTPUT_DIR}
        COMMAND ${CMAKE_COMMAND} -E echo "Signing deb with GPG."
        COMMAND gpg --verify *source.changes > ${DEB_OUTPUT_DIR}/${APP_NAME}.gpg.sign 2>&1
        COMMENT "Signing .deb package"
    )

    add_custom_target(build.deb.package.${APP_ARCH}.sign.gpg DEPENDS check_env_for_deb_packing ${DEB_OUTPUT_DIR}/${APP_NAME}.gpg.sign)
endforeach()
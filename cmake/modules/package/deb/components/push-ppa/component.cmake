get_host_os(HOST_OS)

foreach(DEBARCH ${ALL_DEBARCH_LIST_LINUX})
    map_debarch_to_arch_all(${DEBARCH} APP_ARCH)
    set(DEB_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${DEB_PACKAGE_NAME}/${APP_ARCH})
    set(DEB_PAYLOAD_DIR ${DEB_ROOT_DIR}/${APP_NAME})
    set(DEB_OUTPUT_DIR ${DEB_ROOT_DIR}/output)


    add_custom_command(
        OUTPUT ${DEB_OUTPUT_DIR}/*source.ppa.upload
        WORKING_DIRECTORY ${DEB_OUTPUT_DIR}
        COMMAND ${CMAKE_COMMAND} -E echo "Pushing deb to PPA."
        COMMAND dput -c ${DPUT_CONFIG_PATH} launchpad *source.changes
        COMMENT "Pushing .deb package"
    )

    add_custom_target(build.deb.package.${APP_ARCH}.push.ppa DEPENDS check_env_for_deb_packing ${DEB_OUTPUT_DIR}/*source.ppa.upload)
endforeach()
get_host_os(HOST_OS)
set(PAYLOADS 
    CMakeLists.txt
    .config
    cmake
    Makefile
    src
    docs
)

foreach(DEBARCH ${ALL_DEBARCH_LIST_LINUX})
    map_debarch_to_arch_all(${DEBARCH} APP_ARCH)
    set(DEB_PAYLOAD_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${DEB_PACKAGE_NAME}/${APP_ARCH}/${APP_NAME})

    set(PAYLOAD_OUTPUTS "")
    foreach(PAYLOAD ${PAYLOADS})
        list(APPEND PAYLOAD_OUTPUTS ${DEB_PAYLOAD_DIR}/${PAYLOAD})
        add_custom_command(
            OUTPUT ${DEB_PAYLOAD_DIR}/${PAYLOAD}
            COMMAND cp -r ${SOURCE_ROOT_DIR}/${PAYLOAD} ${DEB_PAYLOAD_DIR}/
            COMMENT "Copying payloads to ${DEB_PAYLOAD_DIR}"
        )
    endforeach()

    add_custom_command(
        OUTPUT ${DEB_PAYLOAD_DIR}/src/vendor
        COMMAND cd ${DEB_PAYLOAD_DIR}/src && go mod tidy
        COMMAND cd ${DEB_PAYLOAD_DIR}/src && go mod vendor
        COMMENT "Preparing payloads to ${DEB_PAYLOAD_DIR}"
    )
add_custom_target(build.deb.package.${APP_ARCH}.payload DEPENDS check_env_for_deb_packing ${PAYLOAD_OUTPUTS} ${DEB_PAYLOAD_DIR}/src/vendor)
endforeach()
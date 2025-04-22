get_host_os(HOST_OS)
set(PAYLOADS 
    CMakeLists.txt
    .config
    cmake
    Makefile
    src
    docs
)

foreach(SNAPARCH ${ALL_SNAPARCH_LIST_LINUX})
    map_snaparch_to_arch_all(${SNAPARCH} APP_ARCH)
    set(SNAP_PAYLOAD_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${SNAP_PACKAGE_NAME}/${APP_ARCH}/${APP_NAME})

    set(PAYLOAD_OUTPUTS "")
    foreach(PAYLOAD ${PAYLOADS})
        list(APPEND PAYLOAD_OUTPUTS ${SNAP_PAYLOAD_DIR}/${PAYLOAD})
        add_custom_command(
            OUTPUT ${SNAP_PAYLOAD_DIR}/${PAYLOAD}
            COMMAND cp -r ${SOURCE_ROOT_DIR}/${PAYLOAD} ${SNAP_PAYLOAD_DIR}/
            COMMENT "Copying payloads to ${SNAP_PAYLOAD_DIR}"
        )
    endforeach()

    add_custom_command(
        OUTPUT ${SNAP_PAYLOAD_DIR}/src/vendor
        COMMAND cd ${SNAP_PAYLOAD_DIR}/src && go mod tidy
        COMMAND cd ${SNAP_PAYLOAD_DIR}/src && go mod vendor
        COMMENT "Preparing payloads to ${SNAP_PAYLOAD_DIR}"
    )
add_custom_target(build.snap.package.${APP_ARCH}.payload DEPENDS check_env_for_snap_packing ${PAYLOAD_OUTPUTS} ${SNAP_PAYLOAD_DIR}/src/vendor)
endforeach()
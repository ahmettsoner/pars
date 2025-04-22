get_host_os(HOST_OS)

foreach(SNAPARCH ${ALL_SNAPARCH_LIST_LINUX})
    map_snaparch_to_arch_all(${SNAPARCH} APP_ARCH)
    set(SNAP_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${SNAP_PACKAGE_NAME}/${APP_ARCH})
    set(SNAP_PAYLOAD_DIR ${SNAP_ROOT_DIR}/${APP_NAME})
    set(SNAP_OUTPUT_DIR ${SNAP_ROOT_DIR}/output)

    add_custom_command(
        OUTPUT ${SNAP_OUTPUT_DIR}
        COMMAND ${CMAKE_COMMAND} -E echo "Building source files."
        COMMAND cd ${SNAP_PAYLOAD_DIR} && snapcraft
        COMMAND mkdir -p ${SNAP_OUTPUT_DIR}
        COMMAND mv ${SNAP_ROOT_DIR}/${APP_NAME}/${APP_NAME}_*.snap ${SNAP_OUTPUT_DIR}/
        COMMENT "Building .snap package"
    )

    add_custom_target(build.snap.package.${APP_ARCH}.package DEPENDS check_env_for_snap_packing ${SNAP_OUTPUT_DIR})

    add_custom_target(build.snap.package.${APP_ARCH})
    add_dependencies(build.snap.package.${APP_ARCH} 
        # build.snap.package.setup
        build.snap.package.${APP_ARCH}.configuration
        build.snap.package.${APP_ARCH}.payload
        build.snap.package.${APP_ARCH}.package
    )
endforeach()
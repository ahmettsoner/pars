get_host_os(HOST_OS)

foreach(CHOCOARCH ${ALL_CHOCOARCH_LIST_WINDOWS})
    map_chocoarch_to_arch_all(${CHOCOARCH} APP_ARCH)
    set(CHOCO_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${CHOCO_PACKAGE_NAME}/${APP_ARCH})
    set(CHOCO_PAYLOAD_DIR ${CHOCO_ROOT_DIR}/${APP_NAME})
    set(CHOCO_OUTPUT_DIR ${CHOCO_ROOT_DIR}/output)

    add_custom_command(
        OUTPUT ${CHOCO_OUTPUT_DIR}
        COMMAND ${CMAKE_COMMAND} -E echo "Building source files."
        COMMAND cd ${CHOCO_PAYLOAD_DIR} && choco pack config.nuspec
        COMMAND mkdir -p ${CHOCO_OUTPUT_DIR}
        COMMAND mv ${CHOCO_ROOT_DIR}/${APP_NAME}/${APP_NAME}_*.choco ${CHOCO_OUTPUT_DIR}/
        COMMENT "Building .choco package"
    )

    add_custom_target(build.choco.package.${APP_ARCH}.package DEPENDS check_env_for_choco_packing ${CHOCO_OUTPUT_DIR})

    add_custom_target(build.choco.package.${APP_ARCH})
    add_dependencies(build.choco.package.${APP_ARCH} 
        # build.choco.package.setup
        build.choco.package.${APP_ARCH}.configuration
        build.choco.package.${APP_ARCH}.payload
        build.choco.package.${APP_ARCH}.package
    )
endforeach()
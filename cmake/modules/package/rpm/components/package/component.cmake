get_host_os(HOST_OS)

foreach(RPMARCH ${ALL_RPMARCH_LIST_LINUX})
    map_rpmarch_to_arch_all(${RPMARCH} APP_ARCH)
    set(RPM_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${RPM_PACKAGE_NAME}/${APP_ARCH})
    set(RPM_PAYLOAD_DIR ${RPM_ROOT_DIR}/${APP_NAME})
    set(RPM_OUTPUT_DIR ${RPM_ROOT_DIR}/output)
    set(RPM_CONF_DIR ${RPM_ROOT_DIR}/${APP_NAME})

    add_custom_command(
        OUTPUT ${RPM_OUTPUT_DIR}
        COMMAND ${CMAKE_COMMAND} -E echo "Building source files."
        COMMAND rpmbuild --define "_topdir ${RPM_PAYLOAD_DIR}" -ba ${RPM_CONF_DIR}/SPECS/config.spec
        WORKING_DIRECTORY ${RPM_PAYLOAD_DIR}
        COMMENT "Building .rpm package"
    )

    add_custom_target(build.rpm.package.${APP_ARCH}.package DEPENDS check_env_for_rpm_packing ${RPM_OUTPUT_DIR})

    add_custom_target(build.rpm.package.${APP_ARCH})
    add_dependencies(build.rpm.package.${APP_ARCH} 
        # build.rpm.package.setup
        build.rpm.package.${APP_ARCH}.configuration
        build.rpm.package.${APP_ARCH}.payload
        build.rpm.package.${APP_ARCH}.package
    )
endforeach()
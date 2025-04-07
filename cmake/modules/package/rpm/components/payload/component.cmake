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

foreach(RPMARCH ${ALL_RPMARCH_LIST_LINUX})
    map_rpmarch_to_arch_all(${RPMARCH} APP_ARCH)
    set(RPM_ROOT_DIR ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${HOST_OS}/pkg/${RPM_PACKAGE_NAME}/${APP_ARCH})
    set(RPM_PAYLOAD_DIR ${RPM_ROOT_DIR}/${APP_NAME})
    set(RPM_TEMP_DIR ${RPM_ROOT_DIR}/temp)


    set(PAYLOAD_OUTPUTS "")

    add_custom_command(
        OUTPUT ${RPM_TEMP_DIR}
        COMMAND mkdir -p ${RPM_TEMP_DIR}
        COMMAND mkdir -p ${RPM_PAYLOAD_DIR}/SOURCES
        COMMAND mkdir -p ${RPM_PAYLOAD_DIR}/BUILD
        COMMAND mkdir -p ${RPM_PAYLOAD_DIR}/SPECS
        COMMAND mkdir -p ${RPM_PAYLOAD_DIR}/SOURCES
        COMMAND mkdir -p ${RPM_PAYLOAD_DIR}/RPMS
        COMMAND mkdir -p ${RPM_PAYLOAD_DIR}/SRPMS
        COMMENT "Creating payloads folders"
    )


    foreach(PAYLOAD ${PAYLOADS})
        list(APPEND PAYLOAD_OUTPUTS ${RPM_TEMP_DIR}/${PAYLOAD})
        add_custom_command(
            OUTPUT ${RPM_TEMP_DIR}/${PAYLOAD}
            COMMAND cp -r ${SOURCE_ROOT_DIR}/${PAYLOAD} ${RPM_TEMP_DIR}/
            COMMENT "Copying ${PAYLOAD} to ${RPM_TEMP_DIR}"
        )
    endforeach()

    add_custom_command(
        OUTPUT ${RPM_TEMP_DIR}/src/vendor
        COMMAND go mod tidy
        COMMAND go mod vendor
        WORKING_DIRECTORY ${RPM_TEMP_DIR}/src
        COMMENT "Preparing payloads to ${RPM_TEMP_DIR}"
    )

    add_custom_command(
        OUTPUT ${RPM_PAYLOAD_DIR}/SOURCES/${APP_NAME}-${VERSION_SEMVER}${ARCHIVE_TAR_GZ_EXT}
        COMMAND tar -czf ${RPM_PAYLOAD_DIR}/SOURCES/${APP_NAME}-${VERSION_SEMVER}${ARCHIVE_TAR_GZ_EXT} ./
        WORKING_DIRECTORY ${RPM_TEMP_DIR}
        COMMENT "Creating archieve to ${RPM_PAYLOAD_DIR}"
    )
add_custom_target(build.rpm.package.${APP_ARCH}.payload DEPENDS check_env_for_rpm_packing ${RPM_TEMP_DIR} ${PAYLOAD_OUTPUTS} ${RPM_TEMP_DIR}/src/vendor ${RPM_PAYLOAD_DIR}/SOURCES/${APP_NAME}-${VERSION_SEMVER}${ARCHIVE_TAR_GZ_EXT})
endforeach()
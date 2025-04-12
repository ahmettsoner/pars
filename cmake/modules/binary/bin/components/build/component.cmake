set(ALL_TARGETS "")
set(ALL_TARGETS_VENDOR "")
foreach(GOOS ${GOOS_LIST})

<<<<<<< HEAD
    set_os_ext(${GOOS} EXE)
=======
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
    set_goos_arch_lists(${GOOS})

    set(OS_ALL_TARGETS "")
    set(OS_ALL_TARGETS_VENDOR "")
    foreach(GOARCH ${ARCH_LIST})
        map_goarch_to_arch(${GOARCH} APP_ARCH)
        
        set_build_output_from_arg(BUILD_OUTPUT_PATH)
        build("${GOOS}" "${GOARCH}" "${BUILD_OUTPUT_PATH}")
<<<<<<< HEAD
        generate_build_output_path_tmp(PATH_OUTPUT)
=======
        generate_build_output_path_tmp(GOOS PATH_OUTPUT)
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
        add_custom_target(build.binary.${GOOS}.${APP_ARCH}
            DEPENDS ${PATH_OUTPUT}
        )
        list(APPEND OS_ALL_TARGETS "build.binary.${GOOS}.${APP_ARCH}")
        list(APPEND ALL_TARGETS "build.binary.${GOOS}.${APP_ARCH}")
    endforeach()
        add_custom_target(build.binary.${GOOS}.all
            DEPENDS ${OS_ALL_TARGETS}
        )
endforeach()
add_custom_target(build.binary.all
    DEPENDS ${ALL_TARGETS}
)



get_host_os(GOOS)
<<<<<<< HEAD
set_os_ext(${GOOS} EXE)
=======
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c

get_host_arch(APP_ARCH)
map_arch_to_goarch(${APP_ARCH} GOARCH)


set_build_output_from_arg(BUILD_OUTPUT_PATH)
build("${GOOS}" "${GOARCH}" "${BUILD_OUTPUT_PATH}")
<<<<<<< HEAD
generate_build_output_path_tmp(PATH_OUTPUT)
=======
generate_build_output_path_tmp(GOOS PATH_OUTPUT)
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
add_custom_target(build.binary
    DEPENDS ${PATH_OUTPUT}
)

set_goos_arch_lists(${GOOS})
foreach(GOARCH ${ARCH_LIST})
    map_goarch_to_arch(${GOARCH} APP_ARCH)

    set_build_output_from_arg(BUILD_OUTPUT_PATH)
    build("${GOOS}" "${GOARCH}" "${BUILD_OUTPUT_PATH}")
<<<<<<< HEAD
    generate_build_output_path_tmp(PATH_OUTPUT)
=======
    generate_build_output_path_tmp(GOOS PATH_OUTPUT)
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
    add_custom_target(build.binary.${APP_ARCH}
        DEPENDS ${PATH_OUTPUT}
    )
endforeach()
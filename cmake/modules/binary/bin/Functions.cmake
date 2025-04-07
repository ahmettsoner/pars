get_host_os(HOST_OS)
set_os_ext(${HOST_OS} EXT)

function(map_arch_to_goarch input_arch output_goarch)
    if(${input_arch} STREQUAL ${ARCH_X86})
        set(${output_goarch} ${GO_ARCH_X86} PARENT_SCOPE)
    elseif(${input_arch} STREQUAL ${ARCH_X86_64})
        set(${output_goarch} ${GO_ARCH_X86_64} PARENT_SCOPE)
    elseif(${input_arch} STREQUAL ${ARCH_ARM})
        set(${output_goarch} ${GO_ARCH_ARM} PARENT_SCOPE)
    elseif(${input_arch} STREQUAL ${ARCH_ARM64})
        set(${output_goarch} ${GO_ARCH_ARM64} PARENT_SCOPE)
    else()
        message(FATAL_ERROR "Unsupported architecture: ${input_arch}")
    endif()
endfunction()
function(map_goarch_to_arch input_goarch output_arch)
    if(${input_goarch} STREQUAL ${GO_ARCH_X86})
        set(${output_arch} ${ARCH_X86} PARENT_SCOPE)
    elseif(${input_goarch} STREQUAL ${GO_ARCH_X86_64})
        set(${output_arch} ${ARCH_X86_64} PARENT_SCOPE)
    elseif(${input_goarch} STREQUAL ${GO_ARCH_ARM})
        set(${output_arch} ${ARCH_ARM} PARENT_SCOPE)
    elseif(${input_goarch} STREQUAL ${GO_ARCH_ARM64})
        set(${output_arch} ${ARCH_ARM64} PARENT_SCOPE)
    else()
        message(FATAL_ERROR "Unsupported Go architecture: ${input_goarch}")
    endif()
endfunction()
# map_arch_to_goarch(${ARCH_X86} GO_ARCH)
# message(STATUS "GOARCH for ${ARCH_X86}: ${GO_ARCH}")

function(build GOOS GOARCH OUTPUT_PATH)

    
    generate_build_output_path(PATH_OUTPUT)
    if("${OUTPUT_PATH}" STREQUAL "")
        set(OUTPUT_PATH ${PATH_OUTPUT})
    endif()

    map_goarch_to_arch(${GOARCH} APP_ARCH)

    set(env_vars "GOOS=${GOOS}" "GOARCH=${GOARCH}")


    if(IS_LINUX)
        set(build_mode "-buildmode=pie")
    endif()
    if (EXISTS "${CMAKE_SOURCE_DIR}/src/vendor")
        set(mod "-mod=vendor")
    endif()

    env_command_for_shell("${HOST_SHELL}" "${env_vars}" BASH_ENV_COMMAND)

    set(GO_BUILD_COMMAND "")
    list(APPEND GO_BUILD_COMMAND ${BASH_ENV_COMMAND})
    list(APPEND GO_BUILD_COMMAND "go build ${mod} ${build_mode} -ldflags=\"-X parsdevkit.net/core/utils.version=${APP_TAG} -X parsdevkit.net/core/utils.stage=${VERSION_CHANNEL} -buildid=${APP_NAME}\" -o ${OUTPUT_PATH} ./pars.go")
    
    command_for_shell("${HOST_SHELL}" "${GO_BUILD_COMMAND}" SHELL_GO_BUILD_COMMAND)


    generate_build_output_path_tmp(PATH_OUTPUT)    
    add_custom_command(
        OUTPUT ${PATH_OUTPUT}
        COMMAND ${SHELL_GO_BUILD_COMMAND}
        VERBATIM
        WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}/src
        COMMENT "Building for ${GOOS} ${APP_ARCH} with tag ${APP_TAG}..."
        )
endfunction()


function(set_goos_arch_lists GOOS)
    if(${GOOS} STREQUAL ${OS_WINDOWS})
        set(ARCH_LIST "${GOARCH_LIST_WINDOWS}" PARENT_SCOPE)
    elseif(${GOOS} STREQUAL ${OS_LINUX})
        set(ARCH_LIST "${GOARCH_LIST_LINUX}" PARENT_SCOPE)
    elseif(${GOOS} STREQUAL ${OS_MACOS})
        set(ARCH_LIST "${GOARCH_LIST_DARWIN}" PARENT_SCOPE)
    elseif(${GOOS} STREQUAL ${OS_FREEBSD})
        set(ARCH_LIST "${GOARCH_LIST_FREEBSD}" PARENT_SCOPE)
    elseif(${GOOS} STREQUAL ${OS_OPENBSD})
        set(ARCH_LIST "${GOARCH_LIST_OPENBSD}" PARENT_SCOPE)
    elseif(${GOOS} STREQUAL ${OS_NETBSD})
        set(ARCH_LIST "${GOARCH_LIST_NETBSD}" PARENT_SCOPE)
    else()
        message(FATAL_ERROR "Unsupported GOOS: ${GOOS}")
    endif()
endfunction()

function(set_build_output_from_arg path_variable)
    set(BASE_PATH "${CMAKE_SOURCE_DIR}")
    if(DEFINED OUTPUT AND NOT "${OUTPUT}" STREQUAL "")
        set(OUTPUT_PATH "${OUTPUT}")
        string(REGEX REPLACE "/$" "" OUTPUT_PATH "${OUTPUT_PATH}")
        cmake_path(IS_RELATIVE OUTPUT_PATH IS_RELATIVE_RESULT)

        if(IS_RELATIVE_RESULT)
            cmake_path(SET FINAL_PATH NORMALIZE "${BASE_PATH}/${OUTPUT_PATH}")
        else()
            set(FINAL_PATH "${OUTPUT_PATH}")
        endif()
        
        string(REGEX REPLACE "/+" "/" FINAL_PATH "${FINAL_PATH}")
        
        set(${path_variable} "${FINAL_PATH}" PARENT_SCOPE)
    endif()
endfunction()



function(generate_build_output_path path_variable)
    set(${path_variable} "${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${GOOS}/bin/${APP_ARCH}/${APP_NAME}${EXT}" PARENT_SCOPE)
endfunction()
function(generate_build_output_path_tmp path_variable)
    set(${path_variable} "${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${GOOS}/bin/${APP_ARCH}/tmp/${APP_NAME}${EXT}" PARENT_SCOPE)
endfunction()
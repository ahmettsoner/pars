

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

function(build GOOS GOARCH EXT)
    map_goarch_to_arch(${GOARCH} APP_ARCH)
    if(${HOST_SHELL} STREQUAL "powershell")
        set(GO_BUILD_ENV_COMMAND $$env:GOOS='${GOOS}' \\\\\\\\\; $$env:GOARCH='${GOARCH}' \\\\\\\\\;)
    else()
        set(GO_BUILD_ENV_COMMAND GOOS=${GOOS} GOARCH=${GOARCH})
    endif()

    set(GO_BUILD_COMMAND ${GO_BUILD_ENV_COMMAND} go build -ldflags='-X parsdevkit.net/core/utils.version=${APP_TAG} -X parsdevkit.net/core/utils.stage=final -buildid=${APP_NAME}' -o ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${GOOS}/bin/${APP_ARCH}/${APP_NAME}${EXT} ./pars.go)

    command_for_default_shell("${GO_BUILD_COMMAND}" SHELL_GO_BUILD_COMMAND)

    
    add_custom_command(
        OUTPUT ${CMAKE_SOURCE_DIR}/${DIST_ROOT_DIR}/${APP_TAG}/${GOOS}/bin/${APP_ARCH}/${APP_NAME}${EXT}
        COMMAND ${SHELL_GO_BUILD_COMMAND}
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


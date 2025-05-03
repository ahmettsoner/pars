function(map_chocoarch_to_arch input_arch output_goarch)
    if(${input_arch} STREQUAL ${CHOCO_ARCH_X86})
        set(${output_goarch} ${ARCH_X86} PARENT_SCOPE)
    elseif(${input_arch} STREQUAL ${CHOCO_ARCH_X86_64})
        set(${output_goarch} ${ARCH_X86_64} PARENT_SCOPE)
    elseif(${input_arch} STREQUAL ${CHOCO_ARCH_ARM64})
        set(${output_goarch} ${ARCH_ARM64} PARENT_SCOPE)
    else()
        message(FATAL_ERROR "Unsupported architecture: ${input_arch}")
    endif()
endfunction()




function(map_chocoarch_to_arch_all input_arch output_goarch)
    if(${input_arch} STREQUAL ${CHOCO_ARCH_ALL})
        set(${output_goarch} ${ARCH_ALL} PARENT_SCOPE)
    else()
        map_chocoarch_to_arch(${input_arch} COMING_ARCH)
        set(${output_goarch} ${COMING_ARCH} PARENT_SCOPE)
    endif()
endfunction()
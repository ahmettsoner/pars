
function(map_msiarch_to_arch input_arch output_goarch)
    if(${input_arch} STREQUAL ${MSI_ARCH_X86})
        set(${output_goarch} ${ARCH_X86} PARENT_SCOPE)
    elseif(${input_arch} STREQUAL ${MSI_ARCH_X86_64})
        set(${output_goarch} ${ARCH_X86_64} PARENT_SCOPE)
    elseif(${input_arch} STREQUAL ${MSI_ARCH_ARM64})
        set(${output_goarch} ${ARCH_ARM64} PARENT_SCOPE)
    else()
        message(FATAL_ERROR "Unsupported architecture: ${input_arch}")
    endif()
endfunction()




function(map_msiarch_to_arch_all input_arch output_goarch)
    map_msiarch_to_arch(${input_arch} COMING_ARCH)
    set(${output_goarch} ${COMING_ARCH} PARENT_SCOPE)
endfunction()



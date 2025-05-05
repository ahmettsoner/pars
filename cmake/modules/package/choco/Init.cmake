if(IS_WINDOWS)
    add_custom_command(
        OUTPUT check_env_for_choco_packing
        COMMAND ${CMAKE_COMMAND} -E echo "Windows detected. Running setup script."
    )
elseif(IS_LINUX)
    if(IS_DEBIAN)
        add_custom_command(
            OUTPUT check_env_for_choco_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux and Debian detected. Running setup script."
        )
    elseif(IS_REDHAT)
        add_custom_command(
            OUTPUT check_env_for_choco_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux and Redhat detected. Running setup script."
        )
    else()
        add_custom_command(
            OUTPUT check_env_for_choco_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux system detected, but not Debian nor Redhat."
            COMMAND exit 1
        )
    endif()
else()
    add_custom_command(
        OUTPUT check_env_for_choco_packing
        COMMAND ${CMAKE_COMMAND} -E echo "Not a Windows system. This target is applicable only for Windows Host."
        COMMAND exit 1
    )
endif()

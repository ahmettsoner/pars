if(IS_LINUX)
    if(IS_REDHAT)
        add_custom_command(
            OUTPUT check_env_for_rpm_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux and Redhat detected. Running setup script."
        )
    else()
        add_custom_command(
            OUTPUT check_env_for_rpm_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux system detected, but not Redhat."
            COMMAND exit 1
        )
    endif()
else()
    add_custom_command(
        OUTPUT check_env_for_rpm_packing
        COMMAND ${CMAKE_COMMAND} -E echo "Not a Linux system. This target is applicable only for Linux/Redhat Host."
        COMMAND exit 1
    )
endif()

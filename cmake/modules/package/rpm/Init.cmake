if(IS_LINUX)
<<<<<<< HEAD
    if(IS_REDHAT)
        add_custom_command(
            OUTPUT check_env_for_rpm_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux and Redhat detected. Running setup script."
=======
if(IS_REDHAT)
    add_custom_command(
        OUTPUT check_env_for_rpm_packing
        COMMAND ${CMAKE_COMMAND} -E echo "Linux and Redhat detected. Running setup script."
    )
    elseif(IS_DEBIAN)
        add_custom_command(
            OUTPUT check_env_for_rpm_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux and Debian detected. Running setup script."
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
        )
    else()
        add_custom_command(
            OUTPUT check_env_for_rpm_packing
<<<<<<< HEAD
            COMMAND ${CMAKE_COMMAND} -E echo "Linux system detected, but not Redhat."
=======
            COMMAND ${CMAKE_COMMAND} -E echo "Linux system detected, but not Redhat nor Debian."
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
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

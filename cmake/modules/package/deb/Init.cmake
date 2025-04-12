if(IS_LINUX)
    if(IS_DEBIAN)
        add_custom_command(
            OUTPUT check_env_for_deb_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux and Debian detected. Running setup script."
        )
<<<<<<< HEAD
    else()
        add_custom_command(
            OUTPUT check_env_for_deb_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux system detected, but not Debian."
=======
    elseif(IS_REDHAT)
        add_custom_command(
            OUTPUT check_env_for_deb_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux and ReadHat detected. Running setup script."
        )
    else()
        add_custom_command(
            OUTPUT check_env_for_deb_packing
            COMMAND ${CMAKE_COMMAND} -E echo "Linux system detected, but not Debian nor ReadHat."
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
            COMMAND exit 1
        )
    endif()
else()
    add_custom_command(
        OUTPUT check_env_for_deb_packing
        COMMAND ${CMAKE_COMMAND} -E echo "Not a Linux system. This target is applicable only for Linux/Debian Host."
        COMMAND exit 1
    )
endif()

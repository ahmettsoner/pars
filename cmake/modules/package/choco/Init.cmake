if(IS_WINDOWS)
    add_custom_command(
        OUTPUT check_env_for_choco_packing
        COMMAND ${CMAKE_COMMAND} -E echo "Windows detected. Running setup script."
    )
else()
    add_custom_command(
        OUTPUT check_env_for_choco_packing
        COMMAND ${CMAKE_COMMAND} -E echo "Not a Windows system. This target is applicable only for Windows Host."
        COMMAND exit 1
    )
endif()

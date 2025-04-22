get_host_os(HOST_OS)
add_custom_command(
    OUTPUT check_env_for_brew_packing
    COMMAND ${CMAKE_COMMAND} -E echo "${HOST_OS} detected. Running setup script."
)

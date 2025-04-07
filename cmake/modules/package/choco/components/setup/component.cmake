set(PACKAGES
    cmake
    make
)
set(COMMANDS
)

command_for_shell("powershell" "${COMMANDS}" SHELL_GO_BUILD_COMMAND)


add_custom_command(
    OUTPUT ./choco-setup
    COMMAND ${CMAKE_COMMAND} -E echo "Setting up the host machine for package build..."
    COMMAND ${CMAKE_COMMAND} -E echo "Running additional setup commands..."
    # COMMAND ${SHELL_GO_BUILD_COMMAND}
    VERBATIM
)

add_custom_target(build.choco.package.setup

    DEPENDS check_env_for_choco_packing ./choco-setup
)

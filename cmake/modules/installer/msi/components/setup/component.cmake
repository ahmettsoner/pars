set(PACKAGES
    gnome-keyring
    lxd
    cmake
    make
    golang-any
)
set(COMMANDS
)

command_for_shell("bash" "${COMMANDS}" SHELL_GO_BUILD_COMMAND)


add_custom_command(
    OUTPUT ./msi-setup
    COMMAND ${CMAKE_COMMAND} -E echo "Setting up the host machine for installer build..."

    COMMAND ${CMAKE_COMMAND} -E echo "Running additional setup commands..."

    COMMAND ${SHELL_GO_BUILD_COMMAND}

)

add_custom_target(build.msi.package.setup

    DEPENDS check_env_for_msi_packing ./msi-setup
)

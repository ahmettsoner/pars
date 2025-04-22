set(PACKAGES
    pkgbuild
    productbuild
)
set(COMMANDS
)

command_for_shell("bash" "${COMMANDS}" SHELL_GO_BUILD_COMMAND)


add_custom_command(
    OUTPUT ./pkg-setup
    COMMAND ${CMAKE_COMMAND} -E echo "Setting up the host machine for installer build..."

    COMMAND ${CMAKE_COMMAND} -E echo "Running additional setup commands..."

    COMMAND ${SHELL_GO_BUILD_COMMAND}

)

add_custom_target(build.pkg.package.setup

    DEPENDS check_env_for_pkg_packing ./pkg-setup
)

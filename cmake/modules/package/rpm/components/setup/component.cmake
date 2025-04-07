set(PACKAGES
    rpm-build
    rpmdevtools
    cmake
    make
)
set(COMMANDS
)

command_for_shell("bash" "${COMMANDS}" SHELL_GO_BUILD_COMMAND)


add_custom_command(
    OUTPUT ./rpm-setup
    COMMAND ${CMAKE_COMMAND} -E echo "Setting up the host machine for package build..."
    COMMAND sudo dnf install -y ${PACKAGES}
    COMMAND ${CMAKE_COMMAND} -E echo "Running additional setup commands..."
    COMMAND ${CMAKE_COMMAND} -E cmake_echo_color --cyan "Installing Snapcraft and initializing LXD..."
    COMMAND ${SHELL_GO_BUILD_COMMAND}
    VERBATIM
)

add_custom_target(build.rpm.package.setup

    DEPENDS check_env_for_rpm_packing ./rpm-setup
)

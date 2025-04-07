set(PACKAGES
    snapd
    gnome-keyring
    lxd
    cmake
    make
    golang-any
)
set(COMMANDS
    "sudo snap install snapcraft --classic"
    "sudo lxd init"
)

command_for_shell("bash" "${COMMANDS}" SHELL_GO_BUILD_COMMAND)


add_custom_command(
    OUTPUT ./snap-setup
    COMMAND ${CMAKE_COMMAND} -E echo "Setting up the host machine for package build..."
    COMMAND sudo apt-get update && sudo apt-get install -y ${PACKAGES}
    COMMAND ${CMAKE_COMMAND} -E echo "Running additional setup commands..."
    COMMAND ${CMAKE_COMMAND} -E cmake_echo_color --cyan "Installing Snapcraft and initializing LXD..."
    COMMAND ${SHELL_GO_BUILD_COMMAND}
    VERBATIM
)

add_custom_target(build.snap.package.setup

    DEPENDS check_env_for_snap_packing ./snap-setup
)

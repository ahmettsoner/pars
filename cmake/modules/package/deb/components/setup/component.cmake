if(IS_REDHAT)
    set(PACKAGES
        rpm-build
        rpmdevtools
        dpkg-dev
        # Add other necessary packages that mimic your build-essential needs
        cmake
        golang
    )

    add_custom_command(
        OUTPUT ./deb-setup
        COMMAND ${CMAKE_COMMAND} -E echo "Setting up the host machine for package build on Red Hat-based system..."
        COMMAND sudo dnf install -y ${PACKAGES}
        COMMAND ${CMAKE_COMMAND} -E echo "Running additional setup commands not fully supporting native deb packaging..."
        VERBATIM
    )
elseif(IS_DEBIAN)
    set(PACKAGES
        build-essential
        devscripts
        dh-make
        dh-golang
        debhelper
        lintian
        fakeroot
        cmake
        golang-any
    )

    add_custom_command(
        OUTPUT ./deb-setup
        COMMAND ${CMAKE_COMMAND} -E echo "Setting up the host machine for package build..."
        COMMAND sudo apt-get update && sudo apt-get install -y ${PACKAGES}
        COMMAND ${CMAKE_COMMAND} -E echo "Running additional setup commands..."
        COMMAND ${CMAKE_COMMAND} -E cmake_echo_color --cyan "Installing Snapcraft and initializing LXD..."
        VERBATIM
    )
endif()

add_custom_target(build.deb.package.setup
    DEPENDS check_env_for_deb_packing ./deb-setup
)

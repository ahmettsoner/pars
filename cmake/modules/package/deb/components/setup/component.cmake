if(IS_DEBIAN)
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
        VERBATIM
    )
endif()

add_custom_target(build.deb.package.setup
    DEPENDS check_env_for_deb_packing ./deb-setup
)

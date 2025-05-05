add_custom_command(
    OUTPUT info_print_version
    COMMAND ${CMAKE_COMMAND} -E echo "APP_TAG ${APP_TAG}"
)
add_custom_target(info.print.version DEPENDS info_print_version)

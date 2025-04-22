add_custom_command(
    OUTPUT check_env_for_pkg_packing
    COMMAND ${CMAKE_COMMAND} -E echo "Macos detected. Running setup script."
)
# if(IS_MACOS)
#     add_custom_command(
#         OUTPUT check_env_for_pkg_packing
#         COMMAND ${CMAKE_COMMAND} -E echo "Macos detected. Running setup script."
#     )
# else()
#     add_custom_command(
#         OUTPUT check_env_for_pkg_packing
#         COMMAND ${CMAKE_COMMAND} -E echo "Not a Macos system. This target is applicable only for Macos Host."
#         COMMAND exit 1
#     )
# endif()

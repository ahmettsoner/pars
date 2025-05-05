add_subdirectory(core)
add_subdirectory(config)
add_subdirectory(common)
add_subdirectory(utils)
add_subdirectory(modules)

# include("${CMAKE_SOURCE_DIR}/cmake/utils/archive/Consts.cmake")
# compress_folder(${ARCHIVE_TAR_GZ_EXT} "/home/ahmetsoner/pars-pro/pars/ttt/com/t" "/home/ahmetsoner/pars-pro/pars/ttt/out" "t${ARCHIVE_TAR_GZ_EXT}")
# add_custom_target(run_compress ALL DEPENDS "/home/ahmetsoner/pars-pro/pars/ttt/out/t${ARCHIVE_TAR_GZ_EXT}")

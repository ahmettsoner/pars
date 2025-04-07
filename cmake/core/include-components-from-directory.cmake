function(include_all_components_from_directory)
    set(module_dir "${CMAKE_CURRENT_LIST_DIR}")

    get_filename_component(module_name ${module_dir} NAME)

    file(GLOB_RECURSE component_files "${module_dir}/components/*/component.cmake")
    message(STATUS "   |- Loading components")

    foreach(component_file ${component_files})
        get_filename_component(component_dir ${component_file} DIRECTORY)
        get_filename_component(component_name ${component_dir} NAME)

        message(STATUS "      |- Including '${component_name}'")
        include(${component_file})
    endforeach()
endfunction()

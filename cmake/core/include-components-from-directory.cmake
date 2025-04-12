function(include_all_components_from_directory)
    set(module_dir "${CMAKE_CURRENT_LIST_DIR}")

    get_filename_component(module_name ${module_dir} NAME)

    file(GLOB_RECURSE component_files "${module_dir}/components/*/component.cmake")
<<<<<<< HEAD
    message(STATUS "   |- Loading components")
=======
    # message(STATUS "   |- Loading components")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c

    foreach(component_file ${component_files})
        get_filename_component(component_dir ${component_file} DIRECTORY)
        get_filename_component(component_name ${component_dir} NAME)

<<<<<<< HEAD
        message(STATUS "      |- Including '${component_name}'")
=======
        # message(STATUS "      |- Including '${component_name}'")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
        include(${component_file})
    endforeach()
endfunction()

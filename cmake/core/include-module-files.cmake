function(include_module_files)
    set(module_dir "${CMAKE_CURRENT_LIST_DIR}")

    get_filename_component(module_name ${module_dir} NAME)

    set(consts_file "${module_dir}/Consts.cmake")
    if(EXISTS ${consts_file})
<<<<<<< HEAD
        message(STATUS "   |- Including Consts")
=======
        # message(STATUS "   |- Including Consts")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
        include(${consts_file})
    endif()

    set(init_file "${module_dir}/Init.cmake")
    if(EXISTS ${init_file})
<<<<<<< HEAD
        message(STATUS "   |- Including Init")
=======
        # message(STATUS "   |- Including Init")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
        include(${init_file})
    endif()

    set(functions_file "${module_dir}/Functions.cmake")
    if(EXISTS ${functions_file})
<<<<<<< HEAD
        message(STATUS "   |- Including Functions")
=======
        # message(STATUS "   |- Including Functions")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
        include(${functions_file})
    endif()

    set(module_file "${module_dir}/Module.cmake")
    if(EXISTS ${module_file})
<<<<<<< HEAD
        message(STATUS "   |- Including Main")
=======
        # message(STATUS "   |- Including Main")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
        include(${module_file})
    else()
        message(FATAL_ERROR "Error: ${module_file} not found. Every module must have a Module.cmake file.")
    endif()
endfunction()


function(include_module_files)
    set(module_dir "${CMAKE_CURRENT_LIST_DIR}")

    get_filename_component(module_name ${module_dir} NAME)

    set(consts_file "${module_dir}/Consts.cmake")
    if(EXISTS ${consts_file})
        message(STATUS "   |- Including Consts")
        include(${consts_file})
    endif()

    set(init_file "${module_dir}/Init.cmake")
    if(EXISTS ${init_file})
        message(STATUS "   |- Including Init")
        include(${init_file})
    endif()

    set(functions_file "${module_dir}/Functions.cmake")
    if(EXISTS ${functions_file})
        message(STATUS "   |- Including Functions")
        include(${functions_file})
    endif()

    set(module_file "${module_dir}/Module.cmake")
    if(EXISTS ${module_file})
        message(STATUS "   |- Including Main")
        include(${module_file})
    else()
        message(FATAL_ERROR "Error: ${module_file} not found. Every module must have a Module.cmake file.")
    endif()
endfunction()


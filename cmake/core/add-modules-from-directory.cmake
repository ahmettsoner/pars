function(add_modules_from_directory)
    file(GLOB directories RELATIVE ${CMAKE_CURRENT_LIST_DIR} ${CMAKE_CURRENT_LIST_DIR}/*)

    foreach(directory ${directories})
        if(IS_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}/${directory})
<<<<<<< HEAD
            message(STATUS "Loading module: ${directory}")
=======
            # message(STATUS "Loading module: ${directory}")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
            add_module(${CMAKE_CURRENT_LIST_DIR}/${directory})
        endif()
    endforeach()
endfunction()

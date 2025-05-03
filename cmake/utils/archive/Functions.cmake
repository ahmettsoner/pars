
function(compress_folder format base_path output_path output_filename)
    include("${CMAKE_SOURCE_DIR}/cmake/utils/archive/Consts.cmake")
    include("${CMAKE_SOURCE_DIR}/cmake/core/shell.cmake")

    if(${format} STREQUAL ${ARCHIVE_TAR_GZ_EXT})
        set(COMPRESS_COMMAND "tar -czf ${output_path}/${output_filename} -C ${base_path} .")
    elseif(${format} STREQUAL ${ARCHIVE_TAR_BZ2_EXT})
        set(COMPRESS_COMMAND "tar -cjf ${output_path}/${output_filename} -C ${base_path} .")
    elseif(${format} STREQUAL ${ARCHIVE_TAR_XZ_EXT})
        set(COMPRESS_COMMAND "tar -cJf ${output_path}/${output_filename} -C ${base_path} .")
    elseif(${format} STREQUAL ${ARCHIVE_ZIP_EXT})
        if(${HOST_SHELL} STREQUAL "powershell")
            set(COMPRESS_COMMAND "Compress-Archive -Path ${base_path} -DestinationPath ${output_path}/${output_filename}")
        else()
            set(COMPRESS_COMMAND "zip -r ${output_path}/${output_filename} .")
        endif()
    elseif(${format} STREQUAL ${ARCHIVE_RAR_EXT})
        if(${HOST_SHELL} STREQUAL "powershell")
            set(COMPRESS_COMMAND "rar a ${output_path}/${output_filename} ${base_path}")
        else()
            set(COMPRESS_COMMAND "rar a ${output_path}/${output_filename} .")
        endif()
    elseif(${format} STREQUAL ${ARCHIVE_SEVEN_Z_EXT})
        if(${HOST_SHELL} STREQUAL "powershell")
            set(COMPRESS_COMMAND "7z a ${output_path}/${output_filename} ${base_path}")
        else()
            set(COMPRESS_COMMAND "7z a ${output_path}/${output_filename} .")
        endif()
    elseif(${format} STREQUAL ${ARCHIVE_LZ_EXT})
        set(COMPRESS_COMMAND "tar --lzma -cf ${output_path}/${output_filename} -C ${base_path} .")
    elseif(${format} STREQUAL ${ARCHIVE_ZST_EXT})
        set(COMPRESS_COMMAND "tar --zstd -cf ${output_path}/${output_filename} -C ${base_path} .")
    else()
        message(FATAL_ERROR "Unsupported format: ${format}")
    endif()

    command_for_default_shell(${COMPRESS_COMMAND} SHELL_COMPRESS_COMMAND)

    add_custom_command(
        OUTPUT ${output_path}/${output_filename}
        COMMAND "${SHELL_COMPRESS_COMMAND}"
        WORKING_DIRECTORY ${base_path}
        COMMENT "Compressing ${base_path} to ${output_path} in format ${format}..."
    )
endfunction()
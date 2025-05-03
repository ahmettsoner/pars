set(SHELL "" CACHE STRING "Custom shell to be used (e.g. bash, zsh, fish, powershell, cmd)")

function(get_shell detected_shell)
    if(SHELL)
        # "message(STATUS" "Using custom shell: ${SHELL}")
        set(${detected_shell} ${SHELL} PARENT_SCOPE)
    else()
        if(WIN32)
            # Check for PowerShell Core (pwsh) first
            execute_process(
                COMMAND pwsh -Command "echo 'pwsh detected'"
                RESULT_VARIABLE PW_SH_DETECTED
                OUTPUT_QUIET ERROR_QUIET
            )

            if(PW_SH_DETECTED EQUAL 0)
                # message(STATUS "Detected PowerShell Core (pwsh)")
                set(${detected_shell} "pwsh" PARENT_SCOPE)
            else()
                # Check for regular PowerShell (powershell)
                execute_process(
                    COMMAND powershell -Command "echo 'powershell detected'"
                    RESULT_VARIABLE PS_DETECTED
                    OUTPUT_QUIET ERROR_QUIET
                )

                if(PS_DETECTED EQUAL 0)
                    # message(STATUS "Detected Windows PowerShell")
                    set(${detected_shell} "powershell" PARENT_SCOPE)
                else()
                    # If neither PowerShell Core nor regular PowerShell is found, fallback to cmd
                    message(WARNING "Neither PowerShell Core (pwsh) nor PowerShell found. Defaulting to cmd.")
                    set(${detected_shell} "cmd" PARENT_SCOPE)
                endif()
            endif()
        else()
            execute_process(
                COMMAND bash -c "echo \$SHELL"
                OUTPUT_VARIABLE DEFAULT_SHELL
                OUTPUT_STRIP_TRAILING_WHITESPACE
            )

            if (NOT DEFAULT_SHELL)
                message(WARNING "Could not detect shell. Using bash as default.")
                set(${detected_shell} "bash" PARENT_SCOPE)
            else()
                # message(STATUS "Detected default shell: ${DEFAULT_SHELL}")

                if("${DEFAULT_SHELL}" MATCHES "bash")
                    set(${detected_shell} "bash" PARENT_SCOPE)
                elseif("${DEFAULT_SHELL}" MATCHES "zsh")
                    set(${detected_shell} "zsh" PARENT_SCOPE)
                elseif("${DEFAULT_SHELL}" MATCHES "fish")
                    set(${detected_shell} "fish" PARENT_SCOPE)
                else()
                    message(WARNING "Unknown shell detected: ${DEFAULT_SHELL}, using bash as fallback.")
                    set(${detected_shell} "bash" PARENT_SCOPE)
                endif()
            endif()
        endif()
    endif()
endfunction()
get_shell(HOST_SHELL)


function(env_command_for_shell shell env_vars output_variable_name)
    get_filename_component(shell_name ${shell} NAME)

    # Set prefix based on shell type
    if(${shell_name} STREQUAL "powershell" OR ${shell_name} STREQUAL "pwsh")
        set(prefix "$env:")
        set(separator " ; ")  # Use ; as separator for PowerShell
    elseif(${shell_name} STREQUAL "cmd")
        set(prefix "set ")
        set(separator " & ")  # Use & as separator for CMD
    elseif(${shell_name} STREQUAL "bash" OR ${shell_name} STREQUAL "zsh" OR ${shell_name} STREQUAL "fish")
        set(prefix "export ")
        set(separator " && ")  # Use && as separator for bash-like shells
    else()
        message(FATAL_ERROR "Unsupported shell: ${shell_name}")
    endif()

    # Prepare environment variable commands
    set(env_commands "")
    foreach(env_var IN LISTS env_vars)
        if(${shell_name} STREQUAL "powershell" OR ${shell_name} STREQUAL "pwsh")
            string(REPLACE "=" "=\"" env_var_quoted ${env_var})
            set(env_var_quoted "${env_var_quoted}\"")
            list(APPEND env_commands "${prefix}${env_var_quoted}")
            elseif(${shell_name} STREQUAL "cmd")
            list(APPEND env_commands "${prefix}${env_var}")
        else()  # For bash, zsh, fish
            list(APPEND env_commands "${prefix}${env_var}")
        endif()
    endforeach()

    # Join all commands with the appropriate separator
    string(JOIN "${separator}" final_command ${env_commands})

    # Set the joined command string in the output variable
    set(${output_variable_name} "${final_command}" PARENT_SCOPE)
endfunction()


function(command_for_shell shell command_list output_variable_name)
    get_filename_component(shell_name ${shell} NAME)

    
    if(${shell_name} STREQUAL "powershell" OR ${shell_name} STREQUAL "pwsh")
        set(separator " \\\\\\; ")
    elseif(${shell_name} STREQUAL "cmd")
        set(separator " & ")
    else()
        set(separator " && ")
    endif()

    string(JOIN "${separator}" command_to_run ${command_list})

    # if(${shell_name} STREQUAL "${HOST_SHELL}") 
    #     set(shell_command ${command_to_run})
    if(${shell_name} STREQUAL "cmd")
        set(shell_command cmd /c ${command_to_run})
    elseif(${shell_name} STREQUAL "powershell")
        set(shell_command powershell -ExecutionPolicy Bypass -Command ${command_to_run})
    elseif(${shell_name} STREQUAL "pwsh")
        set(shell_command pwsh -Command ${command_to_run})
    elseif(${shell_name} STREQUAL "bash")
        set(shell_command bash -c ${command_to_run})
    elseif(${shell_name} STREQUAL "zsh")
        set(shell_command zsh -c ${command_to_run})
    elseif(${shell_name} STREQUAL "fish")
        set(shell_command fish -c ${command_to_run})
    else()
        message(FATAL_ERROR "Unsupported shell: ${shell_name}")
    endif()

    set(${output_variable_name} ${shell_command} PARENT_SCOPE)
endfunction()

function(command_for_default_shell command_list output_variable_name)
    command_for_shell(${HOST_SHELL} "${command_list}" shell_command)

    set(${output_variable_name} ${shell_command} PARENT_SCOPE)

endfunction()

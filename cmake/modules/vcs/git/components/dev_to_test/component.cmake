# set(MERGE_COMMAND "
#     git config user.name \"CI/CD Bot\"
#     git config user.email \"ci-bot@yourdomain.com\"

#     # Check if 'test' branch exists
#     if git ls-remote --heads origin test | grep \"refs/heads/test\"; then
#       # If exists, fetch it and merge dev into it
#       git fetch origin test:test
#       git checkout test
#       git merge --allow-unrelated-histories --strategy-option theirs dev || true
#     else
#       # If not exists, create new 'test' branch from dev
#       git checkout -b test
#     fi

#     # git push origin test
# ")

set(MERGE_COMMAND "")
list(APPEND MERGE_COMMAND "git config user.name \"${GIT_USER_NAME_CI_CD_BOT}\"")
list(APPEND MERGE_COMMAND "git config user.email \"${GIT_USER_EMAIL_CI_CD_BOT}\"")

get_shell(HOST_SHELL)
# if("${HOST_SHELL}" STREQUAL "powershell" OR "${HOST_SHELL}" STREQUAL "pwsh" )
#     set(build_mode "if (git ls-remote --heads origin test | Select-String \"refs/heads/test\")")
# endif()


list(APPEND MERGE_COMMAND "git checkout -b test")
list(APPEND MERGE_COMMAND "git merge dev")

command_for_shell("${HOST_SHELL}" "${MERGE_COMMAND}" SHELL_MERGE_COMMAND)
add_custom_command(
    OUTPUT merge_dev_into_test
    COMMAND ${SHELL_MERGE_COMMAND}
    VERBATIM
)

add_custom_target(vcs.git.move.dev-to-test DEPENDS merge_dev_into_test)

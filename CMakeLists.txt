cmake_minimum_required(VERSION 3.22)
project(Pars)

find_program(GO_EXEC go PATHS /usr/local/bin /usr/bin)
if(NOT GO_EXEC)
    message(FATAL_ERROR "Go executable not found!")
endif()

find_program(GIT_EXEC git)
if(NOT GIT_EXEC)
    message(FATAL_ERROR "Git executable not found!")
endif()

find_program(GRM_EXEC grm)
if(NOT GRM_EXEC)
    message(FATAL_ERROR "grm executable not found!")
endif()

include("${CMAKE_SOURCE_DIR}/cmake/core/version.cmake")
if(NOT DEFINED VERSION OR VERSION STREQUAL "")
    set(VERSION "NONE")
endif()
generate_version_info(${VERSION} APP_TAG)


string(REGEX MATCH "^([A-Za-z]*)([0-9]+\\.[0-9]+\\.[0-9]+)([-]([a-zA-Z]+)([\.]?([0-9]+))?)?$" VERSION_MATCH ${APP_TAG})
if(VERSION_MATCH)
    set(VERSION_PREFIX ${CMAKE_MATCH_1})
    set(VERSION_SEMVER ${CMAKE_MATCH_2})
    set(VERSION_CHANNEL ${CMAKE_MATCH_4})
    set(VERSION_RELEASE ${CMAKE_MATCH_6})
else()
    message(FATAL_ERROR "Invalid version format: ${APP_TAG}")
endif()

string(REGEX MATCH "([0-9]+)\\.([0-9]+)\\.([0-9]+)" SEMVER_MATCH ${VERSION_SEMVER})
if(SEMVER_MATCH)
    set(VERSION_MAJOR ${CMAKE_MATCH_1})
    set(VERSION_MINOR ${CMAKE_MATCH_2})
    set(VERSION_PATCH ${CMAKE_MATCH_3})
else()
    message(FATAL_ERROR "Invalid semver format: ${VERSION_SEMVER}")
endif()

set(RAW_VERSION "${VERSION_SEMVER}")

set(PRE_RELEASE_CHANNELS "none" "dev" "test" "alpha" "beta" "preview")
if(DEFINED VERSION_CHANNEL AND NOT "${VERSION_CHANNEL}" STREQUAL "")
    list(FIND PRE_RELEASE_CHANNELS "${VERSION_CHANNEL}" CHANNEL_INDEX)
    if(CHANNEL_INDEX EQUAL -1)
        message(FATAL_ERROR "Invalid version channel: ${VERSION_CHANNEL}. Allowed values are ${PRE_RELEASE_CHANNELS}.")
    endif()
    
    set(RAW_VERSION "${RAW_VERSION}-${VERSION_CHANNEL}")
else()
    set(VERSION_CHANNEL "stabil")
endif()

if(DEFINED VERSION_RELEASE AND NOT "${VERSION_RELEASE}" STREQUAL "")
    set(RAW_VERSION "${RAW_VERSION}.${VERSION_RELEASE}")
endif()

if(NOT DEFINED RELEASE_DATE OR RELEASE_DATE STREQUAL "")

    execute_process(
        COMMAND git log -1 --format=%cd --date=short
        OUTPUT_VARIABLE GIT_COMMIT_DATE
        OUTPUT_STRIP_TRAILING_WHITESPACE
    )
    set(GIT_DATE ${GIT_COMMIT_DATE})

    set(RELEASE_DATE ${GIT_DATE})

endif()

set(IS_PRERELEASE OFF)
if(NOT VERSION_CHANNEL STREQUAL "stabil")
    set(IS_PRERELEASE ON)
endif()

set(VERSION_SEMVER_BUILD ${VERSION_SEMVER})
if(IS_PRERELEASE)
    list(FIND PRE_RELEASE_CHANNELS "${VERSION_CHANNEL}" CHANNEL_INDEX)
    if(CHANNEL_INDEX EQUAL -1)
        message(FATAL_ERROR "Invalid version channel: ${VERSION_CHANNEL}. Allowed values are ${PRE_RELEASE_CHANNELS}.")
    endif()
    set(VERSION_SEMVER_BUILD "${VERSION_SEMVER}.${CHANNEL_INDEX}${VERSION_RELEASE}")
endif()

message(STATUS "APP_TAG: ${APP_TAG}")
message(STATUS "VERSION_PREFIX: ${VERSION_PREFIX}")
message(STATUS "VERSION_SEMVER: ${VERSION_SEMVER}")
message(STATUS "VERSION_CHANNEL: ${VERSION_CHANNEL}")
message(STATUS "VERSION_RELEASE: ${VERSION_RELEASE}")
message(STATUS "IS_PRERELEASE: ${IS_PRERELEASE}")
message(STATUS "VERSION_SEMVER_BUILD: ${VERSION_SEMVER_BUILD}")


add_subdirectory(cmake)
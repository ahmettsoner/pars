set(PROJECT_NAME "Pars")
set(PROJECT_ORGANIZATION "Pars Dev Kit")
<<<<<<< HEAD
set(PROJECT_MAINTANER "Pars Dev Kit <parsdevkit@gmail.com>")
=======
set(PROJECT_MAINTAINER "Pars Dev Kit <parsdevkit@gmail.com>")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
set(PROJECT_OWNER "Ahmet Soner <ahmettsoner@gmail.com>")
set(PROJECT_HOMEPAGE "https://parsdevkit.net")
set(PROJECT_ICON_URL "https://parsdevkit.net")
set(PROJECT_GIT "https://github.com/parsdevkit/pars")
set(PROJECT_LICENCE_TYPE "Apache-2.0")
set(PROJECT_LICENCE_URL "https://github.com/parsdevkit/pars/blob/main/LICENSE")
set(PROJECT_SUMMARY "${PROJECT_NAME} is a simple utility.")
set(PROJECT_DESCRIPTION "${PROJECT_NAME} is a simple utility.")
set(PROJECT_TAGS "code, generate, boilerplate")

set(DIST_ROOT_DIR dist)

set(PROJECT_DISPLAY_NAME ${PROJECT_NAME})
set(PROJECT_DISPLAY_NAME_RELEASETYPE ${PROJECT_DISPLAY_NAME})
if(IS_PRERELEASE)
    set(PROJECT_DISPLAY_NAME_RELEASETYPE "${PROJECT_DISPLAY_NAME} - ${VERSION_CHANNEL}.${VERSION_RELEASE}")
endif()
<<<<<<< HEAD
message(STATUS "PROJECT_DISPLAY_NAME_RELEASETYPE: ${PROJECT_DISPLAY_NAME_RELEASETYPE}")
=======
# message(STATUS "PROJECT_DISPLAY_NAME_RELEASETYPE: ${PROJECT_DISPLAY_NAME_RELEASETYPE}")
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c

file(WRITE ${CONFIG_FILE_PATH} "Name: ${APP_NAME}\n")
file(APPEND ${CONFIG_FILE_PATH} "Version: ${VERSION_SEMVER}\n")
file(APPEND ${CONFIG_FILE_PATH} "Release: ${VERSION_CHANNEL}.${VERSION_RELEASE}%{?dist}\n")
file(APPEND ${CONFIG_FILE_PATH} "Summary: ${PROJECT_SUMMARY}\n")
file(APPEND ${CONFIG_FILE_PATH} "License: ${PROJECT_LICENCE_TYPE}\n")
file(APPEND ${CONFIG_FILE_PATH} "URL: ${PROJECT_HOMEPAGE}\n")
file(APPEND ${CONFIG_FILE_PATH} "Source0: %{name}-%{version}.tar.gz\n")
if(NOT ${RPMARCH} STREQUAL ${RPM_ARCH_ALL})
    file(APPEND ${CONFIG_FILE_PATH} "BuildArch: ${RPMARCH}\n")
endif()
file(APPEND ${CONFIG_FILE_PATH} "BuildRequires: git, make, cmake, golang\n")
file(APPEND ${CONFIG_FILE_PATH} "Requires: glibc\n")

file(APPEND ${CONFIG_FILE_PATH} "\n%description\n")
file(APPEND ${CONFIG_FILE_PATH} "${PROJECT_DESCRIPTION}\n")

file(APPEND ${CONFIG_FILE_PATH} "\n%prep\n")
file(APPEND ${CONFIG_FILE_PATH} "tar -xzf %{SOURCE0} -C %{_builddir}\n")

file(APPEND ${CONFIG_FILE_PATH} "\n%build\n")
file(APPEND ${CONFIG_FILE_PATH} "make build.cmake.linux VERSION=${APP_TAG}\n")
file(APPEND ${CONFIG_FILE_PATH} "make build.binary OUTPUT=${LINUX_APP_BINARY_DIR}/${APP_NAME}\n")
file(APPEND ${CONFIG_FILE_PATH} "echo %{buildroot}\n")

file(APPEND ${CONFIG_FILE_PATH} "\n%install\n")
# Uncomment if specific build commands are needed
# file(APPEND ${CONFIG_FILE_PATH} "${MAKE} package.rpm.move-binary-to-package-source TAG=${APP_TAG} OS=${OS_LINUX} ARCH=${ARCH_FLAG_VALUE}\n")
file(APPEND ${CONFIG_FILE_PATH} "mkdir -p %{buildroot}/${LINUX_APP_BINARY_DIR}\n")
file(APPEND ${CONFIG_FILE_PATH} "mkdir -p %{buildroot}/${LINUX_APP_CONFIG_DIR}\n")
file(APPEND ${CONFIG_FILE_PATH} "mkdir -p %{buildroot}/${LINUX_APP_LOG_DIR}\n")
file(APPEND ${CONFIG_FILE_PATH} "mkdir -p %{buildroot}/${LINUX_APP_DATA_DATABASE_DIR}\n")
file(APPEND ${CONFIG_FILE_PATH} "mkdir -p %{buildroot}/${LINUX_APP_CACHE_DIR}\n")
file(APPEND ${CONFIG_FILE_PATH} "mkdir -p %{buildroot}/${LINUX_APP_LIB_DIR}\n")
file(APPEND ${CONFIG_FILE_PATH} "mkdir -p %{buildroot}/${LINUX_APP_SHARE_DIR}\n")
file(APPEND ${CONFIG_FILE_PATH} "mkdir -p %{buildroot}/${LINUX_APP_DOCS_DIR}\n")
file(APPEND ${CONFIG_FILE_PATH} "cp -r ${LINUX_APP_BINARY_DIR}/${APP_NAME} %{buildroot}/${LINUX_APP_BINARY_DIR}\n")

file(APPEND ${CONFIG_FILE_PATH} "\n%files\n")
file(APPEND ${CONFIG_FILE_PATH} "%{_bindir}/${APP_NAME}\n")
# file(APPEND ${CONFIG_FILE_PATH} "/${LINUX_APP_DOCS_DIR}/*\n")

file(APPEND ${CONFIG_FILE_PATH} "\n%changelog\n")
file(APPEND ${CONFIG_FILE_PATH} "* ${RELEASE_DATE_RPM} ${PROJECT_MAINTANER} - ${VERSION_SEMVER}\n")

if(EXISTS ${CHANGELOG_PATH})
    file(READ ${CHANGELOG_PATH} CHANGELOG_CONTENT)

    string(REGEX REPLACE "^\*" "-" PROCESSED_CHANGELOG_CONTENT "${CHANGELOG_CONTENT}")

    file(APPEND ${CONFIG_FILE_PATH} "${PROCESSED_CHANGELOG_CONTENT}\n")
else()
    file(APPEND ${CONFIG_FILE_PATH} "- Not specified any changes\n")
endif()

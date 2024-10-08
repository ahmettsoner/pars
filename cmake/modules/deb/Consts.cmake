set(DEB_ARCH_X86 "i386")
set(DEB_ARCH_X86_64 "amd64")
set(DEB_ARCH_ARM "armhf")
set(DEB_ARCH_ARM64 "arm64")


set(DEB_SERIES noble)
set(DEB_PACKAGE_EXT .deb)
set(DEB_PACK_TYPE "source")


set(DEBARCH_LIST_LINUX "${DEB_ARCH_X86};${DEB_ARCH_X86_64};${DEB_ARCH_ARM};${DEB_ARCH_ARM64}")



execute_process(
    COMMAND bash -c "date -d '${RELEASE_DATE}' '+%a, %d %b %Y 00:00:00 +0000'"
    OUTPUT_VARIABLE RELEASE_DATE_DEB
    OUTPUT_STRIP_TRAILING_WHITESPACE
)
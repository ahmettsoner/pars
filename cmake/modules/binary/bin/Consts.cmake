set(GO_ARCH_X86 "386")
set(GO_ARCH_X86_64 "amd64")
set(GO_ARCH_ARM "arm")
set(GO_ARCH_ARM64 "arm64")


set(GOOS_LIST 
    ${OS_LINUX}
    ${OS_WINDOWS}
    ${OS_MACOS}
    ${OS_FREEBSD}
    ${OS_NETBSD}
    ${OS_OPENBSD}
    )

set(GOARCH_LIST_WINDOWS "${GO_ARCH_X86};${GO_ARCH_X86_64};${GO_ARCH_ARM};${GO_ARCH_ARM64}")
# set(GOARCH_LIST_WINDOWS "${GO_ARCH_X86_64};${GO_ARCH_ARM64}")
set(GOARCH_LIST_LINUX "${GO_ARCH_X86};${GO_ARCH_X86_64};${GO_ARCH_ARM};${GO_ARCH_ARM64}")
set(GOARCH_LIST_DARWIN "${GO_ARCH_X86_64};${GO_ARCH_ARM64}")
# set(GOARCH_LIST_FREEBSD "${GO_ARCH_X86};${GO_ARCH_X86_64};${GO_ARCH_ARM};${GO_ARCH_ARM64}")
set(GOARCH_LIST_FREEBSD "${GO_ARCH_X86_64};${GO_ARCH_ARM64}")
# set(GOARCH_LIST_NETBSD "${GO_ARCH_X86};${GO_ARCH_X86_64};${GO_ARCH_ARM};${GO_ARCH_ARM64}")
set(GOARCH_LIST_NETBSD "${GO_ARCH_X86_64}")
# set(GOARCH_LIST_OPENBSD "${GO_ARCH_X86};${GO_ARCH_X86_64};${GO_ARCH_ARM};${GO_ARCH_ARM64}")
set(GOARCH_LIST_OPENBSD "${GO_ARCH_X86_64};${GO_ARCH_ARM64}")


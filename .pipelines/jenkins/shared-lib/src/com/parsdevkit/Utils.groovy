package com.parsdevkit

class Utils implements Serializable {

    def steps
    Utils(steps) {
        this.steps = steps
    }

    def normalizePath(String path) {
        if (isUnix()) {
            return path.replaceAll('\\\\', '/')
        } else {
            return path.replaceAll('/', '\\\\')
        }
    }

    String appExt(String platform) {
        return platform.toLowerCase() == 'windows' ? '.exe' : ''
    }
    String archiveFormat(String platform) {
        return platform.toLowerCase() == 'windows' ? 'zip' : 'tar.gz'
    }

    /**
     * Map general arch to platform-specific value
     * @param genericArch: x86, x86_64, arm, arm64
     * @param platform: linux, deb, rpm, windows, macos, openbsd, freebsd, netbsd
     * @return mapped architecture string
     */
    String mapArch(String platform, String genericArch) {
        def arch = genericArch?.toLowerCase()
        def plat = platform?.toLowerCase()

        switch (plat) {

            case ["linux", "debian", "ubuntu", "deb"]:
                return mapDebArch(arch)
            case ["rhel", "centos", "fedora", "rpm"]:
                return mapRpmArch(arch)
            case ["windows", "win"]:
                return mapWindowsArch(arch)
            case ["darwin", "macos", "mac"]:
                return mapMacArch(arch)
            case ["freebsd"]:
                return mapBsdArch(arch, "freebsd")
            case ["openbsd"]:
                return mapBsdArch(arch, "openbsd")
            case ["netbsd"]:
                return mapBsdArch(arch, "netbsd")
            default:
                steps.echo "⚠️ Unknown platform: ${platform}, using generic arch: ${arch}"
                return arch
        }
    }

    private String mapDebArch(String arch) {
        switch (arch) {
            case "x86_64": return "amd64"
            case "x86": return "i386"
            case "arm": return "armhf"
            case "arm64": return "arm64"
            default: return arch
        }
    }

    private String mapRpmArch(String arch) {
        switch (arch) {
            case "x86_64": return "x86_64"
            case "x86": return "i386"
            case "arm": return "armv7hl"
            case "arm64": return "aarch64"
            default: return arch
        }
    }

    private String mapWindowsArch(String arch) {
        switch (arch) {
            case "x86_64": return "x86_64"
            case "x86": return "x86"
            case "arm64": return "arm64"
            default: return arch
        }
    }

    private String mapMacArch(String arch) {
        switch (arch) {
            case "x86_64": return "x86_64"
            case "arm64":
            case "arm": return "arm64"
            default: return arch
        }
    }

    private String mapBsdArch(String arch, String bsdType) {
        switch (arch) {
            case "x86_64": return "amd64"
            case "x86": return "i386"
            case "arm": return "arm"
            case "arm64": return "aarch64"
            default: return arch
        }
    }
}

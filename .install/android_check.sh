function install::androidCheck() {
    function android_check() {
        [[ -x '/system/bin/getprop' ]] && return 0
        [[ -f '/system/bin/linker' || -f '/system/bin/linker64' ]] && return 0
        [[ -d '/dev/cpuctl' ]] && return 0
        [[ -d '/storage/emulated/0/Android' ]] && return 0
        [[ -d '/data/data/com.termux/' ]] && return 0
        [[ -n "${PREFIX}" && -n "${TERMUX_VERSION}" ]] && return 0
        return 1
    }

    local termux_api="$(
        command am start -n \
            'com.termux.api/invalid.activity' \
            2>&1
    )"

    android_check || {
        echo -e "${R}[!] ${N}Termux environment not detected."
        echo -e "${R}[!] ${N}This tool is designed exclusively for the Termux Android app."
        return 1
    }

    if [[ "${termux_api}" != *"Error: Activity class"* ]] && [[ "${termux_api}" != *"does not exist"* ]]; then
        echo -e "${R}[!] ${N}Termux:API not installed!"
        return 1
    fi
}
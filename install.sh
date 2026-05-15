#!/usr/bin/env bash
# https://github.com/Zeronetsec/TwrapKit

N='\033[0m'
R='\033[1;31m'
B='\033[1;34m'
GG='\033[0;32m'
DG='\033[1;90m'

base="${PREFIX}/opt"
symlink="${PREFIX}/bin"
bkdate="$(command date '+%Y_%b_%d_%H_%M_%S')"

path="$(
    cd -- "$(
        command dirname -- "${BASH_SOURCE[0]}"
    )" &> /dev/null && pwd
)"

if [[ "${1}" == "--backup" ]]; then
    backup="true"
fi

function install() {
    local cmd="${1}"
    local desc="${2}"
    echo -e "\n${B}[*] ${N}${desc}"
    eval "${cmd}" >/dev/null
    local status=$?
    echo -e "    ${DG}└── ${N}exit: ${GG}${status}${N}"
}

function android_check() {
    [[ -x '/system/bin/getprop' ]] && return 0
    [[ -f '/system/bin/linker' || -f '/system/bin/linker64' ]] && return 0
    [[ -d '/dev/cpuctl' ]] && return 0
    [[ -d '/storage/emulated/0/Android' ]] && return 0
    [[ -d '/data/data/com.termux/' ]] && return 0
    [[ -n "${PREFIX}" && -n "${TERMUX_VERSION}" ]] && return 0
    return 1
}

android_check || {
    echo -e "${R}[!] ${N}Termux environment not detected."
    echo -e "${R}[!] ${N}This tool is designed exclusively for the Termux Android app."
    exit 1
}

if [[ ! -d "${path}" ]]; then
    echo -e "${R}[!] ${N}Folder: ${GG}${path} ${N}not found! \n"
    exit 1
fi

termux_api="$(
    command am start -n \
        'com.termux.api/invalid.activity' \
        2>&1
)"

if [[ "${termux_api}" != *"Error: Activity class"* ]] && [[ "${termux_api}" != *"does not exist"* ]]; then
    echo -e "${R}[!] ${N}Termux:API not installed!"
    exit 1
fi

echo -e "${B}[*] ${N}Installing: ${GG}TwrapKit${N}"

pack=(
    "termux-api"
    "golang"
    "zip"
)

for i in "${pack[@]}"; do
    install \
        "command apt install -y ${i}" \
        "Installing: ${GG}${i}${N}"
done

if [[ ! -d "${base}" ]]; then
    install \
        "command mkdir -p ${base}" \
        "Create directory: ${GG}${base}${N}"
fi


if [[ "${backup}" == "true" && -d "${base}/twrapkit" ]]; then
    cd "${base}"
    install \
        "command zip -r twrapkit_${bkdate}.bak.zip twrapkit" \
        "Backup: ${GG}${base}/twrapkit ${DG}=> ${GG}${base}/twrapkit_${bkdate}.bak.zip${N}"
    cd
fi

if [[ -d "${base}/twrapkit" ]]; then
    install \
        "command rm -rf ${base}/twrapkit" \
        "Removing: ${GG}old twrapkit${N}"
fi

install \
    "command mv ${path} ${base}/twrapkit" \
    "Moving: ${GG}${path} ${DG}=> ${GG}${base}/twrapkit${N}"

cd "${base}/twrapkit"
install \
    "command go mod tidy" \
    "Retidy: ${GG}twrapkit${N}"

install \
    "command go build -v -o twrapkit" \
    "Building: ${GG}twrapkit${N}"
cd

install \
    "command ln -sf ${base}/twrapkit/twrapkit ${symlink}/twrapkit" \
    "Symlink: ${GG}${base}/twrapkit/twrapkit ${DG}=> ${GG}${symlink}/twrapkit${N}"

printf '\n'
if command -v twrapkit &>/dev/null; then
    echo -e "${GG}[+] ${N}TwrapKit installed!"
    echo -e "${GG}[+] ${N}Usage: ${GG}twrapkit --help ${N}to show helper"
    exit 0
else
    echo -e "${R}[!] ${N}Failed installing twrapkit!"
    exit 1
fi

# Copyright (c) 2026 Zeronetsec
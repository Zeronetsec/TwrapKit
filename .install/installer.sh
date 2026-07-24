function install::installer() {
    if [[ "${__BACKUP__}" == true && -d "${opt}/twrapkit" ]]; then
        (
            cd "${opt}"
            install::getinstall \
                "
                    command zip -r \
                        twrapkit_${bkdate}.bak.zip \
                        twrapkit
                " \
                "Backup: ${GG}${opt}/twrapkit ${DG}-> ${GG}${opt}/twrapkit_${bkdate}.bak.zip${N}"
            cd
        )
    fi

    if [[ -d "${opt}/twrapkit" ]]; then
        install::getinstall \
            "command rm -rf ${opt}/twrapkit" \
            "Removing old source..."
    fi

    install::getinstall \
        "command mv ${root} ${opt}/twrapkit" \
        "Moving: ${GG}${root} ${DG}-> ${GG}${opt}/twrapkit${N}"

    (
        cd "${opt}/twrapkit"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}twrapkit${N}"

        install::getinstall \
            "command go build -o twrapkit" \
            "Compiling: ${GG}twrapkit${N}"
        cd
    )

    install::getinstall \
        "
            command ln -sf \
                ${opt}/twrapkit/twrapkit \
                ${bin}/twrapkit
        " \
        "Symlink: ${GG}${opt}/twrapkit/twrapkit ${DG}-> ${GG}${bin}/twrapkit${N}"
}; readonly -f install::installer
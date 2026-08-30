function install::installer() {
    (
        cd "${opt}/${targetins}"
        install::getinstall \
            "
                command go mod tidy
                command go build -o ${targetins}
            " \
            "Compiling: ${color_GG}${targetins}${color_N}"
    )
}; readonly -f install::installer
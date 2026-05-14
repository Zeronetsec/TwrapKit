// https://github.com/Zeronetsec/TwrapKit

package service

import (
    "fmt"
    "twrapkit/utils/shell"
    "twrapkit/utils/color"
)

func Runner(action string) {
    var cmd string
    var prefix string

    switch action {
        case "start":
            cmd = "termux-api-start"
            prefix = "Starting service"
        case "stop":
            cmd = "termux-api-stop"
            prefix = "Stopping service"
        default:
            fmt.Printf(
                "%s[!] %sUnknown API action: %s%s%s\n",
                color.R, color.N, color.GG, action, color.N,
            )
            return
    }

    result, err := shell.ExecShell(cmd)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    target := parser(result.Stdout)
    if target == "" {
        fmt.Printf(
            "%s[!] %sFailed to parse target service!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %s%s: %s%s%s\n",
        color.B, color.N, prefix, color.GG, target, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
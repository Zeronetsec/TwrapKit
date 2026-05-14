// https://github.com/Zeronetsec/TwrapKit

package wificonninfo

import (
    "fmt"
    "encoding/json"
    "twrapkit/utils/shell"
    "twrapkit/utils/color"
)

func WFConn() {
    result, err := shell.ExecShell("termux-wifi-connectioninfo")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to get wifi connection info!\n",
            color.R, color.N,
        )
        return
    }

    var conninfo ConnectionInfo

    err = json.Unmarshal([]byte(result.Stdout), &conninfo)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(conninfo)
}

// Copyright (c) 2026 Zeronetsec
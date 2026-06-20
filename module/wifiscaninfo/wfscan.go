// https://github.com/Zeronetsec/TwrapKit

package wifiscaninfo

import (
    "fmt"
    "strings"
    "encoding/json"
    "github.com/Zeronetsec/TwrapKit/utils/shell"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func WFScan() {
    result, err := shell.ExecShell("termux-wifi-scaninfo")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to execute wifi scaninfo!\n",
            color.R, color.N,
        )
        return
    }

    data := strings.TrimSpace(result.Stdout)

    if isError(data) {
        handleError(data)
        return
    }

    var scans []ScanItem

    err = json.Unmarshal([]byte(data), &scans)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(scans)
}

// Copyright (c) 2026 Zeronetsec
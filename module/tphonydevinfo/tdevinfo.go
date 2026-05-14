// https://github.com/Zeronetsec/TwrapKit

package tphonydevinfo

import (
    "fmt"
    "strings"
    "encoding/json"
    "twrapkit/utils/shell"
    "twrapkit/utils/color"
)

func TDevInfo() {
    result, err := shell.ExecShell("termux-telephony-deviceinfo")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to get telephony info!\n",
            color.R, color.N,
        )
        return
    }

    data := strings.TrimSpace(result.Stdout)

    if strings.Contains(data, `"error"`) {
        var e ErrorResp
        _ = json.Unmarshal([]byte(data), &e)

        fmt.Printf(
            "%s[!] %sPermission error: %s%s%s\n",
            color.R, color.N, color.GG, e.Error, color.N,
        )
        return
    }

    var info Info

    err = json.Unmarshal([]byte(data), &info)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(info)
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package batterystat

import (
    "fmt"
    "encoding/json"
    "github.com/Zeronetsec/TwrapKit/utils/shell"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func Battery() {
    result, err := shell.ExecShell("termux-battery-status")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to get battery status!\n",
            color.R, color.N,
        )
        return
    }

    var data Info

    err = json.Unmarshal([]byte(result.Stdout), &data)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(data)
}

// Copyright (c) 2026 Zeronetsec
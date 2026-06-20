// https://github.com/Zeronetsec/TwrapKit

package audioinfo

import (
    "fmt"
    "encoding/json"
    "github.com/Zeronetsec/TwrapKit/utils/shell"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func Audio() {
    result, err := shell.ExecShell("termux-audio-info")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to get audio info!\n",
            color.R, color.N,
        )
        return
    }

    var info Info

    err = json.Unmarshal([]byte(result.Stdout), &info)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to parse json: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(info)
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package camerainfo

import (
    "fmt"
    "encoding/json"
    "twrapkit/utils/shell"
    "twrapkit/utils/color"
)

func CamInfo() {
    result, err := shell.ExecShell("termux-camera-info")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to get camera info!\n",
            color.R, color.N,
        )
        return
    }

    var cams []Info

    err = json.Unmarshal([]byte(result.Stdout), &cams)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(cams)
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package wifiscaninfo

import (
    "fmt"
    "encoding/json"
    "twrapkit/utils/color"
)

func handleError(data string) {
    var e ErrorResp
    _ = json.Unmarshal([]byte(data), &e)

    if e.Error != "" {
        fmt.Printf(
            "%s[!] %sPermission error: %s%s%s\n",
            color.R, color.N, color.GG, e.Error, color.N,
        )
        return
    }

    if e.APIError != "" {
        fmt.Printf(
            "%s[!] %sAPI error: %s%s%s\n",
            color.R, color.N, color.GG, e.APIError, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[!] %sUnknown error format!\n",
        color.R, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
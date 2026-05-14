// https://github.com/Zeronetsec/TwrapKit

package invinput

import (
    "fmt"
    "twrapkit/utils/color"
)

func Invalid() {
    tname := "twrapkit"

    fmt.Printf(
        "%s[!] %sInvalid input!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %s%s --help%s\n",
        color.R, color.N, color.GG, tname, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
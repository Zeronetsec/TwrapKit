// https://github.com/Zeronetsec/TwrapKit

package invinput

import (
    "fmt"
    "twrapkit/utils/color"
)

func Unknown(input string) {
    tname := "twrapkit"

    fmt.Printf(
        "%s[!] %sInvalid command: %s%s%s\n",
        color.R, color.N, color.GG, input, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %s%s --help%s\n",
        color.R, color.N, color.GG, tname, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
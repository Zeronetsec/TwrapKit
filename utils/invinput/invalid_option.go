// https://github.com/Zeronetsec/TwrapKit

package invinput

import (
    "fmt"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func InvalidOption(input string) {
    fmt.Printf(
        "%s[!] %sInvalid option: %s%s%s\n",
        color.R, color.N, color.GG, input, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %stwrapkit --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
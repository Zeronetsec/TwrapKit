// https://github.com/Zeronetsec/TwrapKit

package cursor

import (
    "fmt"
)

func Hide() {
    fmt.Print("\033[?25l")
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package cursor

import (
    "fmt"
)

func Hide() {
    fmt.Print("\x1b[?25l")
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package cursor

import (
    "fmt"
)

func Visible() {
    fmt.Print("\033[?25h")
}

// Copyright (c) 2026 Zeronetsec
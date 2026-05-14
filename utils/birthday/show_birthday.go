// https://github.com/Zeronetsec/TwrapKit

package birthday

import (
    "fmt"
    "time"
    "twrapkit/utils/color"
)

func ShowBirthDay() {
    bdate := "05-15"
    now := time.Now().Format("01-02")
    tname := "twrapkit"

    if now == bdate {
        fmt.Printf(
            "%s› %sHappy birthday for %s%s %s🎉\n",
            color.R, color.N, color.GG, tname, color.N,
        )
        fmt.Println()
    }
}

// Copyright (c) 2026 Zeronetsec
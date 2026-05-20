// https://github.com/Zeronetsec/TwrapKit

package uwu

import (
    "fmt"
    "time"
)

func Uwu(duration time.Duration) {
    faces := []string{
        "(｡◕‿◕｡)",
        "(≧◡≦)",
        "ʕ•ᴥ•ʔ",
        "(・ω・)",
        "(๑˃ᴗ˂)ﻭ",
        "(ง'̀-'́)ง",
        "(=^･ω･^=)",
    }

    delay := 200 * time.Millisecond
    end := time.After(duration)
    kaomoji := 0

    for {
        select {
            case <-end:
                fmt.Print("\033[K")
                return
            default:
                fmt.Printf(
                    "\r%s\033[K",
                    faces[kaomoji%len(faces)],
                )
            time.Sleep(delay)
            kaomoji++
        }
    }
}

// Copyright (c) 2026 Zeronetsec
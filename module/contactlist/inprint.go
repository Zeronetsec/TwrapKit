// https://github.com/Zeronetsec/TwrapKit

package contactlist

import (
    "fmt"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func inprint(list []Contact) {
    if len(list) == 0 {
        fmt.Printf(
            "%s[!] %sNo contacts found!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%sTotal contacts: %s%d%s\n",
        color.N, color.GG, len(list), color.N,
    )

    fmt.Println()

    for i, c := range list {
        if i > 0 {
            fmt.Println()
        }

        fmt.Printf(
            "%sContact: %s%d%s\n",
            color.N, color.B, i+1, color.N,
        )

        fmt.Printf(
            "%sName: %s%s%s\n",
            color.N, color.GG, c.Name, color.N,
        )

        fmt.Printf(
            "%sNumber: %s%s%s\n",
            color.N, color.GG, c.Number, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package calllog

import (
    "fmt"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func inprint(logs []Entry) {
    if len(logs) == 0 {
        fmt.Printf(
            "%s[!] %sNo call history found!\n",
            color.R, color.N,
        )
        return
    }

    for i, c := range logs {
        if i > 0 {
            fmt.Println()
        }

        fmt.Printf(
            "%sLog: %s%d%s\n",
            color.N, color.B, i+1, color.N,
        )

        fmt.Printf(
            "%sName: %s%s%s\n",
            color.N, color.GG, c.Name, color.N,
        )

        fmt.Printf(
            "%sPhone number: %s%s%s\n",
            color.N, color.GG, c.PhoneNumber, color.N,
        )

        fmt.Printf(
            "%sType: %s%s%s\n",
            color.N, color.GG, c.Type, color.N,
        )

        fmt.Printf(
            "%sDate: %s%s%s\n",
            color.N, color.GG, c.Date, color.N,
        )

        fmt.Printf(
            "%sDuration: %s%s%s\n",
            color.N, color.GG, c.Duration, color.N,
        )

        if c.SimID != nil {
            fmt.Printf(
                "%sSIM ID: %s%s%s\n",
                color.N, color.GG, *c.SimID, color.N,
            )
        }
    }
}

// Copyright (c) 2026 Zeronetsec
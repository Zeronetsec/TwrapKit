// https://github.com/Zeronetsec/TwrapKit

package contactlist

import (
    "fmt"
    "strings"
    "encoding/json"
    "github.com/Zeronetsec/TwrapKit/utils/shell"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func DumpList() {
    result, err := shell.ExecShell("termux-contact-list")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to get contact list!\n",
            color.R, color.N,
        )
        return
    }

    data := strings.TrimSpace(result.Stdout)
    if strings.Contains(data, `"error"`) {
        var e ErrorResp
        _ = json.Unmarshal([]byte(data), &e)

        fmt.Printf(
            "%s[!] %sPermission error: %s%s%s\n",
            color.R, color.N, color.GG, e.Error, color.N,
        )
        return
    }

    var list []Contact
    err = json.Unmarshal([]byte(data), &list)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(list)
}

// Copyright (c) 2026 Zeronetsec
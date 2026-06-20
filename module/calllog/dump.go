// https://github.com/Zeronetsec/TwrapKit

package calllog

import (
    "fmt"
    "strings"
    "encoding/json"
    "github.com/Zeronetsec/TwrapKit/utils/shell"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func Dump() {
    result, err := shell.ExecShell("termux-call-log")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to get call log!\n",
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

    var logs []Entry
    err = json.Unmarshal([]byte(data), &logs)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    inprint(logs)
}

// Copyright (c) 2026 Zeronetsec
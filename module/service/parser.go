// https://github.com/Zeronetsec/TwrapKit

package service

import (
    "regexp"
    "strings"
)

func parser(output string) string {
    output = strings.TrimSpace(output)
    re := regexp.MustCompile(`cmp=([^\s}]+)`)
    match := re.FindStringSubmatch(output)

    if len(match) < 2 {
        return ""
    }

    return match[1]
}

// Copyright (c) 2026 Zeronetsec
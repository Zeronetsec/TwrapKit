// https://github.com/Zeronetsec/TwrapKit

package main

import (
    "os"
    "strings"
    "twrapkit/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.TwrapKitConsole(input)
}

// Copyright (c) 2026 Zeronetsec
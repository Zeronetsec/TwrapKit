// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "time"
    "fmt"
    "github.com/Zeronetsec/TwrapKit/module/uwu"
    "github.com/Zeronetsec/TwrapKit/utils/cursor"
)

type UWU struct{}
func (c UWU) Execute(args []string) {
    cursor.Hide()
    uwu.Uwu(5 * time.Second)
    cursor.Visible()

    fmt.Println()
}

// Copyright (c) 2026 Zeronetsec
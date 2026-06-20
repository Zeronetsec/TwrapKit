// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/helper"
)

type Helper struct{}
func (c Helper) Execute(args []string) {
    helper.ShowHelper()
}

// Copyright (c) 2026 Zeronetsec
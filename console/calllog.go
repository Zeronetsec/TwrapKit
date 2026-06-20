// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/calllog"
)

type CallLog struct{}
func (c CallLog) Execute(args []string) {
    calllog.Dump()
}

// Copyright (c) 2026 Zeronetsec
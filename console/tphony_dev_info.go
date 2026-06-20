// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/tphonydevinfo"
)

type TPhonyDevInfo struct{}
func (c TPhonyDevInfo) Execute(args []string) {
    tphonydevinfo.TDevInfo()
}

// Copyright (c) 2026 Zeronetsec
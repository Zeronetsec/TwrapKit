// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/wificonninfo"
)

type WifiConnInfo struct{}
func (c WifiConnInfo) Execute(args []string) {
    wificonninfo.WFConn()
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/wifiscaninfo"
)

type WifiScanInfo struct{}
func (c WifiScanInfo) Execute(args []string) {
    wifiscaninfo.WFScan()
}

// Copyright (c) 2026 Zeronetsec
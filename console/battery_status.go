// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/batterystat"
)

type BatteryStatus struct{}
func (c BatteryStatus) Execute(args []string) {
    batterystat.Battery()
}

// Copyright (c) 2026 Zeronetsec
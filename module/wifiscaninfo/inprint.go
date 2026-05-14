// https://github.com/Zeronetsec/TwrapKit

package wifiscaninfo

import (
    "fmt"
    "twrapkit/utils/color"
)

func inprint(scans []ScanItem) {
    if len(scans) == 0 {
        fmt.Printf(
            "%s[!] %sNo networks found!\n",
            color.R, color.N,
        )
        return
    }

    for i, s := range scans {
        if i > 0 {
            fmt.Println()
        }

        fmt.Printf(
            "%sConnect: %s%d%s\n",
            color.N, color.B, i+1, color.N,
        )

        fmt.Printf(
            "%sSSID: %s%s%s\n",
            color.N, color.GG, s.SSID, color.N,
        )

        fmt.Printf(
            "%sBSSID: %s%s%s\n",
            color.N, color.GG, s.BSSID, color.N,
        )

        fmt.Printf(
            "%sRSSI: %s%d dBm%s\n",
            color.N, color.GG, s.RSSI, color.N,
        )

        fmt.Printf(
            "%sFrequency: %s%d MHz%s\n",
            color.N, color.GG, s.FrequencyMHz, color.N,
        )

        fmt.Printf(
            "%sChannel BW: %s%s%s\n",
            color.N, color.GG, s.ChannelBandwidthMHz, color.N,
        )

        fmt.Printf(
            "%sCapabilities: %s%s%s\n",
            color.N, color.GG, s.Capabilities, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec
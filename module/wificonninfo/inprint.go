// https://github.com/Zeronetsec/TwrapKit

package wificonninfo

import (
    "fmt"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func inprint(i ConnectionInfo) {
    fmt.Printf(
        "%sSSID: %s%s%s\n",
        color.N, color.GG, i.SSID, color.N,
    )

    fmt.Printf(
        "%sHidden SSID: %s%t%s\n",
        color.N, color.GG, i.SSIDHidden, color.N,
    )

    fmt.Printf(
        "%sBSSID: %s%s%s\n",
        color.N, color.GG, i.BSSID, color.N,
    )

    fmt.Printf(
        "%sMAC address: %s%s%s\n",
        color.N, color.GG, i.MACAddress, color.N,
    )

    fmt.Printf(
        "%sIP address: %s%s%s\n",
        color.N, color.GG, i.IP, color.N,
    )

    fmt.Printf(
        "%sRSSI: %s%d dBm%s\n",
        color.N, color.GG, i.RSSI, color.N,
    )

    fmt.Printf(
        "%sLink Speed: %s%d Mbps%s\n",
        color.N, color.GG, i.LinkSpeedMbps, color.N,
    )

    fmt.Printf(
        "%sFrequency: %s%d MHz%s\n",
        color.N, color.GG, i.FrequencyMHz, color.N,
    )

    fmt.Printf(
        "%sNetwork ID: %s%d%s\n",
        color.N, color.GG, i.NetworkID, color.N,
    )

    fmt.Printf(
        "%sSupplicant state: %s%s%s\n",
        color.N, color.GG, i.SupplicantState, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
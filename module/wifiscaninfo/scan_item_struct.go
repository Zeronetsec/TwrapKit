// https://github.com/Zeronetsec/TwrapKit

package wifiscaninfo

type ScanItem struct {
    BSSID string `json:"bssid"`
    FrequencyMHz int `json:"frequency_mhz"`
    RSSI int `json:"rssi"`
    SSID string `json:"ssid"`
    Timestamp int64 `json:"timestamp"`
    ChannelBandwidthMHz string `json:"channel_bandwidth_mhz"`
    Capabilities string `json:"capabilities"`
}

// Copyright (c) 2026 Zeronetsec
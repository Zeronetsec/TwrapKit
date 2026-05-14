// https://github.com/Zeronetsec/TwrapKit

package wificonninfo

type ConnectionInfo struct {
    BSSID string `json:"bssid"`
    FrequencyMHz int `json:"frequency_mhz"`
    IP string `json:"ip"`
    LinkSpeedMbps int `json:"link_speed_mbps"`
    MACAddress string `json:"mac_address"`
    NetworkID int `json:"network_id"`
    RSSI int `json:"rssi"`
    SSID string `json:"ssid"`
    SSIDHidden bool `json:"ssid_hidden"`
    SupplicantState string `json:"supplicant_state"`
}

// Copyright (c) 2026 Zeronetsec
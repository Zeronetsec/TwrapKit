// https://github.com/Zeronetsec/TwrapKit

package batterystat

type Info struct {
    Present bool `json:"present"`
    Technology string `json:"technology"`
    Health string `json:"health"`
    Plugged string `json:"plugged"`
    Status string `json:"status"`
    Temperature float64 `json:"temperature"`
    Voltage int `json:"voltage"`
    Current int `json:"current"`
    CurrentAverage int `json:"current_average"`
    Percentage int `json:"percentage"`
    Level int `json:"level"`
    Scale int `json:"scale"`
    ChargeCounter int `json:"charge_counter"`
    Cycle int `json:"cycle"`
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package batterystat

import (
    "fmt"
    "twrapkit/utils/color"
)

func inprint(b Info) {
    fmt.Printf(
        "%sPresent: %s%t%s\n",
        color.N, color.GG, b.Present, color.N,
    )

    fmt.Printf(
        "%sTechnology: %s%s%s\n",
        color.N, color.GG, b.Technology, color.N,
    )

    fmt.Printf(
        "%sHealth: %s%s%s\n",
        color.N, color.GG, b.Health, color.N,
    )

    fmt.Printf(
        "%sPlugged: %s%s%s\n",
        color.N, color.GG, b.Plugged, color.N,
    )

    fmt.Printf(
        "%sStatus: %s%s%s\n",
        color.N, color.GG, b.Status, color.N,
    )

    fmt.Printf(
        "%sTemperature: %s%.1f°C%s\n",
        color.N, color.GG, b.Temperature, color.N,
    )

    fmt.Printf(
        "%sVoltage: %s%d mV%s\n",
        color.N, color.GG, b.Voltage, color.N,
    )

    fmt.Printf(
        "%sCurrent: %s%d µA%s\n",
        color.N, color.GG, b.Current, color.N,
    )

    fmt.Printf(
        "%sAvg current: %s%d µA%s\n",
        color.N, color.GG, b.CurrentAverage, color.N,
    )

    fmt.Printf(
        "%sBattery: %s%d%%%s\n",
        color.N, color.GG, b.Percentage, color.N,
    )

    fmt.Printf(
        "%sLevel: %s%d/%d%s\n",
        color.N, color.GG, b.Level, b.Scale, color.N,
    )

    fmt.Printf(
        "%sCharge counter: %s%d%s\n",
        color.N, color.GG, b.ChargeCounter, color.N,
    )

    fmt.Printf(
        "%sCycle: %s%d%s\n",
        color.N, color.GG, b.Cycle, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
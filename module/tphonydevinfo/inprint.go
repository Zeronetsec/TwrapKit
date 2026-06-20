// https://github.com/Zeronetsec/TwrapKit

package tphonydevinfo

import (
    "fmt"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func inprint(t Info) {
    fmt.Printf(
        "%sPhone count: %s%d%s\n",
        color.N, color.GG, t.PhoneCount, color.N,
    )

    fmt.Printf(
        "%sPhone type: %s%s%s\n",
        color.N, color.GG, t.PhoneType, color.N,
    )

    fmt.Printf(
        "%sSIM state: %s%s%s\n",
        color.N, color.GG, t.SimState, color.N,
    )

    fmt.Printf(
        "%sSIM operator: %s%s%s\n",
        color.N, color.GG, t.SimOperatorName, color.N,
    )

    fmt.Printf(
        "%sSIM code: %s%s%s\n",
        color.N, color.GG, t.SimOperator, color.N,
    )

    fmt.Printf(
        "%sSIM country: %s%s%s\n",
        color.N, color.GG, t.SimCountryISO, color.N,
    )

    if t.SimSerialNumber != nil {
        fmt.Printf(
            "%sSIM serial: %s%s%s\n",
            color.N, color.GG, *t.SimSerialNumber, color.N,
        )
    }

    if t.SimSubscriberID != nil {
        fmt.Printf(
            "%sSubscriber ID: %s%s%s\n",
            color.N, color.GG, *t.SimSubscriberID, color.N,
        )
    }

    fmt.Printf(
        "%sData enabled: %s%s%s\n",
        color.N, color.GG, t.DataEnabled, color.N,
    )

    fmt.Printf(
        "%sData activity: %s%s%s\n",
        color.N, color.GG, t.DataActivity, color.N,
    )

    fmt.Printf(
        "%sData state: %s%s%s\n",
        color.N, color.GG, t.DataState, color.N,
    )

    fmt.Printf(
        "%sNetwork type: %s%s%s\n",
        color.N, color.GG, t.NetworkType, color.N,
    )

    fmt.Printf(
        "%sNetwork operator: %s%s%s\n",
        color.N, color.GG, t.NetworkOperatorName, color.N,
    )

    fmt.Printf(
        "%sRoaming: %s%t%s\n",
        color.N, color.GG, t.NetworkRoaming, color.N,
    )

    fmt.Printf(
        "%sSoftware version: %s%s%s\n",
        color.N, color.GG, t.DeviceSoftwareVer, color.N,
    )

    if t.DeviceID != nil {
        fmt.Printf(
            "%sDevice ID: %s%s%s\n",
            color.N, color.GG, *t.DeviceID, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec
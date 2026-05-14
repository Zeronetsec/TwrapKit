// https://github.com/Zeronetsec/TwrapKit

package version

import (
    "fmt"
    "twrapkit/utils/color"
)

const (
    name = "TwrapKit"
    version = "v1.0"
    creator = "Zeronetsec"
    homepage = "https://github.com/Zeronetsec/TwrapKit"
)

func ShowVersion() {
    fmt.Printf(
        "%sName: %s%s%s\n",
        color.N, color.GG, name, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sCreator: %s%s%s\n",
        color.N, color.GG, creator, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
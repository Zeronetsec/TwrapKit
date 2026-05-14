// https://github.com/Zeronetsec/TwrapKit

package camerainfo

import (
    "fmt"
    "twrapkit/utils/color"
)

func inprint(cams []Info) {
    if len(cams) == 0 {
        fmt.Printf(
            "%s[!] %sNo camera found!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%sTotal Camera: %s%d%s\n",
        color.N, color.GG, len(cams), color.N,
    )

    fmt.Println()

    for i, cam := range cams {
        if i > 0 {
            fmt.Println()
        }

        fmt.Printf(
            "%sCamera: %s%d%s\n",
            color.N, color.B, i+1, color.N,
        )

        fmt.Printf(
            "%sID: %s%s%s\n",
            color.N, color.GG, cam.ID, color.N,
        )

        fmt.Printf(
            "%sFacing: %s%s%s\n",
            color.N, color.GG, cam.Facing, color.N,
        )

        fmt.Printf(
            "%sJPEG sizes: %s%d modes%s\n",
            color.N, color.GG, len(cam.JPEGOutputSizes), color.N,
        )

        if len(cam.JPEGOutputSizes) > 0 {
            first := cam.JPEGOutputSizes[0]
            last := cam.JPEGOutputSizes[len(cam.JPEGOutputSizes)-1]

            fmt.Printf(
                "%sResolution max: %s%dx%d%s\n",
                color.N, color.GG, first.Width, first.Height, color.N,
            )

            fmt.Printf(
                "%sResolution min: %s%dx%d%s\n",
                color.N, color.GG, last.Width, last.Height, color.N,
            )
        }

        fmt.Printf(
            "%sFocal lengths: %s",
            color.N, color.GG,
        )

        for j, f := range cam.FocalLengths {
            if j > 0 {
                fmt.Printf(", ")
            }
            fmt.Printf("%.2f", f)
        }

        fmt.Printf("%s\n", color.N)

        fmt.Printf(
            "%sAE modes: %s%d%s\n",
            color.N, color.GG, len(cam.AutoExposureModes), color.N,
        )

        for _, mode := range cam.AutoExposureModes {
            fmt.Printf(
                "    %s- %s%s%s\n",
                color.DG, color.CC, mode, color.N,
            )
        }

        fmt.Printf(
            "%sSensor size: %s%.2f x %.2f mm%s\n",
            color.N, color.GG, cam.PhysicalSize.Width, cam.PhysicalSize.Height, color.N,
        )

        fmt.Printf(
            "%sCapabilities: %s%d%s\n",
            color.N, color.GG, len(cam.Capabilities), color.N,
        )

        for _, cpbl := range cam.Capabilities {
            fmt.Printf(
                "    %s- %s%s%s\n",
                color.DG, color.CC, cpbl, color.N,
            )
        }
    }
}

// Copyright (c) 2026 Zeronetsec
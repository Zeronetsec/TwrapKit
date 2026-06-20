// https://github.com/Zeronetsec/TwrapKit

package audioinfo

import (
    "fmt"
    "github.com/Zeronetsec/TwrapKit/utils/color"
)

func inprint(i Info) {
    fmt.Printf(
        "%sOutput sample rate: %s%s%s\n",
        color.N, color.GG, i.OutputSampleRate, color.N,
    )

    fmt.Printf(
        "%sFrames buffer: %s%s%s\n",
        color.N, color.GG, i.OutputFramesPerBuffer, color.N,
    )

    fmt.Printf(
        "%sSample rate: %s%d%s\n",
        color.N, color.GG, i.AudioTrackSampleRate, color.N,
    )

    fmt.Printf(
        "%sLow latency: %s%d%s\n",
        color.N, color.GG, i.AudioTrackLowLatencyRate, color.N,
    )

    fmt.Printf(
        "%sBluetooth A2DP: %s%t%s\n",
        color.N, color.GG, i.BluetoothA2DP, color.N,
    )

    fmt.Printf(
        "%sWired headset: %s%t%s\n",
        color.N, color.GG, i.WiredHeadset, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec
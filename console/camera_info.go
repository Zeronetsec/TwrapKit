// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/camerainfo"
)

type CameraInfo struct{}
func (c CameraInfo) Execute(args []string) {
    camerainfo.CamInfo()
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/audioinfo"
)

type AudioInfo struct{}
func (c AudioInfo) Execute(args []string) {
    audioinfo.Audio()
}

// Copyright (c) 2026 Zeronetsec
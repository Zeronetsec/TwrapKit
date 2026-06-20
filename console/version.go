// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.ShowVersion()
}

// Copyright (c) 2026 Zeronetsec
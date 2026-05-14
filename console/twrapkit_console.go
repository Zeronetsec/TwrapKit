// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "os"
    "twrapkit/utils/invinput"
)

func TwrapKitConsole(input string) {
    args := os.Args
    if len(args) < 2 {
        invinput.Invalid()
        return
    }

    commands := map[string]Command{
        "--help": Helper{},
        "--version": Version{},
        "--uwu": UWU{},
        "--service": Service{},
        "--audio-info": AudioInfo{},
        "--battery-status": BatteryStatus{},
        "--wifi-scaninfo": WifiScanInfo{},
        "--wifi-conninfo": WifiConnInfo{},
        "--telephony-deviceinfo": TPhonyDevInfo{},
        "--call-log": CallLog{},
        "--contact-list": ContactList{},
        "--camera-info": CameraInfo{},
    }

    if cmd, ok := commands[args[1]]; ok {
        cmd.Execute(args)
    } else {
        invinput.Unknown(args[1])
    }
}

// Copyright (c) 2026 Zeronetsec
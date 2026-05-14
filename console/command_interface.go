// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "time"
    "twrapkit/utils/cursor"
    "twrapkit/utils/helper"
    "twrapkit/utils/version"
    "twrapkit/utils/uwu"
    "twrapkit/utils/invinput"
    "twrapkit/module/service"
    "twrapkit/module/audioinfo"
    "twrapkit/module/batterystat"
    "twrapkit/module/wifiscaninfo"
    "twrapkit/module/wificonninfo"
    "twrapkit/module/tphonydevinfo"
    "twrapkit/module/calllog"
    "twrapkit/module/contactlist"
    "twrapkit/module/camerainfo"
)

type Command interface {
    Execute(args []string)
}

type Helper struct{}
func (c Helper) Execute(args []string) {
    helper.ShowHelper()
}

type Version struct{}
func (c Version) Execute(args []string) {
    version.ShowVersion()
}

type UWU struct{}
func (c UWU) Execute(args []string) {
    cursor.Hide()
    uwu.Uwu(5 * time.Second)
    cursor.Visible()
}

type Service struct{}
func (c Service) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    service.Runner(args[2])
}

type AudioInfo struct{}
func (c AudioInfo) Execute(args []string) {
    audioinfo.Audio()
}

type BatteryStatus struct{}
func (c BatteryStatus) Execute(args []string) {
    batterystat.Battery()
}

type WifiScanInfo struct{}
func (c WifiScanInfo) Execute(args []string) {
    wifiscaninfo.WFScan()
}

type WifiConnInfo struct{}
func (c WifiConnInfo) Execute(args []string) {
    wificonninfo.WFConn()
}

type TPhonyDevInfo struct{}
func (c TPhonyDevInfo) Execute(args []string) {
    tphonydevinfo.TDevInfo()
}

type CallLog struct{}
func (c CallLog) Execute(args []string) {
    calllog.Dump()
}

type ContactList struct{}
func (c ContactList) Execute(args []string) {
    contactlist.DumpList()
}

type CameraInfo struct{}
func (c CameraInfo) Execute(args []string) {
    camerainfo.CamInfo()
}

// Copyright (c) 2026 Zeronetsec
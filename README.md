<!-- https://github.com/Zeronetsec/TwrapKit -->

[![version](https://img.shields.io/badge/TwrapKit-Version%200.1-blue.svg?maxAge=259200)]()
[![os](https://img.shields.io/badge/Supported%20OS-Android-blue.svg)]()
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

# TwrapKit
TwrapKit is a CLI wrapper for Termux:API designed for mobile automation. <br>
It provides simplified access to Android system features and sensors.

## Features
- Inspect system, battery, and hardware capabilities
- Monitor network status and Wi-Fi connection details
- Access telephony data, call logs, and contact lists
- Manage background services and tool lifecycle
- Retrieve camera and audio device information
- And more.

## Disclaimer
This tool interacts with the Android API via Termux:API. <br>
Please read the
[DISCLAIMER](https://github.com/Zeronetsec/TwrapKit/blob/main/DISCLAIMER.md)
before use.

## Installation
```bash
git clone https://github.com/Zeronetsec/TwrapKit.git
cd TwrapKit
chmod +x install.sh
./install.sh

# for backup
./install.sh --backup
```

## Usage
``` bash
twrapkit --service <start|stop>
twrapkit --camera-info
twrapkit --wifi-conninfo
twrapkit --telephony-deviceinfo
twrapkit --contact-list
```
And more commands.

## License
This project is licensed under the MIT License.

<!-- Copyright (c) 2026 Zeronetsec -->
<!-- https://github.com/Zeronetsec/TwrapKit -->

<div style="display: flex; gap: 5px; justify-content: left;">
    <img src="https://img.shields.io/badge/TwrapKit-Version%200.1-blue.svg?style=square&logo=go" />
    <img src="https://img.shields.io/badge/Supported%20OS-Android-blue.svg?style=square&logo=android" />
    <img src="https://img.shields.io/badge/License-MIT-blue.svg?style&logo=github" />
</div>

# TwrapKit
TwrapKit is a Termux:API wrapper for Android system and hardware access.

## Features
- Inspect system, battery, and hardware capabilities.
- Monitor network status and Wi-Fi connection details.
- Access telephony data, call logs, and contact lists.
- Retrieve camera and audio device information.
- And more features.

## Disclaimer
Please read [.docs/disclaimer.md](.docs/disclaimer.md) before using this tool. </br>
Use this software at your own risk. </br>
The author is not responsible for any damage, data loss, or issues that may result from its use.

## Installation
Quick install:
```bash
git clone https://github.com/Zeronetsec/TwrapKit
cd TwrapKit
chmod +x install.sh
./install.sh
```
For more detailed installation and uninstallation instructions, see [.docs/install_and_uninstall.md](.docs/install_and_uninstall.md).

## Usage Example
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
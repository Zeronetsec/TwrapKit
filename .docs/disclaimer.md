<!-- https://github.com/Zeronetsec/TwrapKit -->

# DISCLAIMER
## **Version 0.1 (Experimental Status)**
This tool is currently in its early **v0.1** stage and is considered **unstable**. </br>
You may encounter bugs, system errors, or unexpected behavior. </br>
Additionally, core structures, commands, and language configurations are highly subject to change in future updates.

## Termux:API Wrapper
**TwrapKit** is an interface built to interact with the **Termux:API** toolset. </br>
It simplifies command execution but depends entirely on the underlying Termux:API installation for data access and hardware-related features.

## External Dependencies
Installing the **termux-api** package alone is not sufficient for full functionality. </br>
You must also install the **Termux:API Android App** on your device. </br>
Without the Android application, TwrapKit cannot communicate with Android system features. </br>
Download from the official F-Droid repository: [f-droid.org/id/packages/com.termux.api/](https://f-droid.org/id/packages/com.termux.api/)

## Permissions & Access
Termux:API requires access to Android system features via user-granted permissions. </br>
For TwrapKit to function properly, you must manually grant the required permissions (Camera, Contacts, Location, Telephone, etc.) when prompted by Android or through the App Info settings.
- If a command fails or returns no data, ensure the corresponding permission has been granted to the Termux:API app.

<!-- Copyright (c) 2026 Zeronetsec -->
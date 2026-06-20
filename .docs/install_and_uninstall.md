<!-- https://github.com/Zeronetsec/TwrapKit -->

# Installation
`install.sh` optional option:
- `--backup`

Use `--backup` to create a backup of the existing TwrapKit installation before replacing it.

## Termux (only)
```bash
git clone https://github.com/Zeronetsec/TwrapKit
cd TwrapKit
chmod +x install.sh
./install.sh
```

## Uninstallation
```bash
export prefix="${PREFIX:-/usr}"
rm -f "${prefix}/bin/twrapkit"
rm -rf "${prefix}/opt/twrapkit"
```

<!-- Copyright (c) 2026 Zeronetsec -->
// https://github.com/Zeronetsec/TwrapKit

package console

import (
    "github.com/Zeronetsec/TwrapKit/module/contactlist"
)

type ContactList struct{}
func (c ContactList) Execute(args []string) {
    contactlist.DumpList()
}

// Copyright (c) 2026 Zeronetsec
// https://github.com/Zeronetsec/TwrapKit

package calllog

type Entry struct {
    Name string `json:"name"`
    PhoneNumber string `json:"phone_number"`
    Type string `json:"type"`
    Date string `json:"date"`
    Duration string `json:"duration"`
    SimID *string `json:"sim_id"`
}

// Copyright (c) 2026 Zeronetsec
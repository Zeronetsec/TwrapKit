// https://github.com/Zeronetsec/TwrapKit

package tphonydevinfo

type Info struct {
    DataEnabled string `json:"data_enabled"`
    DataActivity string `json:"data_activity"`
    DataState string `json:"data_state"`
    DeviceID *string `json:"device_id"`
    DeviceSoftwareVer string `json:"device_software_version"`
    PhoneCount int `json:"phone_count"`
    PhoneType string `json:"phone_type"`
    NetworkOperator string `json:"network_operator"`
    NetworkOperatorName string `json:"network_operator_name"`
    NetworkCountryISO string `json:"network_country_iso"`
    NetworkType string `json:"network_type"`
    NetworkRoaming bool `json:"network_roaming"`
    SimCountryISO string `json:"sim_country_iso"`
    SimOperator string `json:"sim_operator"`
    SimOperatorName string `json:"sim_operator_name"`
    SimSerialNumber *string `json:"sim_serial_number"`
    SimSubscriberID *string `json:"sim_subscriber_id"`
    SimState string `json:"sim_state"`
}

// Copyright (c) 2026 Zeronetsec
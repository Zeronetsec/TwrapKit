// https://github.com/Zeronetsec/TwrapKit

package audioinfo

type Info struct {
    OutputSampleRate string `json:"PROPERTY_OUTPUT_SAMPLE_RATE"`
    OutputFramesPerBuffer string `json:"PROPERTY_OUTPUT_FRAMES_PER_BUFFER"`
    AudioTrackSampleRate int `json:"AUDIOTRACK_SAMPLE_RATE"`
    AudioTrackBufferSize int `json:"AUDIOTRACK_BUFFER_SIZE_IN_FRAMES"`
    AudioTrackLowLatencyRate int `json:"AUDIOTRACK_SAMPLE_RATE_LOW_LATENCY"`
    AudioTrackLowLatencyBuffer int `json:"AUDIOTRACK_BUFFER_SIZE_IN_FRAMES_LOW_LATENCY"`
    PowerSavingSampleRate int `json:"AUDIOTRACK_SAMPLE_RATE_POWER_SAVING"`
    PowerSavingBuffer int `json:"AUDIOTRACK_BUFFER_SIZE_IN_FRAMES_POWER_SAVING"`
    BluetoothA2DP bool `json:"BLUETOOTH_A2DP_IS_ON"`
    WiredHeadset bool `json:"WIREDHEADSET_IS_CONNECTED"`
}

// Copyright (c) 2026 Zeronetsec
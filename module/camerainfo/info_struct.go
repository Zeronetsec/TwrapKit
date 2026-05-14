// https://github.com/Zeronetsec/TwrapKit

package camerainfo

type Info struct {
    ID string `json:"id"`
    Facing string `json:"facing"`
    JPEGOutputSizes []Size `json:"jpeg_output_sizes"`
    FocalLengths []float64 `json:"focal_lengths"`
    AutoExposureModes []string `json:"auto_exposure_modes"`
    PhysicalSize PhysicalSize `json:"physical_size"`
    Capabilities []string `json:"capabilities"`
}

// Copyright (c) 2026 Zeronetsec
package models

type Access struct {
	Uuid      string    `json:"uuid"`
	Referrer  string    `json:"referrer"`
	Cookie    string    `json:"cookie"`
	UserAgent string    `json:"userAgent"`
	Screen    Screen    `json:"screen"`
	Navigator Navigator `json:"navigator"`
}

type Screen struct {
	AvailHeight int32 `json:"availHeight"`
	AvailWidth  int32 `json:"availWidth"`
	Height      int32 `json:"height"`
	Width       int32 `json:"width"`
	ColorDepth  int16 `json:"colorDepth"`
	PixelDepth  int16 `json:"pixelDepth"`
}

type Navigator struct {
	HardwareConcurrency int16    `json:"hardwareConcurrency"`
	Language            string   `json:"language"`
	Languages           []string `json:"languages"`
}

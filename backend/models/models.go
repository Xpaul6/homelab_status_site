package models

type MCInfo struct {
	Online  bool `json:"online"`
	Players struct {
		Online int `json:"online"`
		Max    int `json:"max"`
	} `json:"players"`
}

type SysInfo struct {
	CPU_temp         float64 `json:"cpu_temp"`
	CPU_load_percent float64 `json:"cpu_load_percent"`
	RAM_percent      float64 `json:"ram_percent"`
}


package models

import (
	"time"
	"encoding/json"
)
// CustomTime adalah wrapper untuk time.Time dengan format kustom
type CustomTime struct {
	time.Time
}

// Format kustom untuk tanggal dan waktu
const customLayout = "2006-01-02 15:04:05"


// MarshalJSON mengatur format saat dimarshal ke JSON
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	formatted := ct.Format(customLayout)
	return json.Marshal(formatted)
}
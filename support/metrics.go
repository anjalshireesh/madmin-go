package support

import (
	"time"

	"github.com/minio/madmin-go"
)

type Metrics struct {
	Version   string    `json:"version"` // The metrics version
	Error     string    `json:"error,omitempty"`
	TimeStamp time.Time `json:"timestamp,omitempty"`

	BgHealState madmin.BgHealState `json:"heal_info"`
	IOStats     []IOStats          `json:"iostats"`
	TLSInfo     madmin.TLSInfo     `json:"tls"`
}

type IOStats struct {
	madmin.NodeCommon
	Stats []IOStat `json:"iostats"`
}
type IOStat struct {
	DriveName  string  `json:"drive_name"`
	ReadsPS    float64 `json:"reads_ps"`
	ReadsMBPS  float64 `json:"reads_mbps"`
	WritesPS   float64 `json:"writes_ps"`
	WritesMBPS float64 `json:"writes_mbps"`
	PercUtil   float64 `json:"perc_util"`
	PercUser   float64 `json:"perc_user"`
	PercIOWait float64 `json:"perc_iowait"`
}

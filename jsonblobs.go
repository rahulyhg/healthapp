package main

type ServerFilesystem struct {
	Path string `json:"Path"`
	Size uint64 `json:"Size"`
	Used uint64 `json:"Used"`
	Free uint64 `json:"Free"`
	Pct  uint64 `json:"Pct"`
}

type ServerExtendedInfo struct {
	Kernel      string `json:"Kernel"`
	OS          string `json:"OS"`
	Name        string `json:"name"`
	LastUpdated string `json:"Last Updated"`
	Filesystems []ServerFilesystem
}

type ServerItemResponse struct {
	Name string             `json:"name"`
	Date string             `json:"time"`
	Good bool               `json:"good"`
	Info ServerExtendedInfo `json:"info"`
}

type ServerListResponse struct {
	Servers []ServerItemResponse `json:"servers"`
}

type PrettyAlertInfo struct {
	StartTime    string             `json:"start_time"`
	EndTime      string             `json:"end_time"`
	Duration     string             `json:"duration"`
	ID           string             `json:"alert_id"`
	StateName    string             `json:"state_name"`
	Description  string             `json:"info"`
	Server       ServerExtendedInfo `json:"server"`
	Ongoing      bool               `json:"ongoing"`
	HumanBadName string             `json:"human_bad"`
}

type AlertList struct {
	Active     []PrettyAlertInfo `json:"active"`
	Historical []PrettyAlertInfo `json:"historical"`
}

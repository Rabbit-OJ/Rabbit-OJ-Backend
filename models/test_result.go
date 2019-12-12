package models

type TestResult struct {
	CaseId    int64   `json:"case_id"`
	Status    string  `json:"status"`
	TimeUsed  uint32  `json:"time_used"`
	SpaceUsed float64 `json:"space_used"`
}

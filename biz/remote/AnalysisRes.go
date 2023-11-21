package remote

type Submission struct {
	SubCnt     int   `json:"subCnt"`
	AcceptTime int64 `json:"acceptTime"`
}

type PersonalRes struct {
	Name        string                `json:"name"`
	Penalty     int                   `json:"penalty"`
	SolveCnt    int                   `json:"solveCnt"`
	Rank        int64                 `json:"rank"`
	Submissions map[int64]*Submission `json:"submissions"`
}

type AnalysisRes struct {
	Result []*PersonalRes `json:"result"`
}

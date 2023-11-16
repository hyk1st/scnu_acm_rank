package remote

type CrawlTrainRes interface {
	GetTrainRes(id string) (*VjRespJson, string, error)
	AnalysisRes(interface{}) (*AnalysisRes, error)
}

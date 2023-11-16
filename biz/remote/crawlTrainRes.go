package remote

type CrawlTrainRes interface {
	GetTrainRes() (string, error)
	AnalysisRes(interface{}) (string, error)
}

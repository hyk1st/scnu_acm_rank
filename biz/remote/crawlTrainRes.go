package remote

type CrawlTrainRes interface {
	GetTrainRes() (string, error)
	AnalysisRes(str string) (string, error)
}

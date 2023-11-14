package remote

type CrawlTrainRes interface {
	Login() string
	GetTrainRes() string
	AnalysisRes() string
	Store2DB()
}

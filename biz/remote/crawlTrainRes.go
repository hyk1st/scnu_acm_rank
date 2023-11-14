package remote

type crawlTrainRes interface {
	Login() string
	GetTrainRes() string
	AnalysisRes() string
	Store2DB()
}

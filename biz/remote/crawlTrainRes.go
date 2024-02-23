package remote

import "scnu_acm_rank/biz/model"

type CrawlTrainRes interface {
	GetTrainRes(id string) (string, error)
	AnalysisRes(interface{}) (*AnalysisRes, *model.Competition, error)
}

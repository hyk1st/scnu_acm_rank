package remote

import json2 "github.com/cloudwego/hertz/pkg/common/json"

type FileCrawler struct{}

func (f *FileCrawler) GetTrainRes(id string) (*VjRespJson, string, error) {
	return nil, "", nil
}
func (f *FileCrawler) AnalysisRes(str interface{}) (*AnalysisRes, error) {
	json := str.(string)
	res := AnalysisRes{}
	err := json2.Unmarshal([]byte(json), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

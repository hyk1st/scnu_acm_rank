package remote

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"io/ioutil"
	"net/http"
	"runtime"
	"scnu_acm_rank/biz/model"
	"strconv"
	"time"
)

const getNCRankUrl = "https://ac.nowcoder.com/acm-heavy/acm/contest/real-time-rank-data?token=&id=%v&searchUserName=%v&limit=0&_=%v"
const getNCRankUrlWithPage = "https://ac.nowcoder.com/acm-heavy/acm/contest/real-time-rank-data?token=&id=%v&page=%v&searchUserName=%v&limit=0&_=%v"

type NcRespJson struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data NcData `json:"data"`
}
type NcProblemData struct {
	AcceptedCount int    `json:"acceptedCount"`
	SubmitCount   int    `json:"submitCount"`
	Name          string `json:"name"`
	ProblemID     int    `json:"problemId"`
}

type NcScoreList struct {
	Accepted          bool    `json:"accepted"`
	AcceptedTime      int64   `json:"acceptedTime"`
	FailedCount       int     `json:"failedCount"`
	FinishJudge       bool    `json:"finishJudge"`
	FirstBlood        bool    `json:"firstBlood"`
	FullScore         float64 `json:"fullScore"`
	ProblemID         int     `json:"problemId"`
	Score             float64 `json:"score"`
	SubmissionID      int     `json:"submissionId"`
	Submit            bool    `json:"submit"`
	TimeConsumption   int     `json:"timeConsumption"`
	WaitingJudgeCount int     `json:"waitingJudgeCount"`
}
type NcRankData struct {
	AcceptedCount int           `json:"acceptedCount"`
	ColorLevel    int           `json:"colorLevel"`
	FullScore     float64       `json:"fullScore"`
	PenaltyTime   int           `json:"penaltyTime"`
	Ranking       int           `json:"ranking"`
	School        string        `json:"school"`
	ScoreList     []NcScoreList `json:"scoreList"`
	Team          bool          `json:"team"`
	TotalScore    float64       `json:"totalScore"`
	UID           int           `json:"uid"`
	UserName      string        `json:"userName"`
}
type NcBasicInfo struct {
	RankCount        int    `json:"rankCount"`
	BasicUID         int    `json:"basicUid"`
	ContestID        int    `json:"contestId"`
	PageCount        int    `json:"pageCount"`
	ContestEndTime   int64  `json:"contestEndTime"`
	ContestBeginTime int64  `json:"contestBeginTime"`
	RankType         string `json:"rankType"`
	PageSize         int    `json:"pageSize"`
	Type             int    `json:"type"`
	SearchUserName   string `json:"searchUserName"`
	PageCurrent      int    `json:"pageCurrent"`
}
type NcData struct {
	ProblemData       []NcProblemData `json:"problemData"`
	RankData          []NcRankData    `json:"rankData"`
	IsContestFinished bool            `json:"isContestFinished"`
	BasicInfo         NcBasicInfo     `json:"basicInfo"`
}
type ncCrawler struct {
}

var NcCrawler *ncCrawler

func init() {
	//config := model.Config{}
	//middle.DB.Model(&model.Config{}).First(&config)

	//VJCrawler.userName = config.VjUserName
	//VJCrawler.passWord = config.VjPassWord
	//VJCrawler.cookie = config.VjCookie
	//if flag, err := VJCrawler.checkLoginStatus(); flag && err != nil {
	//	return
	//}
	//if cookie, err := VJCrawler.Login(); err != nil {
	//	VJCrawler.cookie = cookie
	//}
	NcCrawler = &ncCrawler{}
	runtime.KeepAlive(VJCrawler)
}

func (nc *ncCrawler) GetTrainRes(contest string) (string, error) {

	client := &http.Client{}
	fmt.Printf(getNCRankUrl, contest, "%E5%8D%8E%E5%8D%97%E5%B8%88%E8%8C%83%E5%A4%A7%E5%AD%A6", time.Now().UnixMilli())
	req, err := http.NewRequest("GET", fmt.Sprintf(getNCRankUrl, contest, "%E5%8D%8E%E5%8D%97%E5%B8%88%E8%8C%83%E5%A4%A7%E5%AD%A6", time.Now().UnixMilli()), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Apipost client Runtime/+https://www.apipost.cn/")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	temp := NcRespJson{}
	err = json.Unmarshal(bodyText, &temp)
	if err != nil {
		return "", err
	}
	return string(bodyText), nil
}

func GetTrainResWithPage(contest string, page int) (string, error) {

	client := &http.Client{}
	fmt.Printf(getNCRankUrl, contest, "%E5%8D%8E%E5%8D%97%E5%B8%88%E8%8C%83%E5%A4%A7%E5%AD%A6", time.Now().UnixMilli())
	req, err := http.NewRequest("GET", fmt.Sprintf(getNCRankUrlWithPage, contest, page, "%E5%8D%8E%E5%8D%97%E5%B8%88%E8%8C%83%E5%A4%A7%E5%AD%A6", time.Now().UnixMilli()), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Apipost client Runtime/+https://www.apipost.cn/")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	temp := NcRespJson{}
	err = json.Unmarshal(bodyText, &temp)
	if err != nil {
		return "", err
	}
	return string(bodyText), nil
}

func (nc *ncCrawler) AnalysisRes(v interface{}) (*AnalysisRes, *model.Competition, error) {
	resJson := v.(string)
	res := NcRespJson{}
	err := json.Unmarshal([]byte(resJson), &res)
	if err != nil {
		return nil, nil, err
	}

	proId2Name := make(map[int]int, 0)
	for _, v := range res.Data.ProblemData {
		proId2Name[v.ProblemID] = int(v.Name[0] - 'A')
	}
	sli := make([]*PersonalRes, 0, len(res.Data.RankData))

	for i, v := range res.Data.RankData {
		temp := new(PersonalRes)
		temp.Penalty = v.PenaltyTime
		temp.SolveCnt = v.AcceptedCount
		temp.Rank = int64(i)
		temp.Name = v.UserName
		temp.Submissions = make(map[int64]*Submission, 0)
		for _, sub := range v.ScoreList {
			sub2 := new(Submission)
			sub2.SubCnt = sub.FailedCount
			sub2.AcceptTime = sub.AcceptedTime
			temp.Submissions[int64(proId2Name[sub.ProblemID])] = sub2
		}
		sli = append(sli, temp)
	}

	for res.Data.BasicInfo.PageCurrent < res.Data.BasicInfo.PageCount {
		tempJson, err := GetTrainResWithPage(strconv.Itoa(res.Data.BasicInfo.ContestID), res.Data.BasicInfo.PageCurrent+1)
		if err != nil {
			return nil, nil, err
		}
		err = json.Unmarshal([]byte(tempJson), &res)
		if err != nil {
			return nil, nil, err
		}
		for i, v := range res.Data.RankData {
			temp := new(PersonalRes)
			temp.Penalty = v.PenaltyTime
			temp.SolveCnt = v.AcceptedCount
			temp.Rank = int64(i)
			temp.Name = v.UserName
			temp.Submissions = make(map[int64]*Submission, 0)
			for _, sub := range v.ScoreList {
				sub2 := new(Submission)
				sub2.SubCnt = sub.FailedCount
				sub2.AcceptTime = sub.AcceptedTime
				temp.Submissions[int64(proId2Name[sub.ProblemID])] = sub2
			}
			sli = append(sli, temp)
		}
	}

	r := &AnalysisRes{
		Result: sli,
	}
	return r, &model.Competition{
		CpId:      strconv.Itoa(res.Data.BasicInfo.ContestID),
		Kind:      1,
		StartDate: res.Data.BasicInfo.ContestBeginTime,
		Length:    res.Data.BasicInfo.ContestEndTime - res.Data.BasicInfo.ContestBeginTime,
	}, nil
}
